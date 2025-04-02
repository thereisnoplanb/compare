package compare

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T Comparable](x T) bool {
	return x != x
}
