package writer

import (
	"fmt"
	"github.com/golang-patterns/behavior/visitor/private"
)

// Writer - пакет, который будет реализовывать интерфейс посетителя. Короче говоря, для пакета, закрытого от изменений, здесь реализуем нужное нам поведение.
type Writer struct{}

// VisitFile - конкретная реализация. Пишем логику, которая что-то будет делать с File.
func (w *Writer) VisitFile(f *private.File) {
	fmt.Println("writer write to file")
}

// VisitConsole - конкретная реализация. Пишем логику, которая что-то будет делать с console.
func (w *Writer) VisitConsole(c *private.Console) {
	fmt.Println("writer write to console")
}
