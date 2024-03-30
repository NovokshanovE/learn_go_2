package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func UnpackString(input string) (string, error) {
	var result strings.Builder
	var repeatCount int
	var escape bool
	var char_prev rune
	var first bool = true
	for _, char := range input {
		if unicode.IsDigit(char) && !escape {
			repeatCount = repeatCount*10 + int(char-'0')
			if first {
				return "", errors.New("некорректная строка")
			}

		} else {
			first = false
			if repeatCount > 0 {
				result.WriteString(strings.Repeat(string(char_prev), repeatCount))
				repeatCount = 0
			} else if repeatCount == 0 {
				result.WriteRune(char_prev)
			}
			char_prev = char
			repeatCount = 0
		}
	}

	if repeatCount > 0 {
		result.WriteString(strings.Repeat(string(char_prev), repeatCount))
		repeatCount = 0
	} else if repeatCount == 0 {
		result.WriteRune(char_prev)
	}

	if repeatCount > 0 || escape {
		return "", errors.New("некорректная строка")
	}

	return result.String(), nil
}

func main() {
	res, err := UnpackString("4ihjf488 j")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
