package timeloop

import (
	"time"
)


type job struct {
	execute func()
	interval time.Duration
	complete chan bool
	exit chan bool
}

func newJob(execute func(), interval time.Duration) *job {
	return &job{
		execute: execute,
		interval: interval,
		complete: make(chan bool),
		exit: make(chan bool),
	}
}

func(self *job) startLoop() {
	ticker := time.NewTicker(self.interval)
	for {
		select {
			case <- ticker.C:
				self.execute()
			case <- self.complete:
				ticker.Stop()
				return
		}
	}
}

func(self *job) start() {
	self.startLoop()
	self.exit <- true
}

func(self *job) stop() {
	defer func() {
		<- self.exit
	}()
	self.complete <- true
}