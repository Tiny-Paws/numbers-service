package numbers

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeAddEndpoint(svc NumbersService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		r, err := svc.Add(req.A, req.B)
		if err != nil {
			return addResponse{0, err.Error()}, nil
		}
		return addResponse{r, ""}, nil
	}
}

func makeSubEndpoint(svc NumbersService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(subRequest)
		r, err := svc.Sub(req.A, req.B)
		if err != nil {
			return subResponse{0, err.Error()}, nil
		}
		return subResponse{r, ""}, nil
	}
}

type addRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type addResponse struct {
	R   int    `json:"r"`
	Err string `json:"err,omitempty"`
}

type subRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type subResponse struct {
	R   int    `json:"r"`
	Err string `json:"err,omitempty"`
}
