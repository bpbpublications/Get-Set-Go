func PrintSlice[T any](slice []T) {
	for _, value := range slice {
		fmt.Println(value)
	}
}
   