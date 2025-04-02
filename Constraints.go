package compare

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Integer is a constraint that permits any integer type.
type Integer interface {
	Unsigned | Signed
}

// Float is a constraint that permits any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Real is a constraint that permits any integer or floating-point type.
type Real interface {
	Integer | Float
}

// Complex is a constraint that permits any complex numeric type.
type Complex interface {
	~complex64 | ~complex128
}

// Number is a constraint that permits any real or complex numeric type.
type Number interface {
	Real | Complex
}

// String is a constraint that permits any string type.
type String interface {
	~string
}

// Equatable is an interface that is implemented by all types that can be compared using equal (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types).
// The Equatable interface may only be used as a type parameter constraint, not as the type of a variable.
type Equatable = comparable

// Comparable is a constraint that permits any real numeric type or string type.
type Comparable interface {
	Real | String
}
