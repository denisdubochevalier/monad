package monad

// Result represents a result of an operation that can fail
type Result[T any] interface {
	Error() error
	Value() T
	Failure() bool
	Success() bool
	FBind(func(T) Result[T]) Result[T]
	Bind(Failable[T]) Result[T]
	FMap(Transformable[T]) Result[T]
	Or(ErrorHandler) Result[T]
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

// FBind executes the function on the underlying success and returns its Result response
func (s Success[T]) FBind(f func(T) Result[T]) Result[T] {
	return f(s.val)
}

// Bind executes the callback function and returns a Result object
func (s Success[T]) Bind(f Failable[T]) Result[T] {
	val, err := f(s.val)
	if err != nil {
		return Fail[T](err)
	}
	return Succeed(val)
}

// FMap executes the callback function and returns a success with the value changed to the result of the callback
func (s Success[T]) FMap(t Transformable[T]) Result[T] {
	return Succeed(t(s.val))
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

// FBind returns the current failure
func (f Failure[T]) FBind(_ func(T) Result[T]) Result[T] {
	return f
}

// Bind returns the Failure
func (f Failure[T]) Bind(_ Failable[T]) Result[T] {
	return f
}

// FMap returns the original failure
func (f Failure[T]) FMap(_ Transformable[T]) Result[T] {
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
