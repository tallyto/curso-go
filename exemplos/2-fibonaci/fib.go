package fib

import (
	"fmt"
)

// Função Fibonacci Recursiva
func fibonacciRecursivo(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursivo(n-1) + fibonacciRecursivo(n-2)
}

// Função Fibonacci Imperativa
func fibonacciImperativo(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func main() {
	fmt.Println("Fibonacci recursivo:")
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacciRecursivo(i), " ")
	}
	fmt.Println("\nFibonacci imperativo:")
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacciImperativo(i), " ")
	}
}
