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
	timer := time.NewTimer(r.timeout)
	done := make(chan bool, 1)

	go func() {
		<-timer.C
		done <- true
	}()

	go func() {
		r.runFunc()
		done <- false
		if !timer.Stop() {
			<-timer.C
		}
		return
	}()

	return <-done
}
