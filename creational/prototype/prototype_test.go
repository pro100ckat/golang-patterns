package main

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	prototype1 := &ConcretePrototype{field: "1"}
	clone1 := prototype1.Clone()
	prototype2 := &ConcretePrototype2{field: "1"}
	clone2 := prototype2.Clone()

	if prototype1.GetField() != clone1.GetField() || prototype2.field != clone2.GetField() {
		t.Error("Objects are not equal!\n")
	}
}
