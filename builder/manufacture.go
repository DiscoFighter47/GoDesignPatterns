package builder

// Manufacture ...
type Manufacture struct {
	process BuildProcess
}

// Construct ...
func (m *Manufacture) Construct() {
	m.process.SetWheels().SetSeats().SetStructure()
}

// SetBuilder ...
func (m *Manufacture) SetBuilder(process BuildProcess) {
	m.process = process
}
