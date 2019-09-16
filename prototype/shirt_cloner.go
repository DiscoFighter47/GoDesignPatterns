package prototype

const (
	// White ...
	White = 1
	// Black ...
	Black = 2
	// Blue ...
	Blue = 3
)

// ShirtCloner ...
type ShirtCloner interface {
	GetClone(color int) (ItemInfoGetter, error)
}

// GetShirtCloner ...
func GetShirtCloner() ShirtCloner {
	return &ShirtCache{}
}
