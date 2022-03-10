package main

import "fmt"

type IntSet struct {
	a int
}

func (*IntSet) Hullo() int {
	fmt.Println(20)
	return 5
}

func main() {

	var outer IntSet

	var t = outer.Hullo()

	fmt.Println(t)
}
