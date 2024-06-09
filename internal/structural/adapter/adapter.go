package adapter

import "fmt"

// В большинстве случаев интерфейс определяем на стороне потребетиля, а не конкретной реализации
type Charger interface {
	ChargeFromLightning()
}

// Код который писали мы (условно)
type LightningCharger struct{}

func (l *LightningCharger) ChargeFromLightning() {
	fmt.Println("Charge from lightning...")
}

// Код из сторонней библиотеки (условно)
type TypeCCharger struct{}

func (t *TypeCCharger) ChargeFromTypeC() {
	fmt.Println("Charge from Type-C...")
}

// Созданный нами адаптер
type TypeCAdapter struct {
	TypeCCharger *TypeCCharger
}

func (a *TypeCAdapter) ChargeFromLightning() {
	// Тут можно добавить свою логику
	a.TypeCCharger.ChargeFromTypeC()
}
