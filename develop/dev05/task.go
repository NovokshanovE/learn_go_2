package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		A = flag.Int("A", 0, "print +N lines after match")
		B = flag.Int("B", 0, "print +N lines before match")
		C = flag.Int("C", 0, "print ±N lines around match")
		c = flag.Bool("c", false, "print the count of matching lines")
		i = flag.Bool("i", false, "ignore case")
		v = flag.Bool("v", false, "select non-matching lines")
		F = flag.Bool("F", false, "exact match of the string, not a pattern")
		n = flag.Bool("n", false, "print line number")
	)
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Usage: grep [options] pattern")
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	if *i {
		pattern = strings.ToLower(pattern)
	}

	count := 0
	lineNum := 0
	printLineNum := func() {
		if *n {
			fmt.Printf("%d:", lineNum)
		}
	}

	matches := make([]bool, 0)
	lines := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		lineToMatch := line
		if *i {
			lineToMatch = strings.ToLower(line)
		}

		match := strings.Contains(lineToMatch, pattern)
		if *F {
			match = lineToMatch == pattern
		}
		if *v {
			match = !match
		}

		matches = append(matches, match)
		lines = append(lines, line)

		if match && !*c {
			count++
		}
	}

	if *c {
		fmt.Println(count)
		return
	}

	for index, _ := range lines {
		if matches[index] {
			for j := index - *B; j <= index+*A; j++ {
				if j >= 0 && j < len(lines) {
					printLineNum()
					fmt.Println(lines[j])
				}
			}
			if *C > 0 {
				for j := index - *C; j <= index+*C; j++ {
					if j >= 0 && j < len(lines) {
						printLineNum()
						fmt.Println(lines[j])
					}
				}
			}
		}
	}
}
