package main

import (
	"github.com/golang-patterns/structural/decorator/component"
	"github.com/golang-patterns/structural/decorator/decorators"
)

func main() {
	c := &component.Component{}

	decorator1WithComponent := &decorators.Decorator1{
		Name:      "decorator 1",
		Component: c,
	}
	decorator1WithComponent.DoSomething()

	decorator2WithDecorator1 := &decorators.Decorator2{
		Name:      "decorator 2",
		Component: decorator1WithComponent,
	}

	decorator2WithDecorator1.DoSomething()
}
