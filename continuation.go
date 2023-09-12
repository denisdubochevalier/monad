package monad

import (
	"context"
)

// Continuation represents a monadic interface for continuation-passing style
// (CPS) computations.
//
// T is the type of the value that the computation produces.
type Continuation[T any] interface {
	// Run executes the Continuation, blocking until either a result is ready or
	// the context is cancelled. In the case of cancellation or timeout, the error
	// is taken directly from context.Err(), and is therefore of type `error`.
	Run(ctx context.Context) Result[T, error]

	// Map applies a function to the computed value, yielding a new Continuation.
	Map(func(T) any) Continuation[any]

	// FlatMap composes this computation with another, yielding a new
	// Continuation.
	FlatMap(func(T) Continuation[T]) Continuation[T]
}

// continuation is a concrete implementation of the Continuation interface.
type continuation[T any] struct {
	// runFunc is a function that performs the actual computation.
	// It returns a Result monad containing either the computed value or an error.
	runFunc func(ctx context.Context) Result[T, error]
}

// NewContinuation creates a new Continuation that wraps the given computation
// function.
func NewContinuation[T any](
	runFunc func(ctx context.Context) Result[T, error],
) Continuation[T] {
	return &continuation[T]{runFunc: runFunc}
}

// Run executes the encapsulated computation and returns a Result monad.
func (c continuation[T]) Run(ctx context.Context) Result[T, error] {
	done := make(chan struct{})
	var res Result[T, error]

	go func() {
		res = c.runFunc(ctx)
		close(done)
	}()

	select {
	case <-done:
		return res
	case <-ctx.Done():
		return Fail[T, error](ctx.Err())
	}
}

// Map applies a function to the value that this Continuation produces.
func (c continuation[T]) Map(f func(T) any) Continuation[any] {
	return continuation[any]{runFunc: func(ctx context.Context) Result[any, error] {
		res := c.Run(ctx)
		if res.Failure() {
			return Fail[any, error](res.Error())
		}
		return Succeed[any, error](f(res.Value()))
	}}
}

// FlatMap composes this Continuation with another.
func (c continuation[T]) FlatMap(f func(T) Continuation[T]) Continuation[T] {
	return continuation[T]{runFunc: func(ctx context.Context) Result[T, error] {
		res := c.Run(ctx)
		if res.Failure() {
			return Fail[T, error](res.Error())
		}
		nextC := f(res.Value())
		return nextC.Run(ctx)
	}}
}
