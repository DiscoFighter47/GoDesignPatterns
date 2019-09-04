package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPaymentMethod(t *testing.T) {
	method, err := GetPaymentMethod("method")
	assert.Error(t, err)
	assert.Nil(t, method)
}
