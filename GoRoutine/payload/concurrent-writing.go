package payload

import (
	"fmt"
	"sync"
	"time"
)

func rotateStringRight(rotatedString chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	input := <-rotatedString
	if len(input) == 0 {
		return
	}
	input = string(input[len(input)-1]) + input[:len(input)-1]
	fmt.Println("Rotating right", input)
	time.Sleep(1 * time.Millisecond)
	go writeDataToChan(input, rotatedString)
}

func rotateStringLeft(rotatedString chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	input := <-rotatedString
	if len(input) == 0 {
		return
	}
	input = input[1:] + string(input[0])
	fmt.Println("Rotating left", input)
	time.Sleep(1 * time.Millisecond)
	go writeDataToChan(input, rotatedString)
}

func writeDataToChan(data string, channel chan<- string) {
	channel <- data
}

func ConcurrentStringRotation() {
	payload := make(chan string)
	defer close(payload)

	go writeDataToChan(Data[0], payload)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go rotateStringRight(payload, &wg)
		go rotateStringLeft(payload, &wg)
	}

	wg.Wait()
	formattedString, _ := <-payload
	fmt.Println("Result: ", formattedString)
}
