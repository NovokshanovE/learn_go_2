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
	"strings"
)

func main() {
	var (
		fields    = flag.String("f", "", "choose fields (columns)")
		delimiter = flag.String("d", "\t", "use another delimiter")
		separated = flag.Bool("s", false, "only lines with delimiter")
	)
	flag.Parse()

	if *fields == "" {
		fmt.Fprintln(os.Stderr, "Usage: cut -f <fields> [-d <delimiter>] [-s]")
		os.Exit(1)
	}

	fieldIndexes := parseFields(*fields)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Если установлен флаг -s и разделитель не найден, пропускаем строку
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		columns := strings.Split(line, *delimiter)
		var output []string
		for _, index := range fieldIndexes {
			if index < len(columns) {
				output = append(output, columns[index])
			}
		}

		fmt.Println(strings.Join(output, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}

// parseFields преобразует строку полей в срез индексов
func parseFields(fieldStr string) []int {
	var fields []int
	for _, f := range strings.Split(fieldStr, ",") {
		var fieldIndex int
		fmt.Sscanf(f, "%d", &fieldIndex)
		// Уменьшаем индекс на 1, так как пользователь вводит поля начиная с 1, а индексы в Go начинаются с 0
		fields = append(fields, fieldIndex-1)
	}
	return fields
}
