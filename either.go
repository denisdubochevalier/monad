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

// left represent a left value.
type left[T any] struct {
	val T
}

// NewLVal creates a left value.
func NewLVal[T any](t T) Either[T] {
	return left[T]{
		val: t,
	}
}

// Value gets the underlying value.
func (l left[T]) Value() T {
	return l.val
}

// Left is true.
func (l left[T]) Left() bool {
	return true
}

// Right is false.
func (l left[T]) Right() bool {
	return false
}

// Map applies the mapping function to the left value and returns a new Either
func (l left[T]) Map(f func(T) Either[any]) Either[any] {
	return f(l.val)
}

// FlatMap returns itself.
func (l left[T]) FlatMap(_ func(T) Either[T]) Either[T] {
	return l
}

// Or executes the callback.
func (l left[T]) Or(f func(T) Either[T]) Either[T] {
	return f(l.val)
}

// right represents a right value.
type right[T any] struct {
	val T
}

// NewRVal creates a right value.
func NewRVal[T any](t T) Either[T] {
	return right[T]{
		val: t,
	}
}

// Value gets the underlying value.
func (r right[T]) Value() T {
	return r.val
}

// Left is false.
func (r right[T]) Left() bool {
	return false
}

// Right is false.
func (r right[T]) Right() bool {
	return true
}

// Map applies the mapping function to the right value and returns a new Either
func (r right[T]) Map(f func(T) Either[any]) Either[any] {
	return f(r.val)
}

// FlatMap applies its callback.
func (r right[T]) FlatMap(f func(T) Either[T]) Either[T] {
	return f(r.val)
}

// Or returns itself.
func (r right[T]) Or(f func(T) Either[T]) Either[T] {
	return r
}
