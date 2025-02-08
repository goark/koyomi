//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goark/koyomi/jdn"
	"github.com/goark/koyomi/value"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	for _, s := range args {
		t, err := value.DateFrom(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		j := jdn.GetJDN(t.Time)
		fmt.Printf("Julian Day Number of %v is %v\n", t.Format("2006-01-02"), j)
	}
}

/* Copyright 2022 Spiegel
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
