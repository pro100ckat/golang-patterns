// Package adapter - создаем адаптер, который будет реализовывать метод "Analyze(data xmllib.XML) error".
package adapter

import (
	"fmt"
	"github.com/golang-patterns/structural/adapter/jsonlib"
	"github.com/golang-patterns/structural/adapter/xmllib"
)

// Adapter Адаптер оборачивает один из объектов, так что другой объект даже не знает о наличии первого.
type Adapter struct {
	JsonAnalyzer *jsonlib.Analyzer
}

// Analyze Конвертирует XML данные в json формат и анализирует json формат.
func (a *Adapter) Analyze(data xmllib.XML) error {
	fmt.Println("adapter convert XML data to JSON")
	return a.JsonAnalyzer.Analyze(jsonlib.JSON(data))
}
