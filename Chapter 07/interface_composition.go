type Mover interface {
    Move()
}

type Human interface {
    Speaker
    Mover
}

func (a Athlete) Speak() string {
    return "I am an athlete."
}

func (a Athlete) Move() {
    fmt.Println("Running...")
}

var h Human = Athlete{}
fmt.Println(h.Speak()) // Output: I am an athlete.
h.Move()               // Output: Running...
