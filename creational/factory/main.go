package main

import "fmt"

type Factory interface {
	CreateSomething()
}

type ConcreteProduct1 struct {
}

func (p ConcreteProduct1) CreateSomething() {}

type ConcreteProduct2 struct {
}

func (p ConcreteProduct2) CreateSomething() {

}

func New(factoryType string) Factory {
	switch factoryType {
	case "1":
		return new(ConcreteProduct1)
	default:
		return new(ConcreteProduct2)
	}
}

func main() {
	factory := New("1")
	//factory.CreateSomething("")
	fmt.Println(factory)
}
