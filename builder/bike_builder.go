package builder

// BikeBuilder ...
type BikeBuilder struct {
	vehicle Vehicle
}

// SetWheels ...
func (bike *BikeBuilder) SetWheels() BuildProcess {
	bike.vehicle.Wheels = 2
	return bike
}

// SetSeats ...
func (bike *BikeBuilder) SetSeats() BuildProcess {
	bike.vehicle.Seats = 1
	return bike
}

// SetStructure ...
func (bike *BikeBuilder) SetStructure() BuildProcess {
	bike.vehicle.Structure = "bike"
	return bike
}

// Build ...
func (bike *BikeBuilder) Build() Vehicle {
	return bike.vehicle
}
