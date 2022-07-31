package main

import (
	"fmt"
	"github.com/golang-patterns/behavior/visitor/private"
	"github.com/golang-patterns/behavior/visitor/reader"
	"github.com/golang-patterns/behavior/visitor/writer"
)

func main() {
	reader := &reader.Reader{}
	writer := &writer.Writer{}

	console := private.Console{}
	console.ParseFlags()
	fmt.Println(console.GetType())
	console.Accept(reader)
	console.Accept(writer)

	file := private.File{}
	fmt.Println(file.GetType())
	file.Accept(reader)
	file.Accept(writer)
	file.Save()

}
