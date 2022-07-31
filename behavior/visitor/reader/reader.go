package reader

import (
	"fmt"
	"github.com/golang-patterns/behavior/visitor/private"
)

// Пакет может наследовать код пакета private(reader импортирует private), так как код в пакет reader мы добавляем для расширения функционала private.

// Reader - пакет, который будет реализовывать интерфейс посетителя. Короче говоря, для пакета, закрытого от изменений, здесь реализуем нужное нам поведение.
type Reader struct{}

// VisitFile - конкретная реализация. Пишем логику, которая что-то будет делать с File.
func (r *Reader) VisitFile(f *private.File) {
	fmt.Println("reader read the file")
}

// VisitConsole - конкретная реализация. Пишем логику, которая что-то будет делать с Console.
func (r *Reader) VisitConsole(c *private.Console) {
	// Reader может получить доступ к приватным полям Concosle, если у Console реализовать методы, которые возвращают эти поля.
	fmt.Println(c.GetSomeField())
	fmt.Println("reader read the console")
}
