// This program implement HasPrefix() HasSuffix() Contains()
// defined in strings package

package main

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Use HasPrefix to check wheter substr is within s.
func Contains_HasPrefix(s, substr string) bool {
	if substr == "" {
		return true
	}
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// Use HasSuffix to check wheter substr is within s.
func Contains_HasSuffix(s, substr string) bool {
	// Explicitly check if substr is "", but in HasSuffix version, it's no need.
	if substr == "" {
		return true
	}

	for i := len(s); i > 0; i-- {
		if HasSuffix(s[:i], substr) {
			return true
		}
	}
	return false
}
