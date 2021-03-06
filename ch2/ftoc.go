// celsius to fahrenheit conversions
// create function outside then convert freeze & boiling (0/10)
// F = C * 9/5 + 32

package main

import "fmt"

const BoilingC = 100.0

const FreezingC = 0.0

func main() {
	boilingC, freezingC := BoilingC, FreezingC
	boilingF, freezingF := CtoF(boilingC), CtoF(freezingC)

	fmt.Printf("Boiling temperature is %vC / %vF\n", boilingC, boilingF)
	fmt.Printf("Freezing temperature is %vC / %vF\n", freezingC, freezingF)
}

func CtoF(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}
