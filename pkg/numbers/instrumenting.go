package numbers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type InstrumentingMiddleware struct {
	addCounter prometheus.Counter
	subCounter prometheus.Counter

	addLatency prometheus.Histogram

	next NumbersService
}

func NewInstrumentingMiddleware() func(NumbersService) NumbersService {
	addcnt := promauto.NewCounter(prometheus.CounterOpts{
		Name: "add_counter",
		Help: "Number of Sub call",
	})
	subcnt := promauto.NewCounter(prometheus.CounterOpts{
		Name: "sub_counter",
		Help: "Number of Sub call",
	})

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":5080", nil)

	return func(ns NumbersService) NumbersService {
		return InstrumentingMiddleware{
			addCounter: addcnt,
			subCounter: subcnt,
		}
	}
}

func (i InstrumentingMiddleware) Add(a, b int) (output int, err error) {
	output, err = i.next.Sub(a, b)
	return
}

func (i InstrumentingMiddleware) Sub(a, b int) (output int, err error) {
	output, err = i.next.Sub(a, b)
	return
}
