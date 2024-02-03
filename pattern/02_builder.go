package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

	Порождающий паттерн, позволяющий создавать объекты пошагово, используя один и тот же процесс строительства.

	Применимость:
Когда мы хотим избавиться от "телескопического конструктора" - позволяет собирать объекты пошагово, вызывая только те шаги, которые нужны.
Когда наш код должен создавать разные представления какого-то объекта.
Когда нам нужно собирать сложные составные объекты, например, деревья компоновщика - Строитель конструирует объекты пошагово, а не за один проход.
Более того, шаги строительства можно выполнять рекурсивно.

	Плюсы:
 Позволяет создавать продукты пошагово.
 Позволяет использовать один и тот же код для создания различных продуктов.
 Изолирует сложный код сборки продукта от его основной бизнес-логики.
	Минусы:
 Усложняет код программы из-за введения дополнительных классов.
 Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/

// Car Тип продукта который собираем
type Car struct {
	Body   string
	Engine string
	Wheels string
}

// CarBuilder Интерфейс строителя для создания различных компонентов машины
type CarBuilder interface {
	BuildBody()
	BuildEngine()
	BuildWheels()
	GetCar() Car
}

// SportCarBuilder строитель, которы собирает спорт-кары
type SportCarBuilder struct {
	car Car
}

func (s *SportCarBuilder) BuildBody() {
	s.car.Body = "Sport"
}

func (s *SportCarBuilder) BuildEngine() {
	s.car.Engine = "V8"
}

func (s *SportCarBuilder) BuildWheels() {
	s.car.Body = "R20"
}

func (s *SportCarBuilder) GetCar() Car {
	return s.car
}

//SedanBuilder строитель, которы собирает спорт-кары
type SedanBuilder struct {
	car Car
}

func (s *SedanBuilder) BuildBody() {
	s.car.Body = "Sedan"
}

func (s *SedanBuilder) BuildEngine() {
	s.car.Engine = "V4"
}

func (s *SedanBuilder) BuildWheels() {
	s.car.Body = "R17"
}

func (s *SedanBuilder) GetCar() Car {
	return s.car
}

//Director рулит строительством
type Director struct {
	builder CarBuilder
}

func (d *Director) SetBuilder(builder CarBuilder) {
	d.builder = builder
}

func (d *Director) BuildCar() Car {
	d.builder.BuildEngine()
	d.builder.BuildBody()
	d.builder.BuildWheels()
	return d.builder.GetCar()
}

func main() {
	director := Director{}

	sportCarBuilder := &SportCarBuilder{}
	director.SetBuilder(sportCarBuilder)
	Lamborghini := director.BuildCar()
	fmt.Println("New sport-car: ", Lamborghini)

	sedanBuilder := &SedanBuilder{}
	director.SetBuilder(sedanBuilder)
	VAZ2106 := director.BuildCar()
	fmt.Println("New sedan: ", VAZ2106)
}