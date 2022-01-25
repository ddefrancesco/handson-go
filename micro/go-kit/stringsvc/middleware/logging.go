package middleware

import "github.com/go-kit/kit/log"

type loggingMiddleware struct {
	logger log.Logger
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {

}

func (mw loggingMiddleware) Count(s string) (n int) {

}
