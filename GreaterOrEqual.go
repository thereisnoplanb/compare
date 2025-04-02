package compare

// GreaterOrEqual reports whether x is greater than or equal y.
// For floating-point types, a NaN is considered greater than or equal any non-NaN,
// and 0.0 is not greater than (is equal to) -0.0.
func GreaterOrEqual[T Comparable](x, y T) bool {
	return (!isNaN(x) && isNaN(y)) || x >= y
}
