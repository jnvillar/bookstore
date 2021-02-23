package time

import "time"

type Backend interface {
	Now() time.Time
}

type Factory struct {
	Backend
}

func NewTimeFactory() *Factory {
	return &Factory{newProductionBackend()}
}
