package main

import (
	"errors"
	"strings"
)

// business logic

type StringService interface {
	UpperCase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (g stringService) UpperCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (g stringService) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("empty string")
