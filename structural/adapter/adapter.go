package main

import "fmt"

// analyzer интерфейс клиента. Описывает интерфейс с которым наша система хотела бы работать
type analyzer interface {
	AnalyzeXmlData(data xml) error
}

//  client Клиентский код.
type client struct {
}

// AnalyzeXmlData анализирует xml формат.
func (c *client) AnalyzeXmlData(analyzer analyzer, data xml) error {
	fmt.Println("client analyze xml data")
	return analyzer.AnalyzeXmlData(data)
}

// формат xml для примера
type xml []byte

//  xmlLibrary Сервис
type xmlLibrary struct {
}

// AnalyzeXmlData анализирует данные в формате xml.
func (l *xmlLibrary) AnalyzeXmlData(data xml) error {
	fmt.Println("xml library analyze xml data")
	return nil
}

// формат json для примера
type json []byte

// jsonLibrary  Неизвестный сервис
type jsonLibrary struct{}

// AnalyzeJsonData анализирует данные в формате JSON.
func (l *jsonLibrary) AnalyzeJsonData(data json) error {
	fmt.Println("json analyzer analyze json data")
	return nil
}

// XmlToJsonAdapter Адаптер оборачивает один из объектов, так что другой объект даже не знает о наличии первого.
type XmlToJsonAdapter struct {
	jsonLibrary *jsonLibrary
}

// AnalyzeXmlData Конвертирует xml данные в json формат и анализирует json формат.
func (a XmlToJsonAdapter) AnalyzeXmlData(data xml) error {
	fmt.Println("adapter convert data to json")
	return a.jsonLibrary.AnalyzeJsonData(json(data))
}
