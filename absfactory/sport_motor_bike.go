package absfactory

// SportMotorBike ...
type SportMotorBike struct{}

// NumWheels ...
func (*SportMotorBike) NumWheels() int {
	return 2
}

// NumSeats ...
func (*SportMotorBike) NumSeats() int {
	return 1
}

// GetMotorBikeType ...
func (*SportMotorBike) GetMotorBikeType() int {
	return SportMotorBikeType
}
