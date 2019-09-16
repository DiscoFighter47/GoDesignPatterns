package prototype

import "errors"

var whitePrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}
var blackPrototype = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}
var bluePrototype = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

// ShirtCache ...
type ShirtCache struct{}

// GetClone ...
func (s *ShirtCache) GetClone(color int) (ItemInfoGetter, error) {
	switch color {
	case White:
		item := *whitePrototype
		return &item, nil
	case Black:
		item := *blackPrototype
		return &item, nil
	case Blue:
		item := *bluePrototype
		return &item, nil
	}
	return nil, errors.New("Shirt model not recognized")
}
