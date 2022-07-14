package main

// director Директор определяет порядок вызова строительных шагов для производства той или иной конфигурации продуктов.
type director struct {
	builder Builder
}

// newDirector Конструктор директора. Обычно Клиент подаёт в конструктор директора уже готовый объект-строитель, и в дальнейшем данный директор использует только его.
func newDirector(b Builder) *director {
	return &director{
		builder: b,
	}
}

// setBuilder Другой вариант, когда клиент передаёт строителя через параметр строительного метода директора.
// В этом случае можно каждый раз применять разных строителей для производства различных представлений объектов.
func (d *director) setBuilder(b Builder) {
	d.builder = b
}

// Build порядок вызова строительных шагов для производства той или иной конфигурации продуктов.
func (d *director) Build() *Product {
	d.builder.BuildField1()
	d.builder.BuildField2()
	return d.builder.BuildProduct()
}
