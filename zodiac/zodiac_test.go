package zodiac_test

import (
	"testing"
	"time"

	"github.com/goark/koyomi/value"
	"github.com/goark/koyomi/zodiac"
)

func TestKan10(t *testing.T) {
	testCases := []struct {
		kan   zodiac.Kan10
		name  string
		dir   int
		dirJp string
	}{
		{kan: zodiac.Kinoe, name: "甲", dir: 75, dirJp: "東北東微東"},
		{kan: zodiac.Kinoto, name: "乙", dir: 255, dirJp: "西南西微西"},
		{kan: zodiac.Hinoe, name: "丙", dir: 165, dirJp: "南南東微南"},
		{kan: zodiac.Hinoto, name: "丁", dir: 345, dirJp: "北北西微北"},
		{kan: zodiac.Tsutinoe, name: "戊", dir: 165, dirJp: "南南東微南"},
		{kan: zodiac.Tsutinoto, name: "己", dir: 75, dirJp: "東北東微東"},
		{kan: zodiac.Kanoe, name: "庚", dir: 255, dirJp: "西南西微西"},
		{kan: zodiac.Kanoto, name: "辛", dir: 165, dirJp: "南南東微南"},
		{kan: zodiac.Mizunoe, name: "壬", dir: 345, dirJp: "北北西微北"},
		{kan: zodiac.Mizunoto, name: "癸", dir: 165, dirJp: "南南東微南"},
		{kan: zodiac.Kan10(10), name: "甲", dir: 75, dirJp: "東北東微東"},
	}

	for _, tc := range testCases {
		str := tc.kan.String()
		if str != tc.name {
			t.Errorf("zodiac.Kan10(%v) is \"%v\", want %v", uint(tc.kan), str, tc.name)
		}
		dir := tc.kan.Direction()
		if dir != tc.dir {
			t.Errorf("zodiac.Kan10(%v) is \"%v\", want %v", uint(tc.kan), dir, tc.dir)
		}
		dirJp := tc.kan.DirectionJp()
		if dirJp != tc.dirJp {
			t.Errorf("zodiac.Kan10(%v) is \"%v\", want %v", uint(tc.kan), dirJp, tc.dirJp)
		}
	}
}

func TestShi12(t *testing.T) {
	testCases := []struct {
		shi  zodiac.Shi12
		name string
	}{
		{shi: zodiac.Rat, name: "子"},
		{shi: zodiac.Ox, name: "丑"},
		{shi: zodiac.Tiger, name: "寅"},
		{shi: zodiac.Rabbit, name: "卯"},
		{shi: zodiac.Dragon, name: "辰"},
		{shi: zodiac.Snake, name: "巳"},
		{shi: zodiac.Horse, name: "午"},
		{shi: zodiac.Sheep, name: "未"},
		{shi: zodiac.Monkey, name: "申"},
		{shi: zodiac.Rooster, name: "酉"},
		{shi: zodiac.Dog, name: "戌"},
		{shi: zodiac.Boar, name: "亥"},
		{shi: zodiac.Shi12(12), name: "子"},
	}

	for _, tc := range testCases {
		str := tc.shi.String()
		if str != tc.name {
			t.Errorf("zodiac.Shi12(%v) is \"%v\", want %v", uint(tc.shi), str, tc.name)
		}
	}
}

func TestZodiac(t *testing.T) {
	testCases := []struct {
		t       value.DateJp
		kanYear zodiac.Kan10
		shiYear zodiac.Shi12
		kanDay  zodiac.Kan10
		shiDay  zodiac.Shi12
	}{
		{t: value.NewDate(time.Date(1983, time.January, 1, 0, 0, 0, 0, value.JST)), kanYear: zodiac.Mizunoto, shiYear: zodiac.Boar, kanDay: zodiac.Tsutinoto, shiDay: zodiac.Ox},
		{t: value.NewDate(time.Date(1984, time.January, 1, 0, 0, 0, 0, value.JST)), kanYear: zodiac.Kinoe, shiYear: zodiac.Rat, kanDay: zodiac.Kinoe, shiDay: zodiac.Horse},
		{t: value.NewDate(time.Date(1985, time.January, 1, 0, 0, 0, 0, value.JST)), kanYear: zodiac.Kinoto, shiYear: zodiac.Ox, kanDay: zodiac.Kanoe, shiDay: zodiac.Rat},
		{t: value.NewDate(time.Date(2000, time.January, 1, 0, 0, 0, 0, value.JST)), kanYear: zodiac.Kanoe, shiYear: zodiac.Dragon, kanDay: zodiac.Tsutinoe, shiDay: zodiac.Horse},
		{t: value.NewDate(time.Date(2001, time.January, 1, 0, 0, 0, 0, value.JST)), kanYear: zodiac.Kanoto, shiYear: zodiac.Snake, kanDay: zodiac.Kinoe, shiDay: zodiac.Rat},
		{t: value.NewDate(time.Date(2002, time.January, 1, 0, 0, 0, 0, value.JST)), kanYear: zodiac.Mizunoe, shiYear: zodiac.Horse, kanDay: zodiac.Tsutinoto, shiDay: zodiac.Snake},
	}

	for _, tc := range testCases {
		kanYear, shiYear := zodiac.ZodiacYearNumber(tc.t.Year())
		if kanYear != tc.kanYear {
			t.Errorf("result of ZodiacYearNumber(\"%v\") is \"%v\" (kan), want %v", tc.t, kanYear, tc.kanYear)
		}
		if shiYear != tc.shiYear {
			t.Errorf("result of ZodiacYearNumber(\"%v\") is \"%v\" (shi), want %v", tc.t, shiYear, tc.shiYear)
		}
		kanDay, shiDay := zodiac.ZodiacDayNumber(tc.t)
		if kanDay != tc.kanDay {
			t.Errorf("result of ZodiacDayNumber(\"%v\") is \"%v\" (kan), want %v", tc.t, kanDay, tc.kanDay)
		}
		if shiYear != tc.shiYear {
			t.Errorf("result of ZodiacDayNumber(\"%v\") is \"%v\" (shi), want %v", tc.t, shiDay, tc.shiDay)
		}
	}
}

/* Copyright 2021-2023 Spiegel
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
