package value

import (
	"iter"
	"strconv"
	"strings"
	"time"

	"github.com/goark/errs"
)

// WeekdayJp is a type that represents the days of the week in the Japanese context.
// It is based on the time.Weekday type from the standard library.
type WeekdayJp time.Weekday

const (
	Sunday    WeekdayJp = WeekdayJp(time.Sunday) + iota // 日曜日
	Monday                                              // 月曜日
	Tuesday                                             // 火曜日
	Wednesday                                           // 水曜日
	Thursday                                            // 木曜日
	Friday                                              // 金曜日
	Saturday                                            // 土曜日
)

var weekdayNames = [7]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}
var weekdayShortNames = [7]string{"日", "月", "火", "水", "木", "金", "土"}

// String returns the English name of the Japanese weekday (WeekdayJp)
// by converting it to the standard time.Weekday type and calling its String method.
func (w WeekdayJp) String() string {
	return time.Weekday(w).String()
}

// StringJp returns the Japanese name of the WeekdayJp if it is between Sunday and Saturday.
// If the WeekdayJp is out of this range, it returns the standard time.Weekday string representation.
func (w WeekdayJp) StringJp() string {
	if w < Sunday || w > Saturday {
		return time.Weekday(w).String()
	}
	return weekdayNames[w]
}

// ShortStringJp returns the short Japanese name of the WeekdayJp.
// If the WeekdayJp is not within the valid range (Sunday to Saturday),
// it returns the default string representation of the time.Weekday.
func (w WeekdayJp) ShortStringJp() string {
	if w < Sunday || w > Saturday {
		return time.Weekday(w).String()
	}
	return weekdayShortNames[w]
}

var (
	jstoffset = int64((9 * time.Hour).Seconds())
	JST       = time.FixedZone("JST", int(jstoffset)) // Japan standard Time
)

// DateJp is wrapper class of time.Time
type DateJp struct {
	time.Time
}

// NewDate converts tm to DateJp normalized to a JST calendar day.
//
// For non-zero time values, the returned value is truncated to 00:00:00 in JST
// for the date that tm represents in its original location/offset. This keeps
// comparisons and date-based operations stable regardless of input time-of-day.
// For zero time values, NewDate preserves zero and returns it as-is.
func NewDate(tm time.Time) DateJp {
	if tm.IsZero() {
		return DateJp{tm}
	}
	ut := tm.Unix()
	_, offset := tm.Zone()
	return DateJp{time.Unix(((ut+int64(offset))/86400)*86400-jstoffset, 0).In(JST)}
}

// NewDateYMD returns DateJp from explicit year/month/day values.
//
// The generated value is interpreted in JST and normalized by NewDate.
func NewDateYMD(year int, month time.Month, day int) DateJp {
	return NewDate(time.Date(year, month, day, 0, 0, 0, 0, JST))
}

var timeTemplate = []string{
	"2006-01-02",
	"2006-01",
	time.RFC3339,
}

// DateFrom parses s and returns a normalized DateJp value.
//
// Accepted formats are:
//   - 2006-01-02
//   - 2006-01
//   - RFC3339
//
// The parsed value is normalized by NewDate (JST date boundary).
// Empty string or "null" returns a zero DateJp with nil error.
// If all parse attempts fail, DateFrom returns zero DateJp and aggregated
// wrapped parse errors for all attempted templates. Each wrapped error includes
// context keys "time_string" and "time_template".
func DateFrom(s string) (DateJp, error) {
	if len(s) == 0 || strings.EqualFold(s, "null") {
		return NewDate(time.Time{}), nil
	}
	errlist := &errs.Errors{}
	for _, tmplt := range timeTemplate {
		if tm, err := time.Parse(tmplt, s); err != nil {
			errlist.Add(errs.Wrap(err, errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt)))
		} else {
			return NewDate(tm), nil
		}
	}
	return NewDate(time.Time{}), errlist.ErrorOrNil()
}

// UnmarshalJSON returns result of Unmarshal for json.Unmarshal()
func (t *DateJp) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	tm, err := DateFrom(s)
	if err != nil {
		return err
	}
	*t = tm
	return nil
}

// MarshalJSON returns time string with RFC3339 format
func (t *DateJp) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("\"\""), nil
	}
	if t.IsZero() {
		return []byte("\"\""), nil
	}
	return []byte(strconv.Quote(t.String())), nil
}

// String returns the DateJp value formatted as a string in the "2006-01-02" layout.
// It implements the Stringer interface.
func (t DateJp) String() string {
	return t.Format("2006-01-02")
}

// StringJp returns the date in Japanese format (YYYY年M月D日).
func (t DateJp) StringJp() string {
	return t.Format("2006年1月2日")
}

// Equal reports whether t and dt represent the same time instant.
func (t DateJp) Equal(dt DateJp) bool {
	return t.Year() == dt.Year() && t.Month() == dt.Month() && t.Day() == dt.Day()
}

// Before reports whether the DateJp instant t is before dt.
func (t DateJp) Before(dt DateJp) bool {
	return !t.Equal(dt) && t.Time.Before(dt.Time)
}

// After reports whether the DateJp instant t is after dt.
func (t DateJp) After(dt DateJp) bool {
	return !t.Equal(dt) && t.Time.After(dt.Time)
}

// AddDate method adds years/months/days and returns new Date instance.
func (t DateJp) AddDate(years int, months int, days int) DateJp {
	return NewDate(t.Time.AddDate(years, months, days))
}

// AddDay method adds n days and returns new Date instance.
func (t DateJp) AddDay(days int) DateJp {
	return t.AddDate(0, 0, days)
}

// IterDay returns an iterator that yields DateJp values from t to until.
//
// The iterator yields both endpoints and advances by AddDay(count) on each step.
// A non-zero count is required. If count and until cannot reach each other
// (direction mismatch or non-divisible distance), IterDay returns an error.
func (t DateJp) IterDay(count int, until DateJp) (iter.Seq[DateJp], error) {
	if count == 0 {
		return nil, errs.Wrap(ErrInvalidCount, errs.WithContext("count", count), errs.WithContext("start", t.String()), errs.WithContext("until", until.String()))
	}

	diff := int(until.Sub(t.Time).Hours() / 24)
	if (count > 0 && diff < 0) || (count < 0 && diff > 0) {
		return nil, errs.Wrap(ErrInfiniteLoop, errs.WithContext("count", count), errs.WithContext("start", t.String()), errs.WithContext("until", until.String()))
	}

	absCount := count
	if absCount < 0 {
		absCount = -absCount
	}
	if diff%absCount != 0 {
		return nil, errs.Wrap(ErrInfiniteLoop, errs.WithContext("count", count), errs.WithContext("start", t.String()), errs.WithContext("until", until.String()))
	}

	steps := diff / count
	return func(yield func(DateJp) bool) {
		for i := 0; i <= steps; i++ {
			if !yield(t.AddDay(i * count)) {
				return
			}
		}
	}, nil
}

// WeekdayJp returns the Japanese representation of the weekday for the given DateJp.
// It converts the standard weekday to a WeekdayJp type.
func (t DateJp) WeekdayJp() WeekdayJp {
	return WeekdayJp(t.Weekday())
}

/* Copyright 2020-2026 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
