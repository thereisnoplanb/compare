package compare

// Equal returns
//
//	true if x equals y,
//	false if x is not equal y.
func Equal[T Equatable](x, y T) bool {
	return x == y
}
