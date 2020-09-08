// Test module to try out module versioning
package utils

// Contains returns true if x is in a, false otherwise
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// ContainsInt checks if int `x` contained in `a` slice
func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}