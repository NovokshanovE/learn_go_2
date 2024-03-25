package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import "fmt"

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// Базовая реализация обработчика
type BaseHandler struct {
	nextHandler Handler
}

func (bh *BaseHandler) SetNext(handler Handler) {
	bh.nextHandler = handler
}

// Конкретный обработчик
type ConcreteHandlerA struct {
	BaseHandler
}

func (cha *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA обработал запрос")
	} else if cha.nextHandler != nil {
		cha.nextHandler.Handle(request)
	}
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (chb *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB обработал запрос")
	} else if chb.nextHandler != nil {
		chb.nextHandler.Handle(request)
	}
}

func COfR() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	handlerA.Handle("A")
	handlerA.Handle("B")
	handlerA.Handle("C")
}
