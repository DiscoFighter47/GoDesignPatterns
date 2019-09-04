package factory

import (
	"fmt"
)

// Method ...
type Method string

const (
	// Cash ...
	Cash = "cash"
	// DebitCard ...
	DebitCard = "debit"
)

// PaymentMethod ...
type PaymentMethod interface {
	Pay(float64) string
}

// GetPaymentMethod ...
func GetPaymentMethod(method Method) (PaymentMethod, error) {
	switch method {
	case Cash:
		return &CashMethod{}, nil
	case DebitCard:
		return &DebitMethod{}, nil
	default:
		return nil, fmt.Errorf("Payment method %v not recognized", method)
	}
}
