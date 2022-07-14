package main

import "fmt"

func main() {
	defaultBuilder := NewBuilder("default")
	builder1 := NewBuilder("builder1")

	d := newDirector(defaultBuilder)
	product := d.Build()
	fmt.Println(product)

	d.setBuilder(builder1)
	product = d.Build()
	fmt.Println(product)
}
