package monad

// Either monad represents two equivalent values, left or right.
// Right is by convention the "default value".
type Either[T any] interface {
	Value() T
	Left() bool
	Right() bool
	FMap(func(T) Either[T]) Either[T]
	Or(func(T) Either[T]) Either[T]
}

// Left represent a left value.
type Left[T any] struct {
	val T
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

// FMap returns itself.
func (l Left[T]) FMap(_ func(T) Either[T]) Either[T] {
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

// Value gets the underlying value.
func (l Right[T]) Value() T {
	return l.val
}

// Left is false.
func (l Right[T]) Left() bool {
	return false
}

// Right is false.
func (l Right[T]) Right() bool {
	return true
}

// FMap applies its callback.
func (l Right[T]) FMap(f func(T) Either[T]) Either[T] {
	return f(l.val)
}

// Or returns itself.
func (l Right[T]) Or(f func(T) Either[T]) Either[T] {
	return l
}
