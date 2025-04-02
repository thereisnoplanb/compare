package compare

// Greater reports whether x is greater than y.
// For floating-point types, a NaN is considered greater than any non-NaN,
// and 0.0 is not greater than (is equal to) -0.0.
func Greater[T Comparable](x, y T) bool {
	return (!isNaN(x) && isNaN(y)) || x > y
}
