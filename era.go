package koyomi

import (
	"fmt"
	"time"
)

//eraName は元号名を表す内部型です。
type eraName int

const (
	EraUnknown eraName = iota //不明な元号
	Meiji                     //明治
	Taisho                    //大正
	Showa                     //昭和
	Heisei                    //平成
	Reiwa                     //令和
)

var eraString = map[eraName]string{
	EraUnknown: "",
	Meiji:      "明治",
	Taisho:     "大正",
	Showa:      "昭和",
	Heisei:     "平成",
	Reiwa:      "令和",
}

//koyomi.EraName() 関数は元号の文字列から元号名 eraName を取得します。
//該当する元号名がない場合は era.Unknown を返します。
func EraName(s string) eraName {
	for k, v := range eraString {
		if v == s {
			return k
		}
	}
	return EraUnknown
}

func (e eraName) String() string {
	if s, ok := eraString[e]; ok {
		return s
	}
	return ""
}

var (
	eraTrigger = map[eraName]DateJp{ //各元号の起点
		Meiji:  NewDate(time.Date(1873, time.January, 1, 0, 0, 0, 0, JST)),   //明治（の改暦） : 1873-01-01
		Taisho: NewDate(time.Date(1912, time.July, 30, 0, 0, 0, 0, JST)),     //大正 : 1912-07-30
		Showa:  NewDate(time.Date(1926, time.December, 25, 0, 0, 0, 0, JST)), //昭和 : 1926-12-25
		Heisei: NewDate(time.Date(1989, time.January, 8, 0, 0, 0, 0, JST)),   //平成 : 1989-01-08
		Reiwa:  NewDate(time.Date(2019, time.May, 1, 0, 0, 0, 0, JST)),       //令和 : 2019-05-01
	}
	eraSorted = []eraName{Reiwa, Heisei, Showa, Taisho, Meiji} //ソートされた元号の配列（降順）
)

//koyomi.NewDateEra() 関数は 元号・年・月・日・時・分・秒・タイムゾーン を指定して koyomi.DateJp 型のインスタンスを返します。
//起点が定義されない元号を指定した場合は西暦として処理します。
func NewDateEra(en eraName, year int, month time.Month, day int) DateJp {
	ofset := 0
	if dt, ok := eraTrigger[en]; ok {
		ofset = dt.Year() - 1
	}
	return NewDate(time.Date(year+ofset, month, day, 0, 0, 0, 0, JST))
}

//koyomi.DateJp.Era() メソッドは元号名 koyomi.eraName のインスタンスを返します。
//元号が不明の場合は era.Unknown を返します。
func (t DateJp) Era() eraName {
	for _, es := range eraSorted {
		if !t.Before(eraTrigger[es]) {
			return es
		}
	}
	return EraUnknown

}

//koyomi.DateJp.YearEra() メソッドは元号付きの年の値を返します。
//元号が不明の場合は (era.Unknown, 0) を返します。
func (t DateJp) YearEra() (eraName, int) {
	era := t.Era()
	if era == EraUnknown {
		return EraUnknown, 0
	}
	year := t.Year() - eraTrigger[era].Year() + 1
	if era == Meiji { //明治のみ5年のオフセットを加算する
		return era, year + 5
	}
	return era, year
}

//koyomi.DateJp.YearEraString() メソッドは元号付きの年の値を文字列で返します。
//元号が不明の場合は空文字列を返します。
func (t DateJp) YearEraString() (string, string) {
	era, year := t.YearEra()
	if era == EraUnknown || year < 1 {
		return "", ""
	}
	if year == 1 {
		return era.String(), "元年"
	}
	return era.String(), fmt.Sprintf("%d年", year)
}

/* Copyright 2019-2022 Spiegel
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
