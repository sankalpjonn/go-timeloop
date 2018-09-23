package timeloop

import (
	"testing"
	"time"
	"log"
	"os"
	"syscall"
	"os/signal"
)


func TestTimeloopBlockingStart(*testing.T) {
	log.Println("Running blocking start")

	tl := New()

	tl.Job(func() {
		log.Println("printing Test2SecTimer")
	}, time.Second * 2)

	tl.Job(func() {
		log.Println("printing Test5SecTimer")
	}, time.Second * 5)

	tl.Job(func() {
		log.Println("printing Test10SecTimer")
	}, time.Second * 10)

	tl.BlockingStart()

	log.Println("Exiting blocking start")
}


func TestTimeloopBackgroundStart(*testing.T) {
	log.Println("Running non blocking start")

	tl := New()

	tl.Job(func() {
		log.Println("printing Test2SecTimer")
	}, time.Second * 2)

	tl.Job(func() {
		log.Println("printing Test5SecTimer")
	}, time.Second * 5)

	tl.Job(func() {
		log.Println("printing Test10SecTimer")
	}, time.Second * 10)

	tl.Start()
	defer tl.Stop()

	done := make(chan os.Signal)
    signal.Notify(done, os.Interrupt, syscall.SIGTERM)
    signal.Notify(done, os.Interrupt, syscall.SIGINT)
    <-done

    log.Println("Exiting non blocking start")
}

