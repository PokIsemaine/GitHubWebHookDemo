package main

import (
	"testing"
	"time"
)

func fib(n int) int {
	time.Sleep(10 * time.Millisecond)
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func BenchmarkDemoTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(10)
	}
}
