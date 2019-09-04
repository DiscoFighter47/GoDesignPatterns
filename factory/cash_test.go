package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCashMethod(t *testing.T) {
	method, err := GetPaymentMethod(Cash)
	assert.NoError(t, err)
	assert.Equal(t, &CashMethod{}, method)
	assert.Equal(t, "paid $10.25 using cash card", method.Pay(10.25))
}
