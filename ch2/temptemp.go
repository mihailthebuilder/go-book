package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
	c := Celsius(50)
	fmt.Println(c, c-50)

	f := Fahrenheit(30)
	fmt.Println(c - Celsius(f))
}
