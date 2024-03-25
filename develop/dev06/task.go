package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Определение флагов
	f := flag.String("f", "", "выбрать поля (колонки)")
	d := flag.String("d", "\t", "использовать другой разделитель")
	s := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	// Преобразование строки с номерами полей в слайс
	var fields []int
	if *f != "" {
		for _, field := range strings.Split(*f, ",") {
			fields = append(fields, atoi(field)-1) // -1, потому что индексация начинается с 0
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Если ключ -s активен и в строке нет разделителя, пропускаем строку
		if *s && !strings.Contains(line, *d) {
			continue
		}

		columns := strings.Split(line, *d)
		for _, field := range fields {
			if field < len(columns) {
				fmt.Print(columns[field] + " ")
			}
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка преобразования строки в число:", err)
		os.Exit(1)
	}
	return n
}
