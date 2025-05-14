type Car struct {
    Brand string
    Model string
    Year  int
}

myCar := Car{Brand: "Toyota", Model: "Corolla", 
Year: 2020}
fmt.Println(myCar.Brand) // Output: Toyota
