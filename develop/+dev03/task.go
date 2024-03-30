package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileName := flag.String("file", "", "название файла для сортировки")
	k := flag.Int("k", 0, "указание колонки для сортировки")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	if *fileName == "" {
		fmt.Println("Необходимо указать файл для сортировки")
		os.Exit(1)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Сортировка
	sort.Slice(lines, func(i, j int) bool {
		valI := getColumnValue(lines[i], *k)
		valJ := getColumnValue(lines[j], *k)

		if *n {
			// Сортировка по числовому значению
			numI, _ := strconv.Atoi(valI)
			numJ, _ := strconv.Atoi(valJ)
			if *r {
				return numI > numJ
			}
			return numI < numJ
		} else {
			// Сортировка по строковому значению
			if *r {
				return valI > valJ
			}
			return valI < valJ
		}
	})

	// Удаление дубликатов, если указан ключ -u
	if *u {
		lines = removeDuplicates(lines)
	}

	// Запись отсортированных строк в файл
	err = os.WriteFile(*fileName, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func getColumnValue(line string, index int) string {
	columns := strings.Fields(line)
	if index < len(columns) {
		return columns[index]
	}
	return ""
}

func removeDuplicates(lines []string) []string {
	var uniqueLines []string
	seen := make(map[string]bool)
	for _, line := range lines {
		if !seen[line] {
			uniqueLines = append(uniqueLines, line)
			seen[line] = true
		}
	}
	return uniqueLines
}
