package zodiac

import (
	"time"

	"github.com/goark/koyomi"
)

type Kan10 uint

const (
	Kinoe     Kan10 = iota // 甲（木の兄）
	Kinoto                 // 乙（木の弟）
	Hinoe                  // 丙（火の兄）
	Hinoto                 // 丁（火の弟）
	Tsutinoe               // 戊（土の兄）
	Tsutinoto              // 己（土の弟）
	Kanoe                  // 庚（金の兄）
	Kanoto                 // 辛（金の弟）
	Mizunoe                // 壬（水の兄）
	Mizunoto               // 癸（水の弟）
	KanMax
)

var kanNames = [KanMax]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

func (k Kan10) String() string {
	return kanNames[k%KanMax]
}

type Shi12 uint

const (
	Rat     Shi12 = iota // 子
	Ox                   // 丑
	Tiger                // 寅
	Rabbit               // 卯
	Dragon               // 辰
	Snake                // 巳
	Horse                // 午
	Sheep                // 未
	Monkey               // 申
	Rooster              // 酉
	Dog                  // 戌
	Boar                 // 亥
	ShiMax
)

var shiNames = [ShiMax]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

func (s Shi12) String() string {
	return shiNames[s%ShiMax]
}

var (
	baseDay  = time.Date(2001, time.January, 1, 0, 0, 0, 0, koyomi.JST) // 2001-01-01 is 甲子
	baseYear = 1984                                                     // Year 1984 is 甲子
)

// ZodiacDayNumber function returns japanese zodiac day number from 2001-01-01.
func ZodiacDayNumber(t koyomi.DateJp) (Kan10, Shi12) {
	n := int64(t.Sub(baseDay).Hours()) / 24
	k := n % int64(KanMax)
	if k < 0 {
		k += int64(KanMax)
	}
	s := n % int64(ShiMax)
	if s < 0 {
		s += int64(ShiMax)
	}
	return Kan10(k), Shi12(s)
}

// ZodiacYearNumber function returns japanese zodiac year number from 1984.
func ZodiacYearNumber(y int) (Kan10, Shi12) {
	n := y - baseYear
	k := n % int(KanMax)
	if k < 0 {
		k += int(KanMax)
	}
	s := n % int(ShiMax)
	if s < 0 {
		s += int(ShiMax)
	}
	return Kan10(k), Shi12(s)
}

/* Copyright 2021-2022 Spiegel
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
