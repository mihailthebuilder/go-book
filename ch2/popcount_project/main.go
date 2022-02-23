package main

import (
	"fmt"
	"popcount_project/popcount"
)

func main() {
	var x uint64 = 100
	y := popcount.PopCount(x)

	fmt.Println(y)
}
