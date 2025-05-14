package main

import (
        "fmt"
        "golang.org/x/exp/constraints"
)

// Min returns the minimum value in a slice of ordered elements.
// It returns the zero value of T if the slice is empty.
func Min[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T // Important: Return the zero value for empty slices
		return zero
	}

	min := slice[0]
	for _, v := range slice {
		if v < min {
				min = v
		}
	}
	return min
}

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9, 2, 6}
	floatSlice := []float64{3.14, 1.59, 2.65, 3.58}
	stringSlice := []string{"apple", "banana", "cherry", "date"}

	minInt := Min(intSlice)
	minFloat := Min(floatSlice)
	minString := Min(stringSlice)

	fmt.Printf("Minimum int: %d\n", minInt)
	fmt.Printf("Minimum float: %f\n", minFloat)
	fmt.Printf("Minimum string: %s\n", minString)

	emptySlice := []int{}
	minEmpty := Min(emptySlice)
	fmt.Printf("Minimum empty slice: %d (zero value)\n", minEmpty)
}
