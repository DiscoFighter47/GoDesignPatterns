package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	t.Run("count to 5", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			assert.Equal(t, int64(i), GetInstance().AddOne())
		}
	})
}

func TestGetInstance(t *testing.T) {
	assert.NotNil(t, GetInstance())
	assert.Same(t, GetInstance(), GetInstance())
}
