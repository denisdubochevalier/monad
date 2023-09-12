package monad

import "sync"

// Future represents a monadic interface for future asynchronous computations.
//
// T is the type of the value that the Future operation produces.
// E is the type of the error that can occur.
type Future[T, E any] interface {
	// Await blocks until the Future is either fulfilled or failed, then returns a Result.
	Await() Result[T, E]

	// Map applies a function to the result of the Future operation, yielding a new Future.
	Map(func(T) any) Future[any, E]

	// FlatMap composes this Future operation with another, yielding a new Future.
	FlatMap(func(T) Future[T, E]) Future[T, E]
}

// future is a concrete implementation of the Future interface.
type future[T, E any] struct {
	action func() Result[T, E]
	result Result[T, E]
	once   sync.Once
}

// NewFuture constructs a new Future Monad.
func NewFuture[T, E any](action func() Result[T, E]) Future[T, E] {
	return &future[T, E]{action: action}
}

// Await waits for the Future to be completed and returns the Result.
func (f *future[T, E]) Await() Result[T, E] {
	f.once.Do(func() {
		f.result = f.action()
	})
	return f.result
}

// Map applies a function to the result of the Future operation.
func (f *future[T, E]) Map(transform func(T) any) Future[any, E] {
	return NewFuture[any, E](func() Result[any, E] {
		res := f.Await()
		if res.Failure() {
			return Fail[any, E](res.Error())
		}
		return Succeed[any, E](transform(res.Value()))
	})
}

// FlatMap composes this Future operation with another.
func (f *future[T, E]) FlatMap(compose func(T) Future[T, E]) Future[T, E] {
	return NewFuture[T, E](func() Result[T, E] {
		res := f.Await()
		if res.Failure() {
			return Fail[T, E](res.Error())
		}
		nextFuture := compose(res.Value())
		return nextFuture.Await()
	})
}
