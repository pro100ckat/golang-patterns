package main

import "testing"

func TestAdapter(t *testing.T) {
	client := client{}
	xmlLibrary := xmlLibrary{}
	xmlData := xml{}
	err := client.AnalyzeXmlData(&xmlLibrary, xmlData)
	if err != nil {
		t.Error(err)
	}
	adapter := &XmlToJsonAdapter{
		jsonLibrary: &jsonLibrary{},
	}
	err = client.AnalyzeXmlData(adapter, xmlData)
	if err != nil {
		t.Error(err)
	}
}
