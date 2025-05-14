type Builder struct {
    value int
}

func (b *Builder) Add(v int) *Builder {
    b.value += v
    return b
}

func (b *Builder) Multiply(v int) *Builder {
    b.value *= v
    return b
}

builder := &Builder{}
result := builder.Add(5).Multiply(2).value
fmt.Println(result) // Output: 10
