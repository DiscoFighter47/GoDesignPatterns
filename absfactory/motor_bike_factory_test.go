package absfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMotorbikeFactory(t *testing.T) {
	t.Run("sport bike", func(t *testing.T) {
		factory, err := BuildFactory(MotorBikeFactoryType)
		assert.NoError(t, err)
		vehicle, err := factory.Build(SportMotorBikeType)
		assert.NoError(t, err)
		assert.Equal(t, 2, vehicle.NumWheels())
		assert.Equal(t, 1, vehicle.NumSeats())
		bike, ok := vehicle.(MotorBike)
		assert.True(t, ok)
		assert.Equal(t, SportMotorBikeType, bike.GetMotorBikeType())
	})

	t.Run("cruise bike", func(t *testing.T) {
		factory, err := BuildFactory(MotorBikeFactoryType)
		assert.NoError(t, err)
		vehicle, err := factory.Build(CruiseMotorBikeType)
		assert.NoError(t, err)
		assert.Equal(t, 2, vehicle.NumWheels())
		assert.Equal(t, 1, vehicle.NumSeats())
		bike, ok := vehicle.(MotorBike)
		assert.True(t, ok)
		assert.Equal(t, CruiseMotorBikeType, bike.GetMotorBikeType())
	})
}
