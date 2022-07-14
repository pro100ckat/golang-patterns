package main

import (
	"fmt"
	"github.com/golang-patterns/creational/singleton/singleton"
)

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			pointer := singleton.GetInstance()
			// Печатаем ячейку памяти, чтобы показать, что ссылаемся на один и тот же объект в памяти.
			// Смотрим тесты.
			fmt.Printf("%p", pointer)
		}()
	}
	fmt.Scanln()
}
