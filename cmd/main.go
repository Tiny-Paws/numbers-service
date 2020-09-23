package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/Tiny-Paws/numbers-service/pkg/numbers"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	svc := numbers.NewNumbersService()
	svc = numbers.LoggingMiddleware{logger, svc}

	addHandler := httptransport.NewServer(
		numbers.MakeAddEndpoint(svc),
		numbers.DecodeAddRequest,
		numbers.EncodeResponse,
	)

	subHandler := httptransport.NewServer(
		numbers.MakeSubEndpoint(svc),
		numbers.DecodeSubRequest,
		numbers.EncodeResponse,
	)

	http.Handle("/add", addHandler)
	http.Handle("/sub", subHandler)
	http.ListenAndServe(":8080", nil)
}
