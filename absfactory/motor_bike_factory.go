package absfactory

import "fmt"

const (
	// SportMotorBikeType ...
	SportMotorBikeType = 1
	// CruiseMotorBikeType ...
	CruiseMotorBikeType = 2
)

// MotorBikeFactory ...
type MotorBikeFactory struct{}

// Build ...
func (*MotorBikeFactory) Build(bikeType int) (Vehicle, error) {
	switch bikeType {
	case SportMotorBikeType:
		return &SportMotorBike{}, nil
	case CruiseMotorBikeType:
		return &CruiseMotorBike{}, nil
	}
	return nil, fmt.Errorf("Vehicle of type %d not recognized", bikeType)
}
