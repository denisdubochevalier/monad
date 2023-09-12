package monad

// Identity is a generic interface representing the Identity monad.
type Identity[T any] interface {
	// Value returns the encapsulated value of the Identity monad.
	Value() T

	// Map applies a function to the encapsulated value and returns a new Identity.
	Map(func(T) any) Identity[any]

	// FlatMap applies a function that returns a new Identity and returns it.
	FlatMap(func(T) Identity[T]) Identity[T]
}

// identity is a concrete implementation of the Identity interface.
type identity[T any] struct {
	value T // The encapsulated value
}

// NewIdentity constructs a new Identity monad with an initial value.
func NewIdentity[T any](value T) Identity[T] {
	return identity[T]{value: value}
}

// Value returns the encapsulated value.
func (i identity[T]) Value() T {
	return i.value
}

// Map applies a function to the encapsulated value and returns a new Identity.
func (i identity[T]) Map(f func(T) any) Identity[any] {
	return NewIdentity(f(i.value))
}

// FlatMap applies a function that returns a new Identity and returns it.
func (i identity[T]) FlatMap(f func(T) Identity[T]) Identity[T] {
	return f(i.value)
}
