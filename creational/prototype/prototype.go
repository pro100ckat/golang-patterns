package main

import "fmt"

func clone() {}

type (
	// Cloner Объект, который копируют, называется прототипом
	// Интерфейс прототипов описывает операции клонирования. В большинстве случаев — это единственный метод clone.
	Cloner interface {
		Clone() Cloner
		GetField() string
	}

	// ConcretePrototype Конкретный прототип реализует операцию клонирования самого себя.
	// Помимо банального копирования значений всех полей, здесь могут быть спрятаны различные сложности, о которых не нужно знать клиенту.
	// Например, клонирование связанных объектов, распутывание рекурсивных зависимостей и прочее.
	ConcretePrototype struct {
		field string
	}

	ConcretePrototype2 struct {
		field string
	}

	// PrototypeRegistry Хранилище прототипов облегчает доступ к часто используемым прототипам, храня набор предварительно созданных эталонных, готовых к копированию объектов.
	// Простейшее хранилище может быть построено с помощью хеш-таблицы вида имя-прототипа → прототип.
	// Но для удобства поиска прототипы можно маркировать и другими критериями, а не только условным именем.
	PrototypeRegistry struct {
		items map[string]Cloner
	}
)

func (p *ConcretePrototype) GetField() string {
	return p.field
}
func (p *ConcretePrototype) Clone() Cloner {
	return &ConcretePrototype{
		field: p.field,
	}
}

func (p *ConcretePrototype2) GetField() string {
	return p.field
}
func (p *ConcretePrototype2) Clone() Cloner {
	return &ConcretePrototype2{
		field: p.field,
	}
}

func main() {
	registry := PrototypeRegistry{items: make(map[string]Cloner, 2)}
	p1 := &ConcretePrototype{field: "1"}
	registry.items["1"] = p1
	clone := p1.Clone()
	registry.items["clone"] = clone
	p2 := &ConcretePrototype2{field: "2"}
	registry.items["2"] = p2
	fmt.Println(p1.field == clone.GetField())
}
