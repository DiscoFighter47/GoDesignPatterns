package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShirtCloner(t *testing.T) {
	shirtCache := GetShirtCloner()
	assert.NotNil(t, shirtCache)

	shirt1, err := shirtCache.GetClone(White)
	assert.NoError(t, err)
	assert.NotEqual(t, &whitePrototype, &shirt1)
}
