package main

import (
	"fmt"
	"github.com/golang-patterns/structural/adapter/adapter"
	"github.com/golang-patterns/structural/adapter/jsonlib"
	"github.com/golang-patterns/structural/adapter/xmllib"
)

// Analyzer - интерфейс клиента. Описывает интерфейс с которым наша система хотела бы работать
type analyzer interface {
	Analyze(data xmllib.XML) error
}

//  App - Наш сервис, который умел работать только с форматом XML.
type app struct {
}

// AnalyzeXmlData - анализирует данные в формате xml.
// Прокидываем интерфейс analyzer.
// Теперь наше приложение умеет работать с абстрактным объектом, реализация которого зависит от конкретного типа.
// Наш сервис по-прежнему умеет работать только с xml данными, однако когда в качестве analyzer в метод будет прокинут
// jsonAdapter, он сконвертирует данные в формат json, вызовет библиотеку json, библиотека json проанализирует сконвертированные данные
// и вернет результат.
func (c *app) AnalyzeXmlData(analyzer analyzer, data xmllib.XML) error {
	fmt.Println("app analyze XML data")
	return analyzer.Analyze(data)
}

func main() {
	data := xmllib.XML("hello")
	app := new(app)
	xmlAnalyzer := new(xmllib.Analyzer)
	err := app.AnalyzeXmlData(xmlAnalyzer, data)
	if err != nil {
		return
	}

	jsonAdapter := &adapter.Adapter{
		JsonAnalyzer: new(jsonlib.Analyzer),
	}

	err = app.AnalyzeXmlData(jsonAdapter, data)
	if err != nil {
		return
	}
}
