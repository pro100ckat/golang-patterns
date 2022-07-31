package main

import (
	"github.com/golang-patterns/structural/compositor/compositor"
	"github.com/golang-patterns/structural/compositor/file"
)

func main() {
	file1 := &file.File{Name: "file 1"}
	file2 := &file.File{Name: "file 2"}
	file3 := &file.File{Name: "file 3"}

	compositor1 := compositor.Compositor{Name: "compositor 1"}
	compositor1.AddComponents(file1)
	compositor1.AddComponents(file2)
	compositor1.AddComponents(file3)
	compositor1.Search()

	compositor2 := compositor.Compositor{Name: "compositor 2"}

	compositor2.AddComponents(file1)
	compositor2.AddComponents(file2)
	compositor2.Search()
}
