package LeetCode75

import "testing"

func TestIsSubsequenceWithSimpleInput(t *testing.T) {
	if !IsSubsequence("abc", "axbxc") {
		t.Errorf("Simple subsequence not found")
	}
}

func TestIsSubsequenceFalseReturn(t *testing.T) {
	if IsSubsequence("axc", "ahbgdc") {
		t.Errorf("Wrong subsequence found")
	}
}

func TestIsSubsequenceEmptyInput(t *testing.T) {
	if !IsSubsequence("", "ahbgdc") {
		t.Errorf("Not works properly with empty string")
	}
	if IsSubsequence("abc", "") {
		t.Errorf("Not works properly with empty string")
	}

}
