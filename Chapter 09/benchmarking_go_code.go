package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Fibonacci(20)
  }
}

func Fibonacci(n int) int {
  if n <= 1 {
    return n
  }
  return Fibonacci(n-1) + Fibonacci(n-2)
}
