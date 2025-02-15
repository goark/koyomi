package koyomi_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/value"
)

func ExampleKoyomi() {
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

	csv, err := k.EncodeCSV()
	if err != nil {
		return
	}
	if _, err := io.Copy(os.Stdout, bytes.NewReader(csv)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//Output:
	//"Date","Title"
	//"2019-05-01","休日 (天皇の即位の日)"
	//"2019-05-02","休日"
	//"2019-05-02","八十八夜"
	//"2019-05-03","憲法記念日"
	//"2019-05-04","みどりの日"
	//"2019-05-05","こどもの日"
	//"2019-05-06","休日"
	//"2019-05-06","立夏"
	//"2019-05-21","小満"
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
