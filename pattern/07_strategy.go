package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import "fmt"

// Интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Конкретная стратегия для оплаты кредитной картой
type CreditCardPayment struct{}

func (ccp *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f с помощью кредитной карты", amount)
}

// Конкретная стратегия для оплаты через PayPal
type PayPalPayment struct{}

func (ppp *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f через PayPal", amount)
}

// Контекст, использующий стратегию
type PaymentContext struct {
	strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ExecutePayment(amount float64) string {
	return pc.strategy.Pay(amount)
}

func Strategy_pattern() {
	creditCard := &CreditCardPayment{}
	payPal := &PayPalPayment{}

	context := &PaymentContext{}

	context.SetStrategy(creditCard)
	fmt.Println(context.ExecutePayment(100.0))

	context.SetStrategy(payPal)
	fmt.Println(context.ExecutePayment(50.0))
}

/*

Применимость паттерна:


Паттерн "стратегия" применяется, когда необходимо выбирать алгоритм из семейства алгоритмов во время выполнения.
Используется, когда есть несколько вариантов решения задачи и они должны быть легко расширяемы и заменяемы.
Полезен, когда необходимо изолировать код, который меняется, от остальной части программы.

Плюсы и минусы:


Плюсы:

-Позволяет выбирать алгоритм из семейства алгоритмов во время выполнения.
-Обеспечивает легкую замену и добавление новых алгоритмов.
-Изолирует код, который меняется, от остальной части программы.
Минусы:

-Может привести к увеличению количества классов из-за создания отдельных стратегий для каждого алгоритма.
-Усложняет код из-за введения дополнительных классов.


*/
