package monad

// Maybe is a Monad that allows a value to be either just or Nothing
type Maybe[T any] interface {
	Just() bool
	Nothing() bool
	Value() T
	OrElse(T) T
	Filter(p Predicate[T]) Maybe[T]
	Map(func(T) Maybe[any]) Maybe[any]
	FlatMap(func(T) Maybe[T]) Maybe[T]
}

// just represents a Maybe monad with a just value
type just[T any] struct {
	val T
}

// Just is true
func (j just[T]) Just() bool {
	return true
}

// Nothing is false
func (j just[T]) Nothing() bool {
	return false
}

// Value gives the underlying just value
func (j just[T]) Value() T {
	return j.val
}

// OrElse gives the underlying
func (j just[T]) OrElse(_ T) T {
	return j.val
}

// Filter returns the just value if the predicate is true, nothing elseway
func (j just[T]) Filter(p Predicate[T]) Maybe[T] {
	if p(j.val) {
		return j
	}
	return None[T]()
}

// Map applies a callback to the value and returns a Maybe[any]
func (j just[T]) Map(f func(T) Maybe[any]) Maybe[any] {
	return f(j.val)
}

// FlatMap applies a callback to the value
func (j just[T]) FlatMap(f func(T) Maybe[T]) Maybe[T] {
	return f(j.val)
}

// nothing represents an empty Maybe of type T
type nothing[T any] struct{}

// Just is false
func (n nothing[T]) Just() bool {
	return false
}

// Nothing is true
func (n nothing[T]) Nothing() bool {
	return true
}

// Value returns the zero value of the underlying type
func (n nothing[T]) Value() T {
	return *new(T)
}

// OrElse returns the else value
func (n nothing[T]) OrElse(x T) T {
	return x
}

// Filter returns nothing
func (n nothing[T]) Filter(_ Predicate[T]) Maybe[T] {
	return n
}

// Map does nothing
func (n nothing[T]) Map(_ func(T) Maybe[any]) Maybe[any] {
	return nothing[any](n)
}

// FlatMap does nothing
func (n nothing[T]) FlatMap(_ func(T) Maybe[T]) Maybe[T] {
	return n
}

// Some creates a just Maybe from a value
func Some[T any](x T) Maybe[T] {
	return just[T]{val: x}
}

// None creates a nothing Maybe
func None[T any]() Maybe[T] {
	return nothing[T]{}
}

// Nullable creates a Maybe from its input, nothing if it is nil, just elseway
func Nullable[T any](x *T) Maybe[T] {
	if x == nil {
		return None[T]()
	}
	return Some[T](*x)
}
