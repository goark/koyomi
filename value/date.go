package value

import (
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

// NewDate returns DateJp instance
func NewDate(tm time.Time) DateJp {
	if tm.IsZero() {
		return DateJp{tm}
	}
	ut := tm.Unix()
	_, offset := tm.Zone()
	return DateJp{time.Unix(((ut+int64(offset))/86400)*86400-jstoffset, 0).In(JST)}
}

var timeTemplate = []string{
	"2006-01-02",
	"2006-01",
	time.RFC3339,
}

// DateFrom returns DateJp instance from date string
func DateFrom(s string) (DateJp, error) {
	if len(s) == 0 || strings.EqualFold(s, "null") {
		return NewDate(time.Time{}), nil
	}
	var lastErr error
	for _, tmplt := range timeTemplate {
		if tm, err := time.Parse(tmplt, s); err != nil {
			lastErr = errs.Wrap(err, errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt))
		} else {
			return NewDate(tm), nil
		}
	}
	return NewDate(time.Time{}), lastErr
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
	return t.Time.Year() == dt.Time.Year() && t.Time.Month() == dt.Time.Month() && t.Time.Day() == dt.Time.Day()
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

// WeekdayJp returns the Japanese representation of the weekday for the given DateJp.
// It converts the standard weekday to a WeekdayJp type.
func (t DateJp) WeekdayJp() WeekdayJp {
	return WeekdayJp(t.Weekday())
}

/* Copyright 2020-2025 Spiegel
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
