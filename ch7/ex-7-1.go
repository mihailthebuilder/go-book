package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	pString := string(p)

	scanner := bufio.NewScanner(strings.NewReader(pString))
	scanner.Split(bufio.ScanWords)

	count := 0

	for scanner.Scan() {
		count++
	}

	*c += WordCounter(count)
	return count, nil
}

func main() {
	var c WordCounter
	c.Write([]byte("hello howdy, there"))
	fmt.Println(c)
}
