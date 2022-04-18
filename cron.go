package simplecron

import "time"

/*
    ___ _ __ ___  _ __
   / __| '__/ _ \| '_ \
  | (__| | | (_) | | | |
   \___|_|  \___/|_| |_|
*/

// CronObject - cron object /ᐠ｡‸｡ᐟ\
type CronObject struct {
	timerTime time.Duration
	callback  CronCallback
	stopCh    chan struct{}
	active    bool
	paused    bool
}

// NewCronHandler - create new cron
func NewCronHandler(callback CronCallback, timerTime time.Duration) *CronObject {
	return &CronObject{
		timerTime: timerTime,
		callback:  callback,
		stopCh:    make(chan struct{}, 1),
	}
}

// Stop cron
func (c *CronObject) Stop() {
	c.active = false
	c.stopCh <- struct{}{}
}

// Pause cron event exec
func (c *CronObject) Pause() {
	c.paused = true
}

// IsPaused - check cron is paused status
func (c *CronObject) IsPaused() bool {
	return c.paused
}

// Resume cron event exec
func (c *CronObject) Resume() {
	c.paused = false
}

func (c *CronObject) awaitStop() {
	<-c.stopCh
	c.active = false
}

// CronCallback - cron callback /ᐠ｡‸｡ᐟ\
type CronCallback func()

// Run cron
func (c *CronObject) Run(immediately ...bool) {
	c.active = true
	sleepAtStart := true

	go c.awaitStop()

	if len(immediately) > 0 && immediately[0] {
		sleepAtStart = false
	}
	if sleepAtStart {
		time.Sleep(c.timerTime)
	}

	for c.active {
		if !c.active {
			break
		}
		if !c.paused {
			c.callback()
		}
		time.Sleep(c.timerTime)
	}
}
