type A [3]int
type B [3]int

var a A = [3]int{1, 2, 3}
// var b B = a       // This will cause an error
var b B = B(a)       // Explicit conversion works