package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}

	fmt.Println(time.Since(start).Seconds())

}
