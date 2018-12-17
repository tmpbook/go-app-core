package workers

import (
	"time"
)

type job struct {
	jobID   interface{}
	jobFunc func() (interface{}, error)
}

type jobRst struct {
	JobID    interface{}
	Timedout bool
	Result   interface{}
	Err      error
}

// Master with some workers(goroutine)
type Master struct {
	allDone chan bool
	jobs    chan job
	timeout int64
	jobsRst chan jobRst
}

// NewMaster a Master with n workers
func NewMaster(n int, timeout int64) Master {

	procDone := make(chan bool)
	M := Master{
		allDone: make(chan bool),
		timeout: timeout,
	}

	for i := 0; i < n; i++ {
		go func() {
			for j := range M.jobs {
				funcRst, timedOut, err := WithTimeOut(time.Duration(M.timeout), j.jobFunc)
				// 结果存到 M 对象
				M.jobsRst <- jobRst{
					Timedout: timedOut,
					Err:      err,
					Result:   funcRst,
				}

			}
			procDone <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			_ = <-procDone
		}
		M.allDone <- true

	}()
	return M

}

// Add job function to worker
func (M Master) Add(funcList []func() (interface{}, error)) {
	for index, f := range funcList {
		M.jobs <- job{
			jobFunc: f,
			jobID:   index,
		}
	}
}

//Wait Call Once when all funcs have been added to wait for completion.
//If you don't care to wait, 'go Wait' lol
func (M Master) Wait() {
	close(M.jobs)
	_ = <-M.allDone
}
