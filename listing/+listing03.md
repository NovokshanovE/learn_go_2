Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

```
Функция Foo() возвращает nil, но тип возвращаемого значения - это указатель на os.PathError.
nil для интерфейсов и nil для указателей - это не одно и то же => будет выведено false.