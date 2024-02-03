package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Facade_pattern

	Структурный паттерн.

	Применимость:

Когда нужно предоставить простой или урезанный интерфейс к сложной подсистеме. Например, мы звоним в магазин и делаем заказ по телефону.
Продавец службы поддержки является нашим фасадом ко всем службам и отделам магазина, предоставляющий нам упрощенный интерфейс к системе
создания заказа, платежной системе и отделу доставки.

	Плюсы:

Изолирует клиентов от компонентов сложной подсистемы

	Минусы:

Фасад рискует быть "божественным объектом", привязанным ко всем классам программы
*/
type Order struct {
}

func (o *Order) NewOrder(userID string) {
	fmt.Printf("Create new order for user %s", userID)
}

func (o *Order) DeleteOrder(orderNumber int) {
	fmt.Printf("Delete order %v from database", orderNumber)
}

type Logistics struct {
}

func (l *Logistics) CalculateNewRoute(location string) {
	fmt.Printf("New route for %s is calculated...", location)
}

func (l *Logistics) CreatingNewOrder(orderNumber int) {
	fmt.Printf("Order %v is added to road-map", orderNumber)
}

type BillingSystem struct {
}

func (b *BillingSystem) MakeNewPayment(amount, orderNumber int) {
	fmt.Printf("Amount %v for order %v added to paymant system", amount, orderNumber)
}

func (b *BillingSystem) MakeNewDebit(amount, orderNumber int) {
	fmt.Printf("Amount %v fororder %v debit from paymant system", amount, orderNumber)
}

type MagazineFacade struct {
	order         *Order
	logistic      *Logistics
	billingSystem *BillingSystem
}

func NewMagazineFacade() *MagazineFacade {
	return &MagazineFacade{
		order:         &Order{},
		logistic:      &Logistics{},
		billingSystem: &BillingSystem{},
	}
}

func (m *MagazineFacade) InteractionWithMagazine(userID, location string, orderNumber, amount int) {
	m.order.DeleteOrder(orderNumber)
	m.order.NewOrder(userID)
	m.logistic.CalculateNewRoute(location)
	m.logistic.CreatingNewOrder(orderNumber)
	m.billingSystem.MakeNewDebit(amount, orderNumber)
	m.billingSystem.MakeNewPayment(amount, orderNumber)
}

func main() {
	facade := NewMagazineFacade()

	userID := "d9e7a184-5d5b-11ea-a62a-3499710062d0"
	location := "119180, Россия, г. Москва ул. Большая Полянка, д. 43, стр. 3."
	orderNumber := 234543
	amount := 499

	facade.InteractionWithMagazine(userID, location, orderNumber, amount)

}
