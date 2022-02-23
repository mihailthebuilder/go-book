package main

import (
	"fmt"
	"popcount_project/popcount"
)

func main() {

	fmt.Println(100 >> (1 * 8))

	var x uint64 = 100
	y := popcount.PopCount(x)

	fmt.Println(y)
}
