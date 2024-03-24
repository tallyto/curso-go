package fib

import "testing"

func BenchmarkFibImperativo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciImperativo(20)
	}
}

func BenchmarkFibRecursivo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursivo(20)
	}
}
