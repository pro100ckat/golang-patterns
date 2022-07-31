package compositor

import "fmt"

// Component - помещаем в пакет интерфейс.
// Compositor сможет работать со всеми типами, реализующими данный интерфейс.
type Component interface {
	Search()
}

// Compositor - наш компоновщик, который будет выполнять одну и ту же операцию для разных компонентов.
// Компоновщик группирует множество объектов в древовидную структуру, а затем работать с ней так, как будто это единичный объект.
type Compositor struct {
	Components []Component
	Name       string
}

// Search - функция, которая что-то делает со всеми компонентами. Реализовывать функцию каждый компонент будет по своему.
func (c *Compositor) Search() {
	fmt.Println("search in", c.Name)
	for _, component := range c.Components {
		// Каждый тип реализует собственное поведение.
		component.Search()
	}
}

// AddComponents - добавляет компоненты в Compositor.
func (c *Compositor) AddComponents(component Component) {
	c.Components = append(c.Components, component)
}
