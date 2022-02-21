package main

import (
	"fmt"
	"tempconv_project/tempconv"
)

func main() {
	c := tempconv.AbsoluteZeroC
	k := tempconv.CToK(c)
	f := tempconv.KToF(k)

	fmt.Println(c, k, f)
}
