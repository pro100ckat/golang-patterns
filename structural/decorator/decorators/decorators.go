package decorators

import "fmt"

const logText = "оборачивает компонент и добавляет в него свое поведение"

type ComponentInterface interface {
	DoSomething()
}

// Decorator1 - обертка.
type Decorator1 struct {
	Component ComponentInterface
	Name      string
}

func (d1 *Decorator1) DoSomething() {
	fmt.Println(d1.Name, logText)
	d1.Component.DoSomething()
}

type Decorator2 struct {
	Component ComponentInterface
	Name      string
}

func (d2 *Decorator2) DoSomething() {
	fmt.Println(d2.Name, logText)
	d2.Component.DoSomething()
}

type Decorator3 struct {
	Component ComponentInterface
	Name      string
}

func (d3 *Decorator3) DoSomething() {
	fmt.Println(d3.Name, logText)
	d3.Component.DoSomething()
}
