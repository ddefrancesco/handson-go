package main

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type proxymw struct {
	next      StringService
	uppercase endpoint.Endpoint
}

func (mw proxymw) UpperCase(s string) (output string, err error) {
	response, err := mw.uppercase(nil, uppercaseRequest{S: s})
	if err != nil {
		return "", err
	}
	resp := response.(uppercaseResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}

	return resp.V, nil
}

func proxingMiddleware(proxyURL string) ServiceMiddleware {

}
