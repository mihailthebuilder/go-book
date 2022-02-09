package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(reflect.TypeOf(*os.Stdin))
}
