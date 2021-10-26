package simplecron

import "time"

type runTimeLimitHandler struct {
	timeout time.Duration
	runFunc func()
}

func newRuntimeLimitHandler(timeout time.Duration, runFunc func()) *runTimeLimitHandler {
	return &runTimeLimitHandler{
		timeout: timeout,
		runFunc: runFunc,
	}
}

// returns: bool: true if time is up
func (r *runTimeLimitHandler) run() bool {
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
