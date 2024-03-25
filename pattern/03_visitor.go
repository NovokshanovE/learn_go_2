package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import "fmt"

// Пример интерфейса элемента, который будет принимать посетителя
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент
type ConcreteElement struct {
	name string
}

func (ce *ConcreteElement) Accept(visitor Visitor) {
	visitor.VisitConcreteElement(ce)
}

// Интерфейс посетителя
type Visitor interface {
	VisitConcreteElement(element *ConcreteElement)
}

// Конкретный посетитель
type ConcreteVisitor struct {
	visitorName string
}

func (cv *ConcreteVisitor) VisitConcreteElement(element *ConcreteElement) {
	fmt.Printf("%s посетил элемент %s\n", cv.visitorName, element.name)
}

func VisitorPattern() {
	element := &ConcreteElement{name: "ЭлементA"}
	visitor1 := &ConcreteVisitor{visitorName: "Посетитель1"}
	visitor2 := &ConcreteVisitor{visitorName: "Посетитель2"}

	element.Accept(visitor1)
	element.Accept(visitor2)
}
