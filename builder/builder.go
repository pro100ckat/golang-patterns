package main

import "fmt"

// Builder Интерфейс строителя объявляет возможные шаги конструирования продуктов, общие для всех видов строителей.
type Builder interface {
	BuildField1()
	BuildField2()
	BuildProduct() *Product
}

// NewBuilder создает строителя по типу.
func NewBuilder(builderName string) Builder {
	var builder Builder
	switch builderName {
	case "builder1":
		builder = &builder1{}
	default:
		builder = &defaultBuilder{}
	}
	return builder
}

// builder1 Конкретные строители реализуют строительные шаги, каждый по-своему. Конкретные строители могут производить разнородные объекты, не имеющие общего интерфейса, но в данном примере объект один.
type builder1 struct {
	Filed1 string
	Filed2 string
}

func (b *builder1) BuildProduct() *Product {
	return &Product{
		Field1: b.Filed1,
		Field2: b.Filed2,
	}
}

func (b *builder1) BuildField1() {
	fmt.Println("builder1 BuildField1")
	b.Filed1 = "Field1"
}

func (b *builder1) BuildField2() {
	fmt.Println("builder1 BuildField2")
	b.Filed2 = "Field2"
}

// defaultBuilder ... .
type defaultBuilder struct {
	Field1 string
	Field2 string
}

func (b *defaultBuilder) BuildProduct() *Product {
	return &Product{
		Field1: b.Field1,
		Field2: b.Field2,
	}
}

func (b *defaultBuilder) BuildField1() {
	fmt.Println("defaultBuilder BuildField1")
	b.Field1 = "DefaultField1"
}

func (b *defaultBuilder) BuildField2() {
	fmt.Println("defaultBuilder BuildField2")
	b.Field2 = "DefaultField2"
}
