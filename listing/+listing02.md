Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
Отложенные функции могут считывать и присваивать возвращаемым функциям именованные возвращаемые значения. Поэтому в первом случае будет выводиться 2, а во втором 1. 

```
[defer](https://go.dev/blog/defer-panic-and-recover)

