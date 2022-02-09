package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	var s, sep string

	args := os.Args

	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	fmt.Println(s)
	fmt.Println(time.Since(start).Seconds())
}
