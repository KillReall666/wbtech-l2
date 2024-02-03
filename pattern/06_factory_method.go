package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Фабричный метод — это порождающий паттерн проектирования, который решает проблему создания различных продуктов, без указания конкретных классов продуктов.

	Применимость:
Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.
	Плюсы:
 Избавляет класс от привязки к конкретным классам продуктов.
 Выделяет код производства продуктов в одно место, упрощая поддержку кода.
 Упрощает добавление новых продуктов в программу.
 Реализует принцип открытости/закрытости.
	Минусы:
Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.

P.S.
В Go невозможно реализовать классический вариант паттерна Фабричный метод, поскольу в языке отсутствуют возможности ООП, в том числе классы и наследственность.
Несмотря на это, мы все же можем реализовать базовую версию этого паттерна — Простая фабрика.

P.S.S
Основное отличие паттерна фабрика от паттерна строитель заключается в том, что паттерн фабрика используется для создания объектов определенного типа,
но без конкретной спецификации создаваемых объектов. Это позволяет легко добавлять новые типы объектов без изменения кода самой фабрики.

С другой стороны, паттерн строитель используется для создания объектов с более сложной структурой, где есть несколько шагов и возможны различные варианты конфигурации объекта.
Строитель обеспечивает более гибкое управление процессом создания объекта и позволяет создавать объекты с разными свойствами и параметрами.

Таким образом, основное отличие между этими паттернами заключается в их основной цели: фабрика - для создания объектов определенного типа,
строитель - для создания объектов с более сложной структурой и конфигурацией.
*/

type ISuperHero interface {
	setName(name string)
	setSuperPower(superPower string)
	getName() string
	getSuperPower() string
}

type SuperHero struct {
	name       string
	superPower string
}

func (s *SuperHero) setName(name string) {
	s.name = name
}

func (s *SuperHero) getName() string {
	return s.name
}

func (s *SuperHero) setSuperPower(superPower string) {
	s.superPower = superPower
}

func (s *SuperHero) getSuperPower() string {
	return s.superPower
}

type MarvelSuperHero struct {
	SuperHero
}

func newMarvelSuperHero(name, power string) *SuperHero {
	return &SuperHero{
		name:       name,
		superPower: power,
	}
}

type DCSuperHero struct {
	SuperHero
}

func newDCSuperHero(name, power string) *SuperHero {
	return &SuperHero{
		name:       name,
		superPower: power,
	}
}

func getSuperHero(heroType, name, power string) (*SuperHero, error) {
	if heroType == "marvel" {
		return newMarvelSuperHero(name, power), nil
	}
	if heroType == "dc" {
		return newDCSuperHero(name, power), nil
	}

	return nil, fmt.Errorf("wrong hero type")
}

func printDetails(s ISuperHero) {
	fmt.Printf("Name of hero is: %s", s.getName())
	fmt.Println()
	fmt.Printf("And his super-power is: %d", s.getSuperPower())
	fmt.Println()
}

func main() {
	spiderMan, _ := getSuperHero("marverl", "Spider-Man", "Spider sense")
	batman, _ := getSuperHero("dc", "Batman", "")
	printDetails(spiderMan)
	printDetails(batman)
}
