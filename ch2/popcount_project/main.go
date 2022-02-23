package main

import (
	"fmt"
	"popcount_project/popcount"
)

func main() {

	var x uint64 = 154327
	y := popcount.PopCount(x)

	fmt.Println(y)
}
