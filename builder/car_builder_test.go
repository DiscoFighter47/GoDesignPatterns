package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarBuilder(t *testing.T) {
	manufacture := Manufacture{}
	carBuilder := &CarBuilder{}
	manufacture.SetBuilder(carBuilder)
	manufacture.Construct()
	vehicle := carBuilder.Build()
	assert.Equal(t, 4, vehicle.Wheels)
	assert.Equal(t, 5, vehicle.Seats)
	assert.Equal(t, "car", vehicle.Structure)
}
