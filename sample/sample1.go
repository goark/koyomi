//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/goark/koyomi/value"
)

func main() {
	flag.Parse()
	argsStr := flag.Args()
	tm := time.Now()
	if len(argsStr) > 0 {
		if len(argsStr) < 3 {
			fmt.Fprintln(os.Stderr, "年月日を指定してください")
			return
		}
		args := make([]int, 3)
		for i := 0; i < 3; i++ {
			num, err := strconv.Atoi(argsStr[i])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			args[i] = num
		}
		tm = time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.Local)
	}
	te := value.NewDate(tm)
	n, y := te.YearEraString()
	if len(n) == 0 {
		fmt.Fprintln(os.Stderr, "正しい年月日を指定してください")
		return
	}
	fmt.Printf("%s%s%d月%d日\n", n, y, te.Month(), te.Day())
}

/* Copyright 2019-2023 Spiegel
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
