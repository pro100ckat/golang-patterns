// Package jsonlib - здесь находится другая библиотека, которая умеет анализировать json формат.
// Предположим, что функционал этой библиотеки лучше, чем xmlLibrary, и что нам нужно воспользоваться ей, не меняя при этой клиентский код.
package jsonlib

import "fmt"

// JSON - формат JSON для примера
type JSON []byte

// Analyzer - сервис, который умеет анализировать json формат.
type Analyzer struct{}

// Analyze анализирует данные в формате JSON.
func (l *Analyzer) Analyze(data JSON) error {
	fmt.Println("json analyzer analyze JSON data")
	return nil
}
