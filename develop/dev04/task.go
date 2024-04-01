package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
import (
	"fmt"
	"sort"
	"strings"
)

func sortedWords(word string) string {
	runes := []rune(strings.ToLower(word))
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	sortedWord := string(runes)
	return sortedWord
}
func findAnagrams(input []string) map[string](*[]string) {
	anagramMap := make(map[string](*[]string))
	keyMap := make(map[string]string)
	for _, word := range input {
		sWord := sortedWords(word)
		if val, ok := keyMap[sWord]; ok {
			arr := *anagramMap[val]
			arr = append(arr, word)
			anagramMap[val] = &arr
		} else {
			keyMap[sWord] = word
			arr := []string{word}
			anagramMap[word] = &arr

		}

	}

	// Удаление множеств из одного элемента
	for key, value := range anagramMap {
		if len(*value) < 2 {
			delete(anagramMap, key)
		} else {
			v := *value
			sort.Sort(sort.StringSlice(v))

		}
	}

	return anagramMap
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "кто", "яа", "ая"}

	anagrams := findAnagrams(words)

	for key, value := range anagrams {
		fmt.Printf("Множество анаграмм для %s: %v\n", key, *value)
	}
}
