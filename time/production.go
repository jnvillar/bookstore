package time

import "time"

type productionBackend struct {
}

func (p *productionBackend) Now() time.Time {
	return time.Now()
}

func newProductionBackend() Backend {
	return &productionBackend{}
}
