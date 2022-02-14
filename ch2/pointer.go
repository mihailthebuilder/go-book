package main

import "fmt"

func main() {
	a := 50
	p := &a
	fmt.Println(a, p, *p)

	b := 60
	p = &b
	fmt.Println(a, p, *p)

	*p = 30
	fmt.Println(b, p, *p)
}
