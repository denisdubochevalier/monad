package monad

// Reader is a generic interface representing a Reader monad.
// It wraps a computation that reads from a shared environment of type E and produces a value of type T.
type Reader[E, T any] interface {
	// Run executes the Reader computation with an environment and returns the resulting value.
	Run(E) T

	// Map applies a transformation to the Reader's result and returns a new Reader.
	Map(func(T) any) Reader[E, any]

	// FlatMap applies a transformation that returns a new Reader and returns the combined Reader.
	FlatMap(func(T) Reader[E, T]) Reader[E, T]
}

// reader is a concrete implementation of the Reader interface.
// It holds a function that defines the computation to be run with an environment.
type reader[E, T any] struct {
	computation func(E) T
}

// NewReader constructs a new Reader monad given a computation function.
func NewReader[E, T any](computation func(E) T) Reader[E, T] {
	return reader[E, T]{computation: computation}
}

// Run executes the Reader computation with the provided environment and returns the resulting value.
func (r reader[E, T]) Run(env E) T {
	return r.computation(env)
}

// Map applies a given function to transform the Reader's result into a new type.
// It returns a new Reader that wraps the new computation.
func (r reader[E, T]) Map(f func(T) any) Reader[E, any] {
	return NewReader[E, any](func(env E) any {
		return f(r.Run(env))
	})
}

// FlatMap applies a given function that returns a new Reader monad.
// It returns a new Reader that wraps the combined computation.
func (r reader[E, T]) FlatMap(f func(T) Reader[E, T]) Reader[E, T] {
	return NewReader[E, T](func(env E) T {
		newReader := f(r.Run(env))
		return newReader.Run(env)
	})
}
