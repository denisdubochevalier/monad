// Package monad implements a simple Maybe monad in go
package monad

// Maybe is a Monad that allows a value to be either Just or Nothing
type Maybe[T any] interface {
	Value() T
	OrElse(T) T
	Filter(p Predicate[T]) Maybe[T]
	Bind(Nilable[T]) Maybe[T]
	FMap(Transformable[T]) Maybe[T]
}

// Just represents a Maybe monad with a just value
type Just[T any] struct {
	val T
}

// Value gives the underlying just value
func (j Just[T]) Value() T {
	return j.val
}

// OrElse gives the underlying
func (j Just[T]) OrElse(_ T) T {
	return j.val
}

// Filter returns the just value if the predicate is true, nothing elseway
func (j Just[T]) Filter(p Predicate[T]) Maybe[T] {
	if p(j.val) {
		return j
	}
	return Nothing[T]{}
}

// Bind executes a function that returns a nillable object on the value and returns a Maybe
func (j Just[T]) Bind(n Nilable[T]) Maybe[T] {
	return OfNullable(n(j.val))
}

// FMap applies a callback to the value
func (j Just[T]) FMap(t Transformable[T]) Maybe[T] {
	return OfValue(t(j.val))
}

// Nothing represents an empty Maybe of type T
type Nothing[T any] struct{}

// Value returns the zero value of the underlying type
func (n Nothing[T]) Value() T {
	return *new(T)
}

// OrElse returns the else value
func (n Nothing[T]) OrElse(x T) T {
	return x
}

// Filter returns nothing
func (n Nothing[T]) Filter(_ Predicate[T]) Maybe[T] {
	return n
}

// Bind does nothing
func (n Nothing[T]) Bind(_ Nilable[T]) Maybe[T] {
	return n
}

// FMap does nothing
func (n Nothing[T]) FMap(_ Transformable[T]) Maybe[T] {
	return n
}

// OfValue creates a Just Maybe from a value
func OfValue[T any](x T) Maybe[T] {
	return Just[T]{val: x}
}

// Empty creates a Nothing Maybe
func Empty[T any]() Maybe[T] {
	return Nothing[T]{}
}

// OfNullable creates a Maybe from its input, Nothing if it is nil, Just elseway
func OfNullable[T any](x *T) Maybe[T] {
	if x == nil {
		return Empty[T]()
	}
	return OfValue[T](*x)
}
