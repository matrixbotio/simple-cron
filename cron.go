package simplecron

import (
	"sync"
	"time"
)

type CronObject struct {
	timerTime time.Duration
	callback  func()
	stopCh    chan struct{}
	stopWG    sync.WaitGroup
	active    bool
	paused    bool
}

func NewCronHandler(callback func(), timerTime time.Duration) *CronObject {
	return &CronObject{
		timerTime: timerTime,
		callback:  callback,
		stopCh:    make(chan struct{}),
		active:    false,
		paused:    false,
	}
}

func (c *CronObject) Stop() {
	if c.active {
		c.active = false
		close(c.stopCh)
		c.stopWG.Wait()
	}
}

func (c *CronObject) Run(immediately ...bool) {
	c.active = true
	sleepAtStart := true

	if len(immediately) > 0 && immediately[0] {
		sleepAtStart = false
	}

	c.stopWG.Add(1)

	go func() {
		defer c.stopWG.Done()

		if sleepAtStart {
			select {
			case <-time.After(c.timerTime):
			case <-c.stopCh:
				return
			}
		}

		for c.active {
			if !c.paused {
				c.callback()
			}

			select {
			case <-time.After(c.timerTime):
			case <-c.stopCh:
				return
			}
		}
	}()
}

func (c *CronObject) IsActive() bool {
	return c.active && !c.paused
}

func (c *CronObject) IsPaused() bool {
	return c.paused
}

func (c *CronObject) Pause() {
	c.paused = true
}

func (c *CronObject) Resume() {
	c.paused = false
}
