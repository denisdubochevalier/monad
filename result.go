package monad

// Result represents a result of an operation that can fail
type Result[T any] interface {
	Error() error
	Value() T
	Bind(Func0[T]) Result[T]
	Failure() bool
	Success() bool
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

// Bind executes the callback function and returns a Result object
func (s Success[T]) Bind(f Func0[T]) Result[T] {
	val, err := f(s.val)
	if err != nil {
		return Fail[T](err)
	}
	return Succeed(val)
}

// Failure always return false for a Success
func (s Success[_]) Failure() bool {
	return false
}

// Success always return true for a Success
func (s Success[_]) Success() bool {
	return true
}

// Succeed creates a Success
func Succeed[T any](val T) Success[T] {
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

// Bind returns the Failure
func (f Failure[T]) Bind(_ Func0[T]) Result[T] {
	return f
}

// Failure always returns true for a Failure
func (f Failure[_]) Failure() bool {
	return true
}

// Success always returns false for a Failure
func (f Failure[_]) Success() bool {
	return false
}

// Fail creates a Failure
func Fail[T any](err error) Failure[T] {
	return Failure[T]{err: err}
}

// FromTuple creates a Result object from a tupple value, error
func FromTuple[T any](val T, err error) Result[T] {
	if err != nil {
		return Fail[T](err)
	}
	return Succeed(val)
}
