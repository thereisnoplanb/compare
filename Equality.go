package compare

// Represents the method that reports two objects of the same type are equal.
//
// # Parameters
//
//	x T
//
// The first object to compare for equality.
//
//	y T
//
// The second object to compare for equality.
//
// # Returns
//
//	bool
//
// True if the specified objects are equal; otherwise, false.
type Equality[T any] func(x, y T) bool
