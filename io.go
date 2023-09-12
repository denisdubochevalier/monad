package monad

// IO represents a monadic interface for IO operations.
//
// T is the type of the value that the IO operation produces.
type IO[T, E any] interface {
	// Perform executes the IO operation, producing either a result or an error.
	Perform() Result[T, E]

	// Map applies a function to the result of the IO operation, yielding a new IO.
	Map(func(T) any) IO[any, E]

	// FlatMap composes this IO operation with another, yielding a new IO.
	FlatMap(func(T) IO[T, E]) IO[T, E]
}

// io is a concrete implementation of the IO interface.
type io[T, E any] struct {
	// ioFunc is a function that performs the actual IO operation.
	// It returns a Result monad containing either the computed value or an error.
	ioFunc func() Result[T, E]
}

// NewIO constructs a new IO monad.
func NewIO[T, E any](ioFunc func() Result[T, E]) IO[T, E] {
	return io[T, E]{ioFunc: ioFunc}
}

// Perform executes the encapsulated IO operation and returns a Result.
func (i io[T, E]) Perform() Result[T, E] {
	return i.ioFunc()
}

// Map applies a function to the result of the IO operation.
func (i io[T, E]) Map(f func(T) any) IO[any, E] {
	return io[any, E]{ioFunc: func() Result[any, E] {
		res := i.Perform()
		if res.Failure() {
			return Fail[any, E](res.Error())
		}
		return Succeed[any, E](f(res.Value()))
	}}
}

// FlatMap composes this IO operation with another.
func (i io[T, E]) FlatMap(f func(T) IO[T, E]) IO[T, E] {
	return io[T, E]{ioFunc: func() Result[T, E] {
		res := i.Perform()
		if res.Failure() {
			return Fail[T, E](res.Error())
		}
		nextIO := f(res.Value())
		return nextIO.Perform()
	}}
}
