package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebitMethod(t *testing.T) {
	method, err := GetPaymentMethod(DebitCard)
	assert.NoError(t, err)
	assert.Equal(t, &DebitMethod{}, method)
	assert.Equal(t, "paid $15.75 using debit card", method.Pay(15.75))
}
