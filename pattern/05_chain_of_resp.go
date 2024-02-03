package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	Цепочка обязанностей — это поведенческий паттерн, позволяющий передавать запрос по цепочке потенциальных обработчиков, пока один из них не обработает запрос.

	Применимость:
Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
Когда набор объектов, способных обработать запрос, должен задаваться динамически.
	Плюсы:
 Уменьшает зависимость между клиентом и обработчиками.
 Реализует принцип единственной обязанности.
 Реализует принцип открытости/закрытости.
	Минусы:
 Запрос может остаться никем не обработанным.

*/

type Handler interface {
	Next(handler Handler)
	RequestHandle(request string)
}

type MainHandler struct {
	next Handler
}

func (m *MainHandler) Next(handler Handler) {
	m.next = handler
}

func (m *MainHandler) RequestHandle(request string) {
	if m.next != nil {
		m.next.RequestHandle(request)
	}
}

type AuthenticationHandler struct {
	MainHandler
}

func (a *AuthenticationHandler) RequestHandle(request string) {
	if strings.HasPrefix(request, "auth ") {
		fmt.Println("AuthorizationHandler: обработка запроса", request)
	} else {
		a.MainHandler.RequestHandle(request)
	}
}

type AuthorizationHandler struct {
	MainHandler
}

func (a *AuthorizationHandler) RequestHandle(request string) {
	if strings.HasPrefix(request, "author ") {
		fmt.Println("ValidationHandler: обработка запроса", request)
	} else {
		a.MainHandler.RequestHandle(request)
	}
}

type ValidationHandler struct {
	MainHandler
}

func (v *ValidationHandler) RequestHandle(request string) {
	if strings.HasPrefix(request, "validate ") {
		fmt.Println("ValidationHandler: обработка запроса", request)
	} else {
		v.MainHandler.RequestHandle(request)
	}
}

func main() {
	// Создание обработчиков
	authenticationHandler := &AuthenticationHandler{}
	authorizationHandler := &AuthorizationHandler{}
	validationHandler := &ValidationHandler{}

	// Настройка цепочки вызовов
	authenticationHandler.Next(authorizationHandler)
	authorizationHandler.Next(validationHandler)

	// Обработка запросов
	authenticationHandler.RequestHandle("auth user:password")
	authenticationHandler.RequestHandle("authz user=admin")
	authenticationHandler.RequestHandle("validate data")
	authenticationHandler.RequestHandle("other request")
}
