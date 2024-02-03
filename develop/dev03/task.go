package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	filePath := "input.txt" // Путь к входному файлу
	outputPath := "output.txt" // Путь к выходному файлу

	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Парсинг аргументов командной строки
	args := os.Args[1:]
	column := 0
	sortNumeric := false
	reverse := false
	unique := false

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-k":
			if i+1 < len(args) {
				column, _ = strconv.Atoi(args[i+1])
				i++
			}
		case "-n":
			sortNumeric = true
		case "-r":
			reverse = true
		case "-u":
			unique = true
		}
	}

	// Сортировка строк
	sort.SliceStable(lines, func(i, j int) bool {
		return compare(lines[i], lines[j], column, sortNumeric) < 0
	})

	// Обратный порядок сортировки, если указан флаг -r
	if reverse {
		reverseSlice(lines)
	}

	// Удаление повторяющихся строк, если указан флаг -u
	if unique {
		lines = removeDuplicates(lines)
	}

	// Запись отсортированных строк в выходной файл
	err = writeLines(lines, outputPath)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Файл успешно отсортирован.")
}

// Функция для чтения строк из файла
func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// Функция для записи строк в файл
func writeLines(lines []string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}

// Функция для сравнения строк
func compare(a, b string, column int, sortNumeric bool) int {
	if column > 0 {
		wordsA := strings.Fields(a)
		wordsB := strings.Fields(b)

		if column <= len(wordsA) && column <= len(wordsB) {
			a = wordsA[column-1]
			b = wordsB[column-1]
		}
	}

	if sortNumeric {
		numA, errA := strconv.Atoi(a)
		numB, errB := strconv.Atoi(b)

		if errA == nil && errB == nil {
			if numA < numB {
				return -1
			} else if numA > numB {
				return 1
			} else {
				return 0
			}
		}
	}

	return strings.Compare(a, b)
}

// Функция для обратного порядка сортировки
func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Функция для удаления повторяющихся строк
func removeDuplicates(slice []string) []string {
	uniqueSet := make(map[string]bool)
	result := make([]string, 0)

	for _, item := range slice {
		if !uniqueSet[item] {
			uniqueSet[item] = true
			result = append(result, item)
		}
	}

	return result
}
