package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Состояние — это поведенческий паттерн, позволяющий динамически изменять поведение объекта при смене его состояния.

	Применимость:
Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется.
Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.
Когда вы сознательно используете табличную машину состояний, построенную на условных операторах, но вынуждены мириться с дублированием кода для похожих состояний и переходов.
	Плюсы:
 Избавляет от множества больших условных операторов машины состояний.
 Концентрирует в одном месте код, связанный с определённым состоянием.
 Упрощает код контекста.
	Минусы:
Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type State int

const (
	Newbie       State = 0 //жёсткое ограничение
	FirstCourse  State = 1
	SecondCourse State = 2
	ThirdCourse  State = 3
	Graduate     State = 4
)

type Student struct {
	Name  string
	State State
	Exam  bool
}

func NewStudent(name string) Student {
	return Student{
		Name:  name,
		State: Newbie,
	}
}

func (s *Student) CongratsGraduate() error {
	if s.State != Graduate {
		return fmt.Errorf("current state %d: need state to graduate: %d ", s.State, Graduate)
	}
	fmt.Println(fmt.Sprintf("Congrat to %s with graduation!", s.Name))

	return nil
}

func (s *Student) GoodExam() {
	s.Exam = true
}
func (s *Student) BadExam() {
	s.Exam = false
}

func (s *Student) PromoteToThirdCourse() error {
	if s.State != SecondCourse {
		return fmt.Errorf("current state %d: need state to promote: %d ", s.State, SecondCourse)
	}
	if !s.Exam {
		return fmt.Errorf("unsuccessed examine")
	}
	s.State = ThirdCourse
	s.BadExam()
	return nil

}

func main() {
	student1 := NewStudent("Ragnar")
	student1.State = SecondCourse
	student1.GoodExam()
	err := student1.PromoteToThirdCourse()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(student1.State)

}