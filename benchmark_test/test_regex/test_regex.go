package main

import (
	"fmt"
	"testing"
)

func BenchmarkRegexMatch(b *testing.B) {
	pattern := "example"
	text := "This is an example text"
	for i := 0; i < b.N; i++ {
		_ = regexMatch(pattern, text)
	}
}
func BenchmarkStringContainsMatch(b *testing.B) {
	pattern := "example"
	text := "This is an example text"
	for i := 0; i < b.N; i++ {
		_ = stringContainsMatch(pattern, text)
	}
}
func TestingBenchmark(b *testing.B) {
	fmt.Println("Running benchmarks...")
	testing.Benchmark(BenchmarkRegexMatch)
	testing.Benchmark(BenchmarkStringContainsMatch)
}
