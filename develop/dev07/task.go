package main

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

import (
	"fmt"
	"time"
)

// or функция объединяет один или более каналов в один.
// Она закрывается, когда закрывается любой из составляющих каналов.
var or func(channels ...<-chan interface{}) <-chan interface{}

func init() {
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			// В случае отсутствия каналов возвращаем закрытый канал.
			c := make(chan interface{})
			close(c)
			return c
		case 1:
			// Если канал один, просто возвращаем его.
			return channels[0]
		default:
			// Создаем новый канал, который будет закрыт при закрытии любого из входящих каналов.
			orDone := make(chan interface{})
			go func() {
				defer close(orDone)
				switch len(channels) {
				case 2:
					select {
					case <-channels[0]:
					case <-channels[1]:
					}
				default:
					select {
					case <-channels[0]:
					case <-channels[1]:
					case <-channels[2]:
					case <-or(append(channels[3:], orDone)...):
					}
				}
			}()
			return orDone
		}
	}
}

// sig создает канал, который закрывается через указанное время.
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}
