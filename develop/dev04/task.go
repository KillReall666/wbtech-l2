package main

import (
	"fmt"
	"slices"
	"strings"
)

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

func SetsOfAnagrams(str []string) *map[string][]string {
	out := make(map[string][]string)

	tempHashes := makeMap(str)

	for _, anagrams := range tempHashes {
		if len(anagrams) > 1 {
			firstWord := anagrams[0]
			out[firstWord] = append([]string{}, anagrams...)
		}
	}
	return &out
}

func makeMap(slice []string) map[string][]string {
	out := make(map[string][]string)
	for i := range slice {
		str := strings.ToLower(slice[i])
		runes := []rune(str)
		slices.Sort(runes)
		hash := string(runes)
		out[hash] = append(out[hash], str)
	}
	return out
}


func main() {
	str := []string{"пЯтак", "пяТка", "Тяпка", "лИсток", "сЛиток", "стОлик", "беРеза"}
	fmt.Println(SetsOfAnagrams(str))
}
