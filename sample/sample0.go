//go:build run
// +build run

package main

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/value"
)

func main() {
	start, _ := value.DateFrom("2019-05-01")
	end := value.NewDate(time.Date(2019, time.May, 31, 0, 0, 0, 0, value.JST))
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Holiday, koyomi.SolarTerm),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
	).Get()
	if err != nil {
		return
	}

	csv, err := k.EncodeCSV()
	if err != nil {
		return
	}
	io.Copy(os.Stdout, bytes.NewReader(csv))
}

/* Copyright 2023 Spiegel
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
