package absfactory

import "fmt"

const (
	// LuxuryCarType ...
	LuxuryCarType = 1
	// FamilyCarType ...
	FamilyCarType = 2
)

// CarFactory ...
type CarFactory struct{}

// Build ...
func (factory *CarFactory) Build(carType int) (Vehicle, error) {
	switch carType {
	case LuxuryCarType:
		return &LuxuryCar{}, nil
	case FamilyCarType:
		return &FamilyCar{}, nil
	}
	return nil, fmt.Errorf("Vehicle of type %d not recognized", carType)
}
