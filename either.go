package monad

// Either monad represents two equivalent values, left or right.
// Right is by convention the "default value".
type Either[T any] interface {
	Value() T
	Left() bool
	Right() bool
	Map(func(T) Either[any]) Either[any]
	FlatMap(func(T) Either[T]) Either[T]
	Or(func(T) Either[T]) Either[T]
}

// Left represent a left value.
type Left[T any] struct {
	val T
}

// NewLVal creates a Left value.
func NewLVal[T any](t T) Either[T] {
	return Left[T]{
		val: t,
	}
}

// Value gets the underlying value.
func (l Left[T]) Value() T {
	return l.val
}

// Left is true.
func (l Left[T]) Left() bool {
	return true
}

// Right is false.
func (l Left[T]) Right() bool {
	return false
}

// Map applies the mapping function to the Left value and returns a new Either
func (l Left[T]) Map(f func(T) Either[any]) Either[any] {
	return f(l.val)
}

// FlatMap returns itself.
func (l Left[T]) FlatMap(_ func(T) Either[T]) Either[T] {
	return l
}

// Or executes the callback.
func (l Left[T]) Or(f func(T) Either[T]) Either[T] {
	return f(l.val)
}

// Right represents a right value.
type Right[T any] struct {
	val T
}

// NewRVal creates a Right value.
func NewRVal[T any](t T) Either[T] {
	return Right[T]{
		val: t,
	}
}

// Value gets the underlying value.
func (r Right[T]) Value() T {
	return r.val
}

// Left is false.
func (r Right[T]) Left() bool {
	return false
}

// Right is false.
func (r Right[T]) Right() bool {
	return true
}

// Map applies the mapping function to the Right value and returns a new Either
func (r Right[T]) Map(f func(T) Either[any]) Either[any] {
	return f(r.val)
}

// FlatMap applies its callback.
func (r Right[T]) FlatMap(f func(T) Either[T]) Either[T] {
	return f(r.val)
}

// Or returns itself.
func (r Right[T]) Or(f func(T) Either[T]) Either[T] {
	return r
}
