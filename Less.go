package compare

// Less reports whether x is less than y.
// For floating-point types, a NaN is considered less than any non-NaN,
// and -0.0 is not less than (is equal to) 0.0.
func Less[T Comparable](x, y T) bool {
	return (isNaN(x) && !isNaN(y)) || x < y
}
