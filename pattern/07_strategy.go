package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Стратегия — это поведенческий паттерн, выносит набор алгоритмов в собственные классы и делает их взаимозаменимыми.
Другие объекты содержат ссылку на объект-стратегию и делегируют ей работу. Программа может подменить этот объект другим, если требуется иной способ решения задачи.

	Применимость:
Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого оператора представляет собой вариацию алгоритма.

	Плюсы:
Горячая замена алгоритмов на лету.
Изолирует код и данные алгоритмов от остальных классов.
Уход от наследования к делегированию.
Реализует принцип открытости/закрытости.
	Минусы:
Усложняет программу за счёт дополнительных классов.
Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

type SortStrategy interface {
	Sort(data []int)
}

type BubbleSortStrategy struct {
}

func (b BubbleSortStrategy) Sort(data []int) {
	fmt.Println("BubbleSort processed...", data)
	//Реализация сортировки пузырьком
}

type QuickSortStrategy struct {
}

func (strategy QuickSortStrategy) Sort(data []int) {
	fmt.Println("QuickSort processed...:", data)
	// Реализация быстрой сортировки
}

type SortContext struct {
	strategy SortStrategy
}

func (context SortContext) Sort(data []int) {
	context.strategy.Sort(data)
}

func main() {
	// Создание стратегий
	bubbleSortStrategy := BubbleSortStrategy{}
	quickSortStrategy := QuickSortStrategy{}

	// Создание контекста с пузырьковой сортировкой
	context := SortContext{strategy: bubbleSortStrategy}
	data := []int{4, 2, 7, 1, 5}
	context.Sort(data)

	// Изменение стратегии на быструю сортировку
	context.strategy = quickSortStrategy
	context.Sort(data)
}
