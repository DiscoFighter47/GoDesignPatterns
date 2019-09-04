package builder

// BuildProcess ...
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() Vehicle
}
