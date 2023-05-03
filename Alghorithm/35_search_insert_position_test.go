package Alghorithm

import "testing"

func TestFindInsertPositionOnEOddScenario(t *testing.T) {
	if FindInsertPosition([]int{1, 2, 3, 4, 5}, 3) != 2 {
		t.Errorf("Fails in odd scenario")
	}
}

func TestFindInsertPositionOnEvenScenario(t *testing.T) {
	if FindInsertPosition([]int{1, 2, 3, 4}, 3) != 2 {
		t.Errorf("Fails in even scenario")
	}
}

func TestFindInsertPositionOnOneElementScenario(t *testing.T) {
	if FindInsertPosition([]int{100}, 100) != 0 {
		t.Errorf("Fails in one element scenario")
	}
}

func TestFindInsertPositionOnMinimalisticScenario(t *testing.T) {
	if FindInsertPosition([]int{100, 200}, 200) != 1 {
		t.Errorf("Fails in a minimalistic scenario")
	}
}

func TestFindInsertPositionOnFirstElementScenario(t *testing.T) {
	if FindInsertPosition([]int{1, 2, 3, 4, 5, 6, 7}, 1) != 0 {
		t.Errorf("Fails in first element scenario")
	}
}

func TestFindInsertPositionOnLastElementScenario(t *testing.T) {
	if FindInsertPosition([]int{1, 2, 3, 4, 5, 6, 7}, 6) != 5 {
		t.Errorf("Fails in last element scenario")
	}
}

func TestFindInsertPositionOnNotFoundScenario(t *testing.T) {
	if FindInsertPosition([]int{1, 2, 4, 5, 6, 7}, 3) != 2 {
		t.Errorf("Fails in not found")
	}
}

func TestFindInsertPositionOnLeetCodeCase1(t *testing.T) {
	if FindInsertPosition([]int{1, 3, 5, 6}, 5) != 2 {
		t.Errorf("Fails in leetcode case1")
	}
}

func TestFindInsertPositionOnLeetCodeCase2(t *testing.T) {
	if FindInsertPosition([]int{1, 3, 5, 6}, 2) != 1 {
		t.Errorf("Fails in leetcode case2")
	}
}

func TestFindInsertPositionOnLeetCodeCase3(t *testing.T) {
	if FindInsertPosition([]int{1, 3, 5, 6}, 7) != 4 {
		t.Errorf("Fails in leetcode case3")
	}
}

func TestFindInsertPositionOnLeetCodeCase4(t *testing.T) {
	if FindInsertPosition([]int{1, 3, 5, 6}, 0) != 0 {
		t.Errorf("Fails in leetcode case4")
	}
}
