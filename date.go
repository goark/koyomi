package koyomi

import (
	"strconv"
	"strings"
	"time"

	"github.com/spiegel-im-spiegel/errs"
)

const jstoffset = 9 * 60 * 60

var jst = time.FixedZone("JST", jstoffset)

//DateJp is wrapper class of time.Time
type DateJp struct {
	time.Time
}

//NewDate returns DateJp instance
func NewDate(tm time.Time) DateJp {
	if tm.IsZero() {
		return DateJp{tm}
	}
	ut := tm.Unix()
	_, offset := tm.Zone()
	return DateJp{time.Unix(((ut+int64(offset))/86400)*86400-jstoffset, 0).In(jst)}
}

var timeTemplate = []string{
	"2006-01-02",
	"2006-01",
	time.RFC3339,
}

//DateFrom returns DateJp instance from date string
func DateFrom(s string) (DateJp, error) {
	if len(s) == 0 || strings.EqualFold(s, "null") {
		return NewDate(time.Time{}), nil
	}
	var lastErr error
	for _, tmplt := range timeTemplate {
		if tm, err := time.Parse(tmplt, s); err != nil {
			lastErr = errs.WrapWithCause(err, nil, errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt))
		} else {
			return NewDate(tm), nil
		}
	}
	return NewDate(time.Time{}), lastErr
}

//UnmarshalJSON returns result of Unmarshal for json.Unmarshal()
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

//MarshalJSON returns time string with RFC3339 format
func (t *DateJp) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("\"\""), nil
	}
	if t.IsZero() {
		return []byte("\"\""), nil
	}
	return []byte(strconv.Quote(t.String())), nil
}

func (t DateJp) String() string {
	return t.Format("2006-01-02")
}

//Equal reports whether t and dt represent the same time instant.
func (t DateJp) Equal(dt DateJp) bool {
	return t.Time.Year() == dt.Time.Year() && t.Time.Month() == dt.Time.Month() && t.Time.Day() == dt.Time.Day()
}

//Before reports whether the DateJp instant t is before dt.
func (t DateJp) Before(dt DateJp) bool {
	return !t.Equal(dt) && t.Time.Before(dt.Time)
}

//After reports whether the DateJp instant t is after dt.
func (t DateJp) After(dt DateJp) bool {
	return !t.Equal(dt) && t.Time.After(dt.Time)
}

/* Copyright 2020 Spiegel
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
