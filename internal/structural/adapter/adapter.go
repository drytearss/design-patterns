package adapter

import "fmt"

type Charger interface {
	ChargeFromLightning()
}

type LightningCharger struct{}

func (l *LightningCharger) ChargeFromLightning() {
	fmt.Println("Charge from lightning...")
}

type TypeCCharger struct{}

func (t *TypeCCharger) ChargeFromTypeC() {
	fmt.Println("Charge from Type-C...")
}

type TypeCAdapter struct {
	*TypeCCharger
}

func (a *TypeCAdapter) ChargeFromLightning() {
	a.ChargeFromTypeC()
}
