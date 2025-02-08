package value

import (
	"testing"
	"time"
)

func TestNameToString(t *testing.T) {
	testCases := []struct {
		n eraName
		s string
	}{
		{n: EraUnknown, s: ""},
		{n: Meiji, s: "明治"},
		{n: Taisho, s: "大正"},
		{n: Showa, s: "昭和"},
		{n: Heisei, s: "平成"},
		{n: Reiwa, s: "令和"},
		{n: eraName(6), s: ""},
	}
	for _, tc := range testCases {
		if tc.n.String() != tc.s {
			t.Errorf("eraName.String(%d) = \"%v\", want \"%v\".", int(tc.n), tc.n, tc.s)
		}
	}
}

func TestStringToName(t *testing.T) {
	testCases := []struct {
		n eraName
		s string
	}{
		{n: EraUnknown, s: ""},
		{n: EraUnknown, s: "元号"},
		{n: Meiji, s: "明治"},
		{n: Taisho, s: "大正"},
		{n: Showa, s: "昭和"},
		{n: Heisei, s: "平成"},
		{n: Reiwa, s: "令和"},
	}
	for _, tc := range testCases {
		n := EraName(tc.s)
		if tc.n != n {
			t.Errorf("GetName(%v) = \"%v\", want \"%v\".", tc.s, n, tc.n)
		}
	}
}

func TestTimeToEra(t *testing.T) {
	testCases := []struct {
		year    int
		month   time.Month
		day     int
		n       eraName
		yearEra int
		eraStr  string
		yearStr string
	}{
		{year: 1872, month: time.December, day: 31, n: EraUnknown, yearEra: 0, eraStr: "", yearStr: ""},
		{year: 1873, month: time.January, day: 1, n: Meiji, yearEra: 6, eraStr: "明治", yearStr: "6年"},
		{year: 1912, month: time.July, day: 29, n: Meiji, yearEra: 45, eraStr: "明治", yearStr: "45年"},
		{year: 1912, month: time.July, day: 30, n: Taisho, yearEra: 1, eraStr: "大正", yearStr: "元年"},
		{year: 1926, month: time.December, day: 24, n: Taisho, yearEra: 15, eraStr: "大正", yearStr: "15年"},
		{year: 1926, month: time.December, day: 25, n: Showa, yearEra: 1, eraStr: "昭和", yearStr: "元年"},
		{year: 1989, month: time.January, day: 7, n: Showa, yearEra: 64, eraStr: "昭和", yearStr: "64年"},
		{year: 1989, month: time.January, day: 8, n: Heisei, yearEra: 1, eraStr: "平成", yearStr: "元年"},
		{year: 2019, month: time.April, day: 30, n: Heisei, yearEra: 31, eraStr: "平成", yearStr: "31年"},
		{year: 2019, month: time.May, day: 1, n: Reiwa, yearEra: 1, eraStr: "令和", yearStr: "元年"},
		{year: 2118, month: time.December, day: 31, n: Reiwa, yearEra: 100, eraStr: "令和", yearStr: "100年"},
	}
	for _, tc := range testCases {
		tm := time.Date(tc.year, tc.month, tc.day, 0, 0, 0, 0, time.UTC)
		n, y := NewDate(tm).YearEra()
		if tc.n != n || tc.yearEra != y {
			t.Errorf("[%v].Era() = \"%v %d\", want \"%v %d\".", tm, n, y, tc.n, tc.yearEra)
		}
		ns, ys := NewDate(tm).YearEraString()
		if tc.eraStr != ns || tc.yearStr != ys {
			t.Errorf("[%v].Era() = \"%v %v\", want \"%v %v\".", tm, ns, ys, tc.eraStr, tc.yearStr)
		}
	}
}

func TestEraToDate(t *testing.T) {
	testCases := []struct {
		n       eraName
		year    int
		month   time.Month
		day     int
		timeStr string
	}{
		{n: EraUnknown, year: 2019, month: time.May, day: 1, timeStr: "2019-05-01"},
		{n: Heisei, year: 31, month: time.May, day: 1, timeStr: "2019-05-01"},
		{n: Reiwa, year: 1, month: time.May, day: 1, timeStr: "2019-05-01"},
		{n: Showa, year: 100, month: time.January, day: 1, timeStr: "2025-01-01"},
	}
	for _, tc := range testCases {
		tm := NewDateEra(tc.n, tc.year, tc.month, tc.day)
		s := tm.Format("2006-01-02")
		if tc.timeStr != s {
			t.Errorf("Date() = \"%v\", want \"%v\".", s, tc.timeStr)
		}
	}
}

/* Copyright 2019-2025 Spiegel
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
