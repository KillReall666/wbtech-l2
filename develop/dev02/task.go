package main

import (
	"errors"
	"strings"
	"unicode"
)

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

var ErrIncorrectString = errors.New("incorrect string")

func unpackingString(s string) (string, error) {
	var prevRune rune
	var sb strings.Builder

	for i, r := range s {
		isDigit := unicode.IsDigit(r)

		if (isDigit && i == 0) || (isDigit && unicode.IsDigit(prevRune)) {
			return s, ErrIncorrectString
		}
		if !isDigit {
			sb.WriteRune(r)
		}
		if isDigit {
			repeat := int(r - '1')
			str := strings.Repeat(string(prevRune), repeat)
			sb.WriteString(str)
		}
		prevRune = r
	}
	return sb.String(), nil
}

