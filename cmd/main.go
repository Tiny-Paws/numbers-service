package main

import (
	"net/http"

	"github.com/Tiny-Paws/numbers-service/pkg/numbers"
	httptransport "github.com/go-kit/kit/transport/http"
)

func init() {

}

func main() {
	svc := numbers.NewNumbersService()

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
