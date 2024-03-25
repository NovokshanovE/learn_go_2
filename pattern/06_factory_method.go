package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "fmt"

// Интерфейс продукта
type Prod interface {
	Use() string
}

// Конкретный продукт A
type ConcreteProdA struct{}

func (cpa *ConcreteProdA) Use() string {
	return "Используем продукт A"
}

// Конкретный продукт B
type ConcreteProdB struct{}

func (cpb *ConcreteProdB) Use() string {
	return "Используем продукт B"
}

// Интерфейс фабрики
type Factory interface {
	CreateProd() Prod
}

// Конкретная фабрика A
type ConcreteFactoryA struct{}

func (cfa *ConcreteFactoryA) CreateProd() Prod {
	return &ConcreteProdA{}
}

// Конкретная фабрика B
type ConcreteFactoryB struct{}

func (cfb *ConcreteFactoryB) CreateProd() Prod {
	return &ConcreteProdB{}
}

func Factory_method_pattern() {
	factoryA := &ConcreteFactoryA{}
	ProdA := factoryA.CreateProd()
	fmt.Println(ProdA.Use())

	factoryB := &ConcreteFactoryB{}
	ProdB := factoryB.CreateProd()
	fmt.Println(ProdB.Use())
}
