//go:build ignore
// +build ignore

package main

import (
	"os"
	"text/template"
	"time"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/value"
)

func main() {
	start, _ := value.DateFrom("2019-05-01")
	end := value.NewDate(time.Date(2019, time.May, 31, 0, 0, 0, 0, value.JST))
	td, err := os.MkdirTemp(os.TempDir(), "sample")
	if err != nil {
		return
	}
	defer func() { _ = os.RemoveAll(td) }()
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Holiday, koyomi.SolarTerm),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
		koyomi.WithTempDir(td),
	).Get()
	if err != nil {
		return
	}

	myTemplate := `| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
{{ range . }}| {{ .Date.StringJp }} | {{ .Date.WeekdayJp.ShortStringJp }} | {{ .Title }} |
{{ end -}}`

	t, err := template.New("").Parse(myTemplate)
	if err != nil {
		return
	}
	if err := t.Execute(os.Stdout, k.Events()); err != nil {
		return
	}
	//Output:
	//| 日付 | 曜日 | 内容 |
	//| ---- |:----:| ---- |
	//| 2019年5月1日 | 水 | 休日 (天皇の即位の日) |
	//| 2019年5月2日 | 木 | 休日 |
	//| 2019年5月2日 | 木 | 八十八夜 |
	//| 2019年5月3日 | 金 | 憲法記念日 |
	//| 2019年5月4日 | 土 | みどりの日 |
	//| 2019年5月5日 | 日 | こどもの日 |
	//| 2019年5月6日 | 月 | 休日 |
	//| 2019年5月6日 | 月 | 立夏 |
	//| 2019年5月21日 | 火 | 小満 |
}

/* Copyright 2025 Spiegel
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
