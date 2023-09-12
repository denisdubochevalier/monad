package monad

// Maybe is a Monad that allows a value to be either Just or Nothing
type Maybe[T any] interface {
	IsJust() bool
	IsNothing() bool
	Value() T
	OrElse(T) T
	Filter(p Predicate[T]) Maybe[T]
	Map(func(T) Maybe[any]) Maybe[any]
	FlatMap(func(T) Maybe[T]) Maybe[T]
}

// Just represents a Maybe monad with a just value
type Just[T any] struct {
	val T
}

// IsJust is true
func (j Just[T]) IsJust() bool {
	return true
}

// IsNothing is false
func (j Just[T]) IsNothing() bool {
	return false
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
	return None[T]()
}

// Map applies a callback to the value and returns a Maybe[any]
func (j Just[T]) Map(f func(T) Maybe[any]) Maybe[any] {
	return f(j.val)
}

// FlatMap applies a callback to the value
func (j Just[T]) FlatMap(f func(T) Maybe[T]) Maybe[T] {
	return f(j.val)
}

// Nothing represents an empty Maybe of type T
type Nothing[T any] struct{}

// IsJust is false
func (n Nothing[T]) IsJust() bool {
	return false
}

// IsNothing is true
func (n Nothing[T]) IsNothing() bool {
	return true
}

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

// Map does nothing
func (n Nothing[T]) Map(_ func(T) Maybe[any]) Maybe[any] {
	return Nothing[any](n)
}

// FlatMap does nothing
func (n Nothing[T]) FlatMap(_ func(T) Maybe[T]) Maybe[T] {
	return n
}

// Some creates a Just Maybe from a value
func Some[T any](x T) Maybe[T] {
	return Just[T]{val: x}
}

// None creates a Nothing Maybe
func None[T any]() Maybe[T] {
	return Nothing[T]{}
}

// Nullable creates a Maybe from its input, Nothing if it is nil, Just elseway
func Nullable[T any](x *T) Maybe[T] {
	if x == nil {
		return None[T]()
	}
	return Some[T](*x)
}
