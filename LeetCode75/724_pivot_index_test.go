package LeetCode75

import "testing"

type pivotTest struct {
	input    []int
	expected int
}

var pivotTests = []pivotTest{
	{input: []int{1, 7, 3, 6, 5, 6}, expected: 3},
	{input: []int{1, 2, 3}, expected: -1},
	{input: []int{2, 1, -1}, expected: 0},
}

func TestPivotIndex(t *testing.T) {
	for _, test := range pivotTests {
		if output := PivotIndex(test.input); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
