package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		args := strings.Fields(input)
		cmd := args[0]
		switch cmd {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Необходимо указать аргумент для cd")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Ошибка при смене директории:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка при получении текущей директории:", err)
			} else {
				fmt.Println(dir)
			}
		case "echo":
			if len(args) < 2 {
				fmt.Println("Необходимо указать аргумент для echo")
				continue
			}
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Необходимо указать процесс для завершения")
				continue
			}
			out, err := exec.Command("kill", args[1]).Output()
			if err != nil {
				fmt.Println("Ошибка при завершении процесса:", err)
			} else {
				fmt.Println(string(out))
			}
		case "ps":
			out, err := exec.Command("ps").Output()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды ps:", err)
			} else {
				fmt.Println(string(out))
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды:", err)
			}
		}
	}
}
