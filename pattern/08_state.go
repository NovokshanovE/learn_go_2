package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import "fmt"

// Интерфейс состояния
type State interface {
	DoAction(context *Context)
}

// Конкретное состояние A
type ConcreteStateA struct{}

func (csa *ConcreteStateA) DoAction(context *Context) {
	fmt.Println("Выполнение действия в состоянии A")
	context.SetState(&ConcreteStateB{})
}

// Конкретное состояние B
type ConcreteStateB struct{}

func (csb *ConcreteStateB) DoAction(context *Context) {
	fmt.Println("Выполнение действия в состоянии B")
	context.SetState(&ConcreteStateA{})
}

// Контекст
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.DoAction(c)
}

func StagePattern() {
	context := &Context{state: &ConcreteStateA{}}

	context.Request()
	context.Request()
}

/*
Применимость паттерна:

Паттерн "состояние" применяется, когда поведение объекта зависит от его состояния и может изменяться во время выполнения.
Используется, когда в коде присутствует множество условных операторов, зависящих от состояния объекта.
Полезен, когда поведение, связанное с различными состояниями, нужно инкапсулировать и сделать независимым от контекста.

Плюсы и минусы:

Плюсы:

-Инкапсулирует поведение, связанное с каждым состоянием, делая его независимым от контекста.
-Упрощает код за счет удаления множества условных операторов, зависящих от состояния объекта.
-Обеспечивает легкую добавляемость новых состояний.
Минусы:

-Может привести к увеличению количества классов из-за создания отдельных классов для каждого состояния.
-Усложняет код из-за введения дополнительных классов.
*/
