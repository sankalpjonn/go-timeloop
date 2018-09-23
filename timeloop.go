package timeloop

import (
	"time"
	"os"
	"os/signal"
	"syscall"
)

type Timeloop struct {
 jobs []*job
}



func New() *Timeloop {
	return new(Timeloop)
}

func (self *Timeloop) Job(f func(), interval time.Duration) {
	self.jobs = append(self.jobs, newJob(f, interval))
}


func (self *Timeloop) start(block bool) {
	for _, j := range self.jobs {
		go j.start()
	}

	if block {
		defer self.Stop()

		done := make(chan os.Signal)
	    signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	    signal.Notify(done, os.Interrupt, syscall.SIGINT)

	    <-done
	}
}

func (self *Timeloop) BlockingStart() {
	self.start(true)
}

func (self *Timeloop) Start() {
	self.start(false)
}

func (self *Timeloop) Stop() {
	for _, j := range self.jobs {
		j.stop()
	}
}
