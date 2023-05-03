package Alghorithm

import (
	"math/rand"
	"testing"
	"time"
)

func TestSimpleOddScenario(t *testing.T) {
	if BinarySearch([]int{1, 4, 7}, 4) != 1 {
		t.Errorf("Fails in simple case")
	}
}

func TestSimpleEvenScenario(t *testing.T) {
	if BinarySearch([]int{1, 4, 6, 7}, 4) != 1 {
		t.Errorf("Fails in simple case")
	}
}

func TestMinimalisticScenario(t *testing.T) {
	if BinarySearch([]int{2, 5}, 2) != 0 {
		t.Errorf("Fails in minimalistic case")
	}
}

func TestPositiveScenario(t *testing.T) {
	if BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		t.Errorf("Search fails")
	}
}

func TestNegativeScenario(t *testing.T) {
	if BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) != -1 {
		t.Errorf("Search fails")
	}
}

func TestNotFoundScenario(t *testing.T) {
	res := BinarySearch([]int{1, 4, 7}, 99)
	if res != -1 {
		t.Errorf("Not found scenario returns %v, expected '-1'", res)
	}
}

func TestFirstPositionScenario(t *testing.T) {
	res := BinarySearch([]int{1, 4, 5, 7, 10}, 1)
	if res != 0 {
		t.Errorf("First position scenario returns %v, expected '0'", res)
	}
}

func TestLastPositionScenario(t *testing.T) {
	res := BinarySearch([]int{1, 4, 5, 7, 10}, 10)
	if res != 4 {
		t.Errorf("Last position scenario returns %v, expected '4'", res)
	}
}

func TestPerformance(t *testing.T) {
	const SIZE = 100_000_000
	const STEP = 100
	var bigSlice = make([]int, SIZE)
	bigSlice[0] = rand.Intn(STEP)
	for i := 1; i < SIZE; i++ {
		bigSlice[i] = bigSlice[i-1] + rand.Intn(STEP)
	}
	var target = bigSlice[rand.Intn(SIZE)]
	start := time.Now()
	BinarySearch(bigSlice, target)
	duration := time.Since(start)
	if duration > time.Millisecond {
		t.Errorf("Too long execution time: %v", duration)
	}
}
