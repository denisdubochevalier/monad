package monad

// Maybe is a Monad that allows a value to be either Just or Nothing
type Maybe[T any] interface {
	IsJust() bool
	IsNothing() bool
	Value() T
	OrElse(T) T
	Filter(p Predicate[T]) Maybe[T]
	FMap(func(T) Maybe[T]) Maybe[T]
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

// FMap applies a callback to the value
func (j Just[T]) FMap(t func(T) Maybe[T]) Maybe[T] {
	return t(j.val)
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

// FMap does nothing
func (n Nothing[T]) FMap(_ func(T) Maybe[T]) Maybe[T] {
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
