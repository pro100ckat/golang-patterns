package private

// Пакет содержит иерархию типов File и Console

import "fmt"

// File - наш объект, поведение которого мы будем расширять через посетителя.
type File struct{}

// Save - метод который уже что-то делает в этом пакете.
func (f *File) Save() {
	fmt.Println("file saved")
}

// GetType - еще один метод, который что-то делает.
func (f *File) GetType() string {
	return "i am file"
}

// Accept - принимает некоторое поведение, которое не хочется или нельзя писать в пакете private.
func (f *File) Accept(visitor Visitor) {
	visitor.VisitFile(f)
}
