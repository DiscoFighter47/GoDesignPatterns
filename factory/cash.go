package factory

import "fmt"

// CashMethod ...
type CashMethod struct{}

// Pay ...
func (method *CashMethod) Pay(amount float64) string {
	return fmt.Sprintf("paid $%.2f using cash card", amount)
}
