package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

*/

func main() {

	buf := bufio.NewScanner(os.Stdin)
	fmt.Print("Shell: ")
	for buf.Scan() {
		InputLine := buf.Text()
		Shell(InputLine)

		// fmt.Println("Для выхода используйте ctrl+c.")
		fmt.Print("Shell: ")
	}
}

func Shell(inputLine string) {

	conveyor := strings.Split(inputLine, "|")

	if strings.Contains(inputLine, "&") {
		conveyor[len(conveyor)-1] = strings.ReplaceAll(conveyor[len(conveyor)-1], "&", "") // уберем & в конце
		go doCommands(conveyor)

		return
	}

	doCommands(conveyor)

}

func doCommands(conveyor []string) {

	cmds := make([]*exec.Cmd, 0, len(conveyor))

	for _, v := range conveyor {
		v = strings.TrimSpace(v)
		command := strings.Split(v, " ")

		cmds = append(cmds, exec.Command(command[0], command[1:]...))
	}

	if len(cmds) == 1 { // если pipes не используются
		cmds[0].Stdin = os.Stdin
		cmds[0].Stdout = os.Stdout
		cmds[0].Run()
		return
	}

	output, stdErr, err := Pipeline(cmds)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, output)
	io.Copy(os.Stderr, stdErr)

	fmt.Print("Shell: ")
}

func Pipeline(cmds []*exec.Cmd) (*bytes.Buffer, *bytes.Buffer, error) {

	var output bytes.Buffer
	var stderr bytes.Buffer

	cmds[0].Stdin = os.Stdin

	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {
		var err error
		cmds[i+1].Stdin, err = cmd.StdoutPipe()
		if err != nil {
			return nil, nil, err
		}

		cmd.Stderr = &stderr
	}

	cmds[last].Stdout, cmds[last].Stderr = &output, &stderr

	for _, cmd := range cmds {
		err := cmd.Start()
		if err != nil {
			return nil, nil, err
		}
	}

	for _, cmd := range cmds {
		cmd.Wait()
	}

	return &output, &stderr, nil
}