// Package xmllib - Предположим, что наш сервис умеет работать только с xml форматом и для этого у нас есть специальная библиотека xmllib,
// код которой мы по каким-то причинам не можем изменять. Эта библиотека каким-то образом анализирует xml формат.
// Однако, нашей библиотеке понадобился какой-то крутой функционал, который есть только в другой библиотеке jsonLib.
// Для того чтобы воспользоваться функционалом библиотеки jsonLib, не меняя при этом наш клиентский код, нужно написать adapter.
package xmllib

import (
	"fmt"
)

// XML формат XML для примера
type XML []byte

// Analyzer - Сервис
type Analyzer struct{}

// Analyze анализирует данные в формате XML.
func (l *Analyzer) Analyze(data XML) error {
	fmt.Println("XML library analyze XML data")
	return nil
}
