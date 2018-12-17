package workers

import (
	"errors"
	"time"
)

var (
	errTimeout = errors.New("withtimeout: Operation timed out")
)

type resultAndError struct {
	result interface{}
	err    error
}

// WithTimeOut 让方法可以超时，方法返回的第一个必须是结果，第二个参数必须是错误
func WithTimeOut(timeout time.Duration, fn func() (interface{}, error)) (result interface{}, timedOut bool, err error) {
	rstCh := make(chan *resultAndError, 1)

	go func() {
		rst, err := fn()
		rstCh <- &resultAndError{rst, err}
	}()

	select {
	case <-time.After(timeout * time.Second):
		return nil, true, errTimeout
	case rAe := <-rstCh:
		return rAe.result, false, rAe.err
	}
}
