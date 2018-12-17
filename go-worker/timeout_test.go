package workers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func sleep(second time.Duration) {
	time.Sleep(second * time.Second)
}

func TestTimeout(t *testing.T) {

	funRst, timedout, err := WithTimeOut(
		time.Duration(4),
		func() (interface{}, error) {
			sleep(time.Duration(3))
			return "over", nil
		})
	// 方法本身没有报错
	assert.Equal(t, nil, err)
	// 没有超时
	assert.Equal(t, false, timedout)
	// 拿结果
	assert.Equal(t, "over", funRst)

	funRst, timedout, err = WithTimeOut(
		time.Duration(2),
		func() (interface{}, error) {
			sleep(time.Duration(3))
			return "over", nil
		})
	assert.Equal(t, errTimeout, err)
	assert.Equal(t, true, timedout)
}
