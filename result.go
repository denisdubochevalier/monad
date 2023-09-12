package monad

// Result represents a result of an operation that can fail
type Result[T any] interface {
	Error() error
	Value() T
	Failure() bool
	Success() bool
	Map(func(T) Result[any]) Result[any]
	FlatMap(func(T) Result[T]) Result[T]
	Or(ErrorHandler[T]) Result[T]
}

// Success is the successful return of a failable operation
type Success[T any] struct {
	val T
}

// Error always returns nil for a success
func (s Success[_]) Error() error {
	return nil
}

// Value returns the underlying value
func (s Success[T]) Value() T {
	return s.val
}

// Failure always return false for a Success
func (s Success[_]) Failure() bool {
	return false
}

// Success always return true for a Success
func (s Success[_]) Success() bool {
	return true
}

// Map executes the callback functions and returns a result with the value changed to any type.
func (s Success[T]) Map(f func(T) Result[any]) Result[any] {
	return f(s.val)
}

// FlatMap executes the callback function and returns a success with the value changed to the result of the callback
func (s Success[T]) FlatMap(f func(T) Result[T]) Result[T] {
	return f(s.val)
}

// Or returns the success
func (s Success[T]) Or(_ ErrorHandler[T]) Result[T] {
	return s
}

// Succeed creates a Success
func Succeed[T any](val T) Result[T] {
	return Success[T]{val: val}
}

// Failure represents an operation failure
type Failure[T any] struct {
	err error
}

// Error returns the underlying error
func (f Failure[_]) Error() error {
	return f.err
}

// Value returns the zero value of the type
func (f Failure[T]) Value() T {
	return *new(T)
}

// Failure always returns true for a Failure
func (f Failure[_]) Failure() bool {
	return true
}

// Success always returns false for a Failure
func (f Failure[_]) Success() bool {
	return false
}

// Map returns the original failure in a Result[any] monad
func (f Failure[T]) Map(_ func(T) Result[any]) Result[any] {
	return Failure[any](f)
}

// FlatMap returns the original failure
func (f Failure[T]) FlatMap(_ func(T) Result[T]) Result[T] {
	return f
}

// Or executes the callback on the error and returns a new Failure with the result of the error - or the same Failure if error returned is nil
func (f Failure[T]) Or(e ErrorHandler[T]) Result[T] {
	e(f.err)
	return f
}

// Fail creates a Failure
func Fail[T any](err error) Result[T] {
	return Failure[T]{err: err}
}

// FromTuple creates a Result object from a tupple value, error
func FromTuple[T any](val T, err error) Result[T] {
	if err != nil {
		return Fail[T](err)
	}
	return Succeed(val)
}
