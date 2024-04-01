package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Вы не ввели аргумент")
		return
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Создание файла для сохранения содержимого страницы
	fileName := getFileName(url)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Копирование содержимого страницы в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Ошибка при сохранении содержимого страницы:", err)
		return
	}

	fmt.Println("Страница успешно загружена и сохранена в файл", fileName)
}

// Функция для получения имени файла из URL
func getFileName(url string) string {
	if url[len(url)-1] == '/' {
		parts := strings.Split(url, "/")
		fileName := parts[len(parts)-2]
		return fileName
	}
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1]
	return fileName
}
