package private

// Visitor - добавляем новое поведение в пакет.
type Visitor interface {
	VisitConsole(console *Console)
	VisitFile(console *File)
}
