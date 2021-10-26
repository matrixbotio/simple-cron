package simplecron

import "time"

// RunTimeLimitHandler - func runtime limit handler
type RunTimeLimitHandler struct {
	timeout time.Duration
	runFunc func()
}

// NewRuntimeLimitHandler - create new func runtime limit handler
func NewRuntimeLimitHandler(timeout time.Duration, runFunc func()) *RunTimeLimitHandler {
	return &RunTimeLimitHandler{
		timeout: timeout,
		runFunc: runFunc,
	}
}

// Run - run func & limit runtime.
// returns: bool: true if time is up
func (r *RunTimeLimitHandler) Run() bool {
	timeTo := time.After(r.timeout)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-timeTo:
				done <- true
				return
			default:
				// wait
			}
		}
	}()

	go func() {
		r.runFunc()
		done <- false
		return
	}()

	return <-done
}
