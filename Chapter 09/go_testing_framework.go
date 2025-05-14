package main

import "testing"

func TestCalculateDiscount(t *testing.T) {
  result := CalculateDiscount(100, 10) // Apply 10% discount to $100
  expected := 90.0

  if result != expected {
    t.Errorf("CalculateDiscount(100, 10) = %.2f; want %.2f", result, expected)
  }
}

func CalculateDiscount(price float64, discountPercent float64) float64 {
  return price - (price * discountPercent / 100)
}
