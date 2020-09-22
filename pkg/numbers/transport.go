package numbers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeAddEndpoint(svc NumbersService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		r, err := svc.Add(req.A, req.B)
		if err != nil {
			return addResponse{0, err.Error()}, nil
		}
		return addResponse{r, ""}, nil
	}
}

func MakeSubEndpoint(svc NumbersService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(subRequest)
		r, err := svc.Sub(req.A, req.B)
		if err != nil {
			return subResponse{0, err.Error()}, nil
		}
		return subResponse{r, ""}, nil
	}
}

func DecodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeSubRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request subRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
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
