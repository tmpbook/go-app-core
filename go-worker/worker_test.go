package workers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaster(t *testing.T) {
	M := NewMaster(5, 10)

	// var funcList []func() (interface{}, error)
	// for i := 0; i <= 5; i++ {
	// 	funcList = append(funcList, sleep)
	// }

	// M.Add(funcList)
	M.Wait()
	// for i := 0; i <= 5; i++ {
	// 	<-M.jobsRst
	// }
	assert.Equal(t, t, len(M.jobsRst))
}
