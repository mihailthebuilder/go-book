// create cost of 212 float
// c = (f-32)*5 . 9
// say boiling point in C or F

package main

import "fmt"

const BoilingF = 212.0

func main() {
	boilingF := BoilingF
	celsiusF := (boilingF - 32) * 5 / 9
	fmt.Printf("Boiling point is %v Celsius / %v Fahrenheit", celsiusF, boilingF)
}
