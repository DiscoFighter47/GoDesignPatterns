package absfactory

// CruiseMotorBike ...
type CruiseMotorBike struct{}

// NumWheels ...
func (*CruiseMotorBike) NumWheels() int {
	return 2
}

// NumSeats ...
func (*CruiseMotorBike) NumSeats() int {
	return 1
}

// GetMotorBikeType ...
func (*CruiseMotorBike) GetMotorBikeType() int {
	return CruiseMotorBikeType
}
