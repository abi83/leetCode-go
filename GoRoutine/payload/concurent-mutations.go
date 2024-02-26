package payload

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
}

func addOneYear(person *Person, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	currentAge := person.age
	person.age = currentAge + 1
	mu.Unlock()
}

func AgeRunner() {
	person := &Person{
		name: "Vladimir", age: 0,
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1_000; i++ {
		wg.Add(1)
		go addOneYear(person, &wg, &mu)
	}
	wg.Wait()
	fmt.Println("Person.age: ", person.age)

}
