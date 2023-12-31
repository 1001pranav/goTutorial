package main

import "fmt"

func fibonacci(n int) []int {
	fibSeries := make([]int, n)
	fibSeries[0], fibSeries[1] = 0, 1

	for i := 2; i < n; i++ {
		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
	}

	return fibSeries
}

func main() {
	n := 10 // You can change this to generate a different number of Fibonacci numbers
	result := fibonacci(n)

	fmt.Printf("Fibonacci Series of %d numbers: %v\n", n, result)
}
