package builder

// CarBuilder ...
type CarBuilder struct {
	vehicle Vehicle
}

// SetWheels ...
func (car *CarBuilder) SetWheels() BuildProcess {
	car.vehicle.Wheels = 4
	return car
}

// SetSeats ...
func (car *CarBuilder) SetSeats() BuildProcess {
	car.vehicle.Seats = 5
	return car
}

// SetStructure ...
func (car *CarBuilder) SetStructure() BuildProcess {
	car.vehicle.Structure = "car"
	return car
}

// Build ...
func (car *CarBuilder) Build() Vehicle {
	return car.vehicle
}
