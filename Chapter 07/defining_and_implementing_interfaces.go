type Speaker interface {
    Speak() string
}

func (d Dog) Speak() string {
    return "Woof!"
}

var s Speaker = Dog{}
fmt.Println(s.Speak()) // Output: Woof!
