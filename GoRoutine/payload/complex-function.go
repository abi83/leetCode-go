package payload

import (
	"fmt"
	"sync"
	"time"
)

const BASE_COMPLEXITY = 1_000_000
const COMPLEXITY_MULTIPLIER = 100

var Data = []string{
	"01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
	"0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
	"012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
	"01234567890123456789012345678901234567890123456789012345678901234567890123456789",
	"0123456789012345678901234567890123456789012345678901234567890123456789",
	"012345678901234567890123456789012345678901234567890123456789",
	"01234567890123456789012345678901234567890123456789",
	"0123456789012345678901234567890123456789",
	"012345678901234567890123456789",
	"01234567890123456789",
	"0123456789",
}

func HeavyComputationsSync(input string, multiplier int) string {
	for i := 0; i < BASE_COMPLEXITY*multiplier; i++ {
		input = string(input[len(input)-1]) + input[:len(input)-1]
	}
	return input
}

func HeavyComputationsRoutine(input string, multiplier int, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < BASE_COMPLEXITY*multiplier; i++ {
		input = string(input[len(input)-1]) + input[:len(input)-1]
	}
	result <- input
}

func RunHeavyComputationsSync() {
	fmt.Println("Starting Sync")
	startTime := time.Now()
	for i, element := range Data {
		result := HeavyComputationsSync(element, COMPLEXITY_MULTIPLIER)
		fmt.Printf("Received sync result (%d):(%.2f) sec. %s\n", i, time.Now().Sub(startTime).Seconds(), result)
	}
	fmt.Printf("Elapsed time Sync computations: %.2f\n", time.Now().Sub(startTime).Seconds())
}

func RunHeavyComputationsInRoutites() {
	fmt.Println("Starting go routines")
	startTime := time.Now()
	resultChan := make(chan string)
	defer close(resultChan)

	receiveResults := func(resultChan <-chan string) {
		i := 1
		for result := range resultChan {
			fmt.Printf("Received go routines result (%d):(%.2f) sec. %s\n", i, time.Now().Sub(startTime).Seconds(), result)
			i++
		}
	}

	var wg sync.WaitGroup
	for _, element := range Data {
		wg.Add(1)
		go HeavyComputationsRoutine(element, COMPLEXITY_MULTIPLIER, resultChan, &wg)
	}

	go receiveResults(resultChan)

	wg.Wait()
	fmt.Printf("Elapsed time routines: %.2f\n", time.Now().Sub(startTime).Seconds())
}
