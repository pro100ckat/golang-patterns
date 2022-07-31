package component

import "fmt"

// Component - конкретный компонент. Поведение этого объекта мы будем расширять через декораторы.
type Component struct{}

// DoSomething - какой-нибудь метод.
func (c *Component) DoSomething() {
	fmt.Println("component do something")
}
