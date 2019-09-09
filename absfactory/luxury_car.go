package absfactory

// LuxuryCar ...
type LuxuryCar struct{}

// NumDoors ...
func (*LuxuryCar) NumDoors() int {
	return 2
}

// NumWheels ...
func (*LuxuryCar) NumWheels() int {
	return 4
}

// NumSeats ...
func (*LuxuryCar) NumSeats() int {
	return 2
}
