package numbers

import (
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   NumbersService
}

func (mw LoggingMiddleware) Add(a, b int) (output int, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "Add",
			"input", a, b,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Add(a, b)
	return
}

func (mw LoggingMiddleware) Sub(a, b int) (output int, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "Sub",
			"input", a, b,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Sub(a, b)
	return
}
