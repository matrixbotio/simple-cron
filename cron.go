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
}

// NewCronHandler - create new cron
func NewCronHandler(callback CronCallback, timerTime time.Duration) *CronObject {
	return &CronObject{
		timerTime: timerTime,
		callback:  callback,
		stopCh:    make(chan struct{}),
	}
}

// Stop cron
func (c *CronObject) Stop() {
	c.stopCh <- struct{}{}
}

// CronCallback - cron callback /ᐠ｡‸｡ᐟ\
type CronCallback func()

// Run cron
func (c *CronObject) Run(immediately ...bool) {
	active := true
	sleepAtStart := true
	if len(immediately) > 0 && immediately[0] {
		sleepAtStart = false
	}
	if sleepAtStart {
		time.Sleep(c.timerTime)
	}

	go func() {
		<-c.stopCh
		active = false
	}()

	for active {
		c.callback()
		time.Sleep(c.timerTime)
	}
}
