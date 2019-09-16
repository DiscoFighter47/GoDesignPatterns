package prototype

import "fmt"

// ShirtColor ...
type ShirtColor byte

// Shirt ...
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// GetInfo ...
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}
