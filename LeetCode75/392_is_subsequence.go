package LeetCode75

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sIndex := 0
	for tIndex := 0; tIndex < len(t); tIndex++ {
		if t[tIndex] == s[sIndex] {
			sIndex++
		}
		if sIndex == len(s) {
			return true
		}
	}
	return false
}
