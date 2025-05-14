func modifyValue(x int) {
    x = 100 // Modifies the copy, not the original
}

func main() {
    num := 50
    fmt.Println("Before function call:", num)

    modifyValue(num)
    fmt.Println("After function call:", num)
}
