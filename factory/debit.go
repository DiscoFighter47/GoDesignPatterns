package factory

import "fmt"

// DebitMethod ...
type DebitMethod struct{}

// Pay ...
func (method *DebitMethod) Pay(amount float64) string {
	return fmt.Sprintf("paid $%.2f using debit card", amount)
}
