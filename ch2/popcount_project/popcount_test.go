package main

import (
	"popcount_project/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(143404)
	}
}
