package numbers

import "github.com/go-kit/kit/metrics/prometheus"

type InstrumentingMiddleware struct {
	addCounter prometheus.Counter

	next NumbersService
}

func NewInstrumentingMiddleware() func(NumbersService) NumbersService {
	return func(ns NumbersService) NumbersService {
		return InstrumentingMiddleware{}
	}
}

func (InstrumentingMiddleware) Add(a, b int) (int, error) {

}

func (InstrumentingMiddleware) Sub(a, b int) (int, error) {

}
