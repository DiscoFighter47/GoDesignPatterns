package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBikeBuilder(t *testing.T) {
	manufacture := Manufacture{}
	bikeBuilder := &BikeBuilder{}
	manufacture.SetBuilder(bikeBuilder)
	manufacture.Construct()
	vehicle := bikeBuilder.Build()
	assert.Equal(t, 2, vehicle.Wheels)
	assert.Equal(t, 1, vehicle.Seats)
	assert.Equal(t, "bike", vehicle.Structure)
}
