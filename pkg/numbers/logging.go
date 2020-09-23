package numbers

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   NumbersService
}

func (mw loggingMiddleware) Add(a, b int) (output int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Add",
			"input", a, b,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Add(a, b)
	return
}

func (mw loggingMiddleware) Sub(a, b int) (output int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Sub",
			"input", a, b,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Sub(a, b)
	return
}
