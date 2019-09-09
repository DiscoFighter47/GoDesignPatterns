package absfactory

import (
	"fmt"
)

const (
	// CarFactoryType ...
	CarFactoryType = 1
	// MotorBikeFactoryType ...
	MotorBikeFactoryType = 2
)

// VehicleFactory ...
type VehicleFactory interface {
	Build(int) (Vehicle, error)
}

// BuildFactory ...
func BuildFactory(factoryType int) (VehicleFactory, error) {
	switch factoryType {
	case CarFactoryType:
		return &CarFactory{}, nil
	case MotorBikeFactoryType:
		return &MotorBikeFactory{}, nil
	}
	return nil, fmt.Errorf("Factory with id %d not recognized", factoryType)
}
