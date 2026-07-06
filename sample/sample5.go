//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/goark/koyomi/value"
)

func main() {
	start := value.NewDateYMD(2024, time.June, 1)
	until := start.AddDay(4)

	seq, err := start.IterDay(1, until)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for d := range seq {
		fmt.Println(d.String())
	}
}
