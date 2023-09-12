package monad

import "reflect"

// ErrorHandler represents a function that handles an error
type ErrorHandler[T, E any] func(E) Result[T, E]

// Result represents a result of an operation that can fail
type Result[T, E any] interface {
	Error() E
	Value() T
	Failure() bool
	Success() bool
	Map(func(T) any) Result[any, E]
	FlatMap(func(T) Result[T, E]) Result[T, E]
	Or(ErrorHandler[T, E]) Result[T, E]
}

// success is the successful return of a failable operation
type success[T, E any] struct {
	val T
}

// Error always returns nil for a success
func (s success[T, E]) Error() E {
	return *(new(E))
}

// Value returns the underlying value
func (s success[T, E]) Value() T {
	return s.val
}

// Failure always return false for a success
func (s success[_, _]) Failure() bool {
	return false
}

// Success always return true for a success
func (s success[_, _]) Success() bool {
	return true
}

// Map executes the callback functions and returns a result with the value changed to any type.
func (s success[T, E]) Map(f func(T) any) Result[any, E] {
	return Succeed[any, E](f(s.val))
}

// FlatMap executes the callback function and returns a success with the value changed to the result of the callback
func (s success[T, E]) FlatMap(f func(T) Result[T, E]) Result[T, E] {
	return f(s.val)
}

// Or returns the success
func (s success[T, E]) Or(_ ErrorHandler[T, E]) Result[T, E] {
	return s
}

// Succeed creates a success
func Succeed[T, E any](val T) Result[T, E] {
	return success[T, E]{val: val}
}

// failure represents an operation failure
type failure[T, E any] struct {
	err E
}

// Error returns the underlying error
func (f failure[_, E]) Error() E {
	return f.err
}

// Value returns the zero value of the type
func (f failure[T, E]) Value() T {
	return *new(T)
}

// Failure always returns true for a failure
func (f failure[_, _]) Failure() bool {
	return true
}

// Success always returns false for a failure
func (f failure[_, _]) Success() bool {
	return false
}

// Map returns the original failure in a Result[any] monad
func (f failure[T, E]) Map(_ func(T) any) Result[any, E] {
	return failure[any, E](f)
}

// FlatMap returns the original failure
func (f failure[T, E]) FlatMap(_ func(T) Result[T, E]) Result[T, E] {
	return f
}

// Or executes the callback on the error and returns a new failure with the result of the error - or the same failure if error returned is nil
func (f failure[T, E]) Or(e ErrorHandler[T, E]) Result[T, E] {
	return e(f.err)
}

// Fail creates a failure
func Fail[T, E any](err E) Result[T, E] {
	return failure[T, E]{err: err}
}

// FromTuple creates a Result object from a tupple value, error
func FromTuple[T, E any](val T, err E) Result[T, E] {
	v := reflect.ValueOf(err)
	if v.IsValid() && !v.IsZero() {
		return Fail[T, E](err)
	}
	return Succeed[T, E](val)
}
