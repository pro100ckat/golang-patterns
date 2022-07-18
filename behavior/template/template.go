package main

// Шаблонный метод
type foo interface {
	some()
}

type foostr struct {
	foo foo
}

func (f *foostr) Foo() {
	// Шаги алгоритма
}

// Другой вариант вызова
//type foostr struct {
//
//}

//func (f *foostr) Foo(foo foo)  {
//	foo.some()
//}

//  Конкретная реализация. Реализует шаги алгоритма. Реализуеет foo interface
type A struct {
	f *foostr
}

func (a *A) some() {
}

//  Конкретная реализация. Реализует шаги алгоритма. Реализуеет foo interface
type B struct {
	f *foostr
}

func (b *B) some() {
}

func main() {

}
