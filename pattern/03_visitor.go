package pattern

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Посетитель — поведенческий паттерн, который позволяет добавлять в программу новые операции, не изменяя классы объектов, над которыми эти операции могут выполняться.

	Применимость:
Когда нужно выполнить операцию над всеми элементами сложной структуры объектов, например деревом.
Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции, но мы не хотим «засорять» классы такими операциями.
Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.

	Плюсы:
Упрощает добавление операций, работающих со сложными структурами объектов.
Объединяет родственные операции в одном классе.
Посетитель может накапливать состояние при обходе структуры элементов.
	Минусы:
Паттерн не оправдан, если иерархия элементов часто меняется.
Может привести к нарушению инкапсуляции элементов.
*/

type VisitorShape interface {
	VisitRectangle(rectangle *Rectangle)
	VisitCircle(circle *Circle)
}

type VisitorShapeToPrint struct {
}

func (v *VisitorShapeToPrint) VisitRectangle(rectangle *Rectangle) {
	fmt.Printf("Visitor was here: %+v\n", *rectangle)
}

func (v *VisitorShapeToPrint) VisitCircle(circle *Circle) {
	fmt.Printf("Visitor vas here: %+v\n", *circle)
}

type Rectangle struct {
	width  float32
	height float32
}

type Circle struct {
	radius float32
}

type Shapes struct {
	rectangle Rectangle
	circle    Circle
}

func (s *Shapes) Visit(v VisitorShape) {
	v.VisitRectangle(&s.rectangle)
	v.VisitCircle(&s.circle)
}

func NewShapes() *Shapes {
	s := new(Shapes)
	rand.Seed(time.Now().UnixNano())
	s.rectangle.width = rand.Float32() * 5
	s.rectangle.height = rand.Float32() * 7
	s.circle.radius = rand.Float32() * 9
	return s
}

func main() {
	shapes := NewShapes()
	visitorShape := VisitorShapeToPrint{}
	shapes.Visit(&visitorShape)

}
