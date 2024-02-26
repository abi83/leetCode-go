package payload

import (
	"fmt"
	"time"
)

func ticker(timeout int, ch chan<- int) {
	time.Sleep(time.Second * time.Duration(timeout))
	ch <- timeout
}

func RunTimer() {
	ch := make(chan int)
	defer close(ch)

	for i := 0; i < 10; i++ {
		go ticker(i, ch)
	}

	for data := range ch {
		fmt.Println("Tick: ", data)
	}
}
