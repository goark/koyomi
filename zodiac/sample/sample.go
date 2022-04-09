//go:build run
// +build run

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/goark/koyomi/zodiac"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	for _, s := range args {
		t, err := time.Parse("2006-01-02", s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		kan, shi := zodiac.ZodiacYearNumber(t.Year())
		fmt.Printf("Year %v is %v%v\n", t.Year(), kan, shi)
		kan, shi = zodiac.ZodiacDayNumber(t)
		fmt.Printf("Day %v is %v%v\n", t.Format("2006-01-02"), kan, shi)
	}
}

/* Copyright 2021 Spiegel
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
