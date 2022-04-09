//go:build run
// +build run

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/goark/koyomi/era"
)

func main() {
	flag.Parse()
	argsStr := flag.Args()

	if len(argsStr) < 4 {
		fmt.Fprintln(os.Stderr, "元号 年 月 日 を指定してください")
		return
	}
	name := argsStr[0]
	args := make([]int, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi(argsStr[i+1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		args[i] = num
	}
	te := era.Date(era.Name(name), args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.Local)
	fmt.Println(te.Format("西暦2006年1月2日"))
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
