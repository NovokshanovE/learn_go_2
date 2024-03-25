package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

import "fmt"

// Пример структуры продукта, который мы хотим построить
type Product struct {
	Part1 string
	Part2 string
}

// Интерфейс строителя
type Builder interface {
	BuildPart1()
	BuildPart2()
	GetProduct() Product
}

// Конкретный строитель
type ConcreteBuilder struct {
	product Product
}

func (cb *ConcreteBuilder) BuildPart1() {
	cb.product.Part1 = "Part1"
}

func (cb *ConcreteBuilder) BuildPart2() {
	cb.product.Part2 = "Part2"
}

func (cb *ConcreteBuilder) GetProduct() Product {
	return cb.product
}

// Директор, который управляет процессом строительства
type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
}

func BuilderPattern() {
	builder := &ConcreteBuilder{}
	director := Director{builder: builder}

	director.Construct()
	product := builder.GetProduct()

	fmt.Printf("Product parts: %s, %s\n", product.Part1, product.Part2)
}
