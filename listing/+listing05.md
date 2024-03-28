Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
nil для интерфейсов и nil для указателей - это не одно и то же. Интерфейс считается nil только в том случае, если и его тип, и значение равны nil. В данном случае, тип возвращаемого значения не nil (это *customError), поэтому err не равно nil, даже если значение равно nil.
```
