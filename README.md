# timeloop
timeloop is a service that can be used to run periodic tasks after a certain interval.

![timeloop](http://66.42.57.109/timeloop.jpg)


## Installation
```sh
go get github.com/sankalpjonn/go-timeloop
```

## Writing jobs
```go
package main

import (
	"log"
	"github.com/sankalpjonn/go-timeloop"
)
func main() {
	tl := timeloop.New()

	tl.AddJob(func() {
		log.Println("printing Test2SecTimer")
	}, time.Second * 2)

	tl.AddJob(func() {
		log.Println("printing Test5SecTimer")
	}, time.Second * 5)

	tl.AddJob(func() {
		log.Println("printing Test10SecTimer")
	}, time.Second * 10)	
}
```

## Start time loop in separate thread
By default timeloop starts in a separate thread.

Please do not forget to call ```tl.stop``` before exiting the program, Or else the jobs wont shut down gracefully.

```go
tl.Start()
defer tl.Stop()

done := make(chan os.Signal)
signal.Notify(done, os.Interrupt, syscall.SIGTERM)
signal.Notify(done, os.Interrupt, syscall.SIGINT)
<-done

```

## Start time loop in main thread
Doing this will automatically shut down the jobs gracefully when the program is killed, so no need to  call ```tl.stop```
```go
tl.BlockingStart()
```

## Author
* **Sankalp Jonna**

Email me with any queries: [sankalpjonna@gmail.com](sankalpjonna@gmail.com).
