// Package monad implements a simple Maybe monad in go
package monad

// Maybe is a Monad that allows a value to be either Just or Nothing
type Maybe[T any] interface {
	Value() T
	OrElse(T) T
	Filter(p Predicate[T]) Maybe[T]
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

// Map applies a Func to a Maybe
func Map[V, T any](m Maybe[V], f Func[V, T]) Maybe[T] {
	switch mm := m.(type) {
	case Just[V]:
		return Just[T]{
			val: f(mm.val),
		}
	}
	return Nothing[T]{}
}

// FlatMap applies a Func returning a Maybe to a Maybe
func FlatMap[V, T any](m Maybe[V], f Func[V, Maybe[T]]) Maybe[T] {
	switch mm := m.(type) {
	case Just[V]:
		return f(mm.val)
	}
	return Nothing[T]{}
}
