package private

// Пакет содержит иерархию типов File и Console.

import "fmt"

// Console - наш объект, поведение которого мы будем расширять через посетителя.
type Console struct {
	// Вы можете столкнуться с ситуацией, когда посетителю нужен будет доступ к приватным полям элементов.
	// В этом случае вы можете либо раскрыть доступ к этим полям, нарушив инкапсуляцию элементов,
	// либо сделать класс посетителя вложенным в класс элемента, если вам повезло писать на языке,
	// который поддерживает вложенность классов.
	// Также можно реализовать методы возвращающие приватные поля.(func GetSomeField() string{return console.someField})
	someField string
}

func (c *Console) GetSomeField() string {
	c.someField = "visitor get some private field from console"
	return c.someField
}

// GetType - метод который уже что-то делает в этом пакете.
func (c *Console) GetType() string {
	return "i am console"
}

// ParseFlags - метод который уже что-то делает в этом пакете.
func (c *Console) ParseFlags() {
	fmt.Println("flags parsed")
}

// Accept - Console принимает некоторое поведение, которое не хочется или нельзя писать в пакете private.
func (c *Console) Accept(visitor Visitor) {
	visitor.VisitConsole(c)
}
