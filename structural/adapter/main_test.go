package main

import (
	"github.com/golang-patterns/structural/adapter/adapter"
	"github.com/golang-patterns/structural/adapter/jsonlib"
	"github.com/golang-patterns/structural/adapter/xmllib"
	"testing"
)

func TestAdapter(t *testing.T) {
	app := app{}
	xmlAnalyzer := xmllib.Analyzer{}
	xmlData := xmllib.XML{}
	err := app.AnalyzeXmlData(&xmlAnalyzer, xmlData)
	if err != nil {
		t.Error(err)
	}
	jsonAdapter := &adapter.Adapter{
		JsonAnalyzer: &jsonlib.Analyzer{},
	}
	err = app.AnalyzeXmlData(jsonAdapter, xmlData)
	if err != nil {
		t.Error(err)
	}
}
