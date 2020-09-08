// Test module to try out module versioning
package utils

// InSlice returns true if x is in a, false otherwise
func InSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// InSliceInt checks if int `x` contained in `a` slice
func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}