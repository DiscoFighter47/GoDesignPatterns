package singleton

// Counter ...
type Counter interface {
	AddOne() int64
}

type counter struct {
	count int64
}

func (c *counter) AddOne() int64 {
	c.count++
	return c.count
}

var instance *counter

// GetInstance ...
func GetInstance() Counter {
	if instance == nil {
		instance = &counter{}
	}
	return instance
}
