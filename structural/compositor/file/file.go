package file

import "fmt"

// File - конкретная реализация Component.
type File struct {
	Name string
}

// Search - функция которая что-то делает для конкретного компонента.
func (f *File) Search() {
	fmt.Println("search in file", f.Name)
}
