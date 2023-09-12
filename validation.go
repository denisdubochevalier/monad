package monad

// Validation is an interface that models the Validation monad.
// It either contains a value of type T or an aggregated list of errors of type E.
type Validation[E, T any] interface {
	// Valid returns true if Validation contains a value, false otherwise.
	Valid() bool

	// Value returns the encapsulated value of type T.
	Value() T

	// Errors returns the aggregated errors of type E.
	Errors() []E

	// Map applies a transformation to the encapsulated value, assuming it's valid,
	// and returns a new Validation instance.
	Map(func(T) any) Validation[E, any]

	// FlatMap applies a transformation to the encapsulated value, assuming it's valid,
	// and returns a new Validation instance. It's useful for chaining multiple Validation operations.
	FlatMap(func(T) Validation[E, T]) Validation[E, T]
}

// validation is a concrete implementation of the Validation interface.
type validation[E, T any] struct {
	value  T
	errors []E
}

// NewValid returns a new Validation containing a value.
func NewValid[E, T any](value T) Validation[E, T] {
	return validation[E, T]{value: value, errors: nil}
}

// NewInvalid returns a new Validation containing an array of errors.
func NewInvalid[E, T any](errors []E) Validation[E, T] {
	return validation[E, T]{errors: errors}
}

// Valid checks if the Validation instance contains a value, not errors.
func (v validation[E, T]) Valid() bool {
	return v.errors == nil
}

// Value returns the encapsulated value.
func (v validation[E, T]) Value() T {
	return v.value
}

// Errors returns the encapsulated errors.
func (v validation[E, T]) Errors() []E {
	return v.errors
}

// Map applies a function to the encapsulated value (if present)
// and returns a new Validation instance.
func (v validation[E, T]) Map(f func(T) any) Validation[E, any] {
	if v.Valid() {
		newValue := f(v.value)
		return NewValid[E, any](newValue)
	}
	return NewInvalid[E, any](v.errors)
}

// FlatMap applies a function to the encapsulated value (if present)
// and returns a new Validation instance. The function should return
// a Validation instance itself, allowing for chaining.
func (v validation[E, T]) FlatMap(f func(T) Validation[E, T]) Validation[E, T] {
	if v.Valid() {
		return f(v.value)
	}
	return v
}
