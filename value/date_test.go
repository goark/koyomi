package value

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type ForTestStruct struct {
	DateTaken DateJp `json:"date_taken,omitempty"`
}

func TestWeekdayJp(t *testing.T) {
	testCases := []struct {
		s     string
		name  string
		short string
	}{
		{s: "2024-06-01T09:00:00+09:00", name: "土曜日", short: "土"},
		{s: "2024-06-02T09:00:00+09:00", name: "日曜日", short: "日"},
		{s: "2024-06-03T09:00:00+09:00", name: "月曜日", short: "月"},
		{s: "2024-06-04T09:00:00+09:00", name: "火曜日", short: "火"},
		{s: "2024-06-05T09:00:00+09:00", name: "水曜日", short: "水"},
		{s: "2024-06-06T09:00:00+09:00", name: "木曜日", short: "木"},
		{s: "2024-06-07T09:00:00+09:00", name: "金曜日", short: "金"},
	}

	for _, tc := range testCases {
		dt, err := DateFrom(tc.s)
		if err != nil {
			t.Errorf("value.DateFrom() is \"%v\", want nil.", err)
			continue
		}
		wd := dt.WeekdayJp()
		if wd.String() != tc.name {
			t.Errorf("DateJp.WeekdayJp() is \"%v\", want \"%v\".", wd.String(), tc.name)
		}
		if wd.ShortString() != tc.short {
			t.Errorf("DateJp.WeekdayJp() is \"%v\", want \"%v\".", wd.ShortString(), tc.short)
		}
	}
}
func TestUnmarshal(t *testing.T) {
	testCases := []struct {
		s    string
		str  string
		str2 string
		jsn  string
	}{
		{s: `{"date_taken": "2005-03-26T00:00:00+01:00"}`, str: "2005-03-26", str2: "2005-03-26T00:00:00+09:00", jsn: `{"date_taken":"2005-03-26"}`},
		{s: `{"date_taken": "2005-03-26T12:34:56+09:00"}`, str: "2005-03-26", str2: "2005-03-26T00:00:00+09:00", jsn: `{"date_taken":"2005-03-26"}`},
		{s: `{"date_taken": "2005-03-26"}`, str: "2005-03-26", str2: "2005-03-26T00:00:00+09:00", jsn: `{"date_taken":"2005-03-26"}`},
		{s: `{"date_taken": ""}`, str: "0001-01-01", str2: "0001-01-01T00:00:00Z", jsn: `{"date_taken":""}`},
		{s: `{}`, str: "0001-01-01", str2: "0001-01-01T00:00:00Z", jsn: `{"date_taken":""}`},
	}

	for _, tc := range testCases {
		tst := &ForTestStruct{}
		if err := json.Unmarshal([]byte(tc.s), tst); err != nil {
			t.Errorf("json.Unmarshal() is \"%v\", want nil.", err)
			continue
		}
		str := tst.DateTaken.String()
		if str != tc.str {
			t.Errorf("DateJp = \"%v\", want \"%v\".", str, tc.str)
		}
		str2 := tst.DateTaken.Time.Format(time.RFC3339)
		if str2 != tc.str2 {
			t.Errorf("DateJp = \"%v\", want \"%v\".", str2, tc.str2)
		}
		b, err := json.Marshal(tst)
		if err != nil {
			t.Errorf("json.Marshal() is \"%v\", want nil.", err)
			continue
		}
		str = string(b)
		if str != tc.jsn {
			t.Errorf("ForTestStruct = \"%v\", want \"%v\".", str, tc.jsn)
		}
	}
}

func TestUnmarshalErr(t *testing.T) {
	data := `{"date_taken": "2005/03/26"}`
	tst := &ForTestStruct{}
	if err := json.Unmarshal([]byte(data), tst); err == nil {
		t.Error("Unmarshal() error = nil, not want nil.")
	} else {
		fmt.Printf("Info: %+v\n", err)
	}
}

func TestEqual(t *testing.T) {
	dt1, _ := DateFrom("2020-03-31")
	dt2, _ := DateFrom("2020-04-01")
	if dt1.Equal(dt2) {
		t.Error("DateJp.Equal() is true, want false.")
	}
	if !dt1.Before(dt2) {
		t.Error("DateJp.Before() is false, want true.")
	}
	if dt1.After(dt2) {
		t.Error("DateJp.After() is true, want false.")
	}
	if !dt1.AddDay(1).Equal(dt2) {
		t.Error("DateJp.Equal() is false, want true.")
	}
}

/* Copyright 2020-2024 Spiegel
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
