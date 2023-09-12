package monad

// State is a generic interface for representing a stateful computation.
// It wraps a value of type T and a state of type S, along with the capability to
// transform itself using Map and FlatMap operations.
type State[S, T any] interface {
	// Value returns the value of type T wrapped by the State monad.
	Value() T

	// State returns the state of type S stored in the State monad.
	State() S

	// Run performs the stateful computation and returns the resulting State
	Run(S) (T, S)

	// Map applies a given function to the wrapped value without affecting the state.
	// It returns a new State monad wrapping the transformed value.
	Map(func(T) any) State[S, any]

	// FlatMap applies a given function that returns a new State monad.
	// It effectively combines the state transformations of both the original and the new monad.
	FlatMap(func(T) State[S, T]) State[S, T]
}

// Stateful is a concrete implementation of the State interface.
// It uses an internal function, runner, to define its stateful computation.
type Stateful[S, T any] struct {
	state  S              // The current state
	val    T              // The value wrapped by the State monad
	runner func(S) (T, S) // Function to perform the stateful computation
}

// NewState creates a new State monad given a function that represents a stateful computation.
func NewState[S, T any](f func(S) (T, S)) State[S, T] {
	return Stateful[S, T]{runner: f}
}

// Value returns the value wrapped by the Stateful monad.
func (s Stateful[S, T]) Value() T {
	return s.val
}

// State returns the current state stored in the Stateful monad.
func (s Stateful[S, T]) State() S {
	return s.state
}

// Run performs the stateful computation defined by runner and returns the resulting value and state.
// Run performs the stateful computation defined by runner and returns the resulting state.
func (s Stateful[S, T]) Run(state S) (T, S) {
	return s.runner(state)
}

// Map applies a given function to the wrapped value without affecting the state.
// It returns a new State monad with the transformed value.
// Map applies a given function to the wrapped value without affecting the state.
func (s Stateful[S, T]) Map(f func(T) any) State[S, any] {
	return NewState[S, any](func(state S) (any, S) {
		val, newState := s.Run(state)
		return f(val), newState
	})
}

// FlatMap applies a given function that returns a new State monad.
// It effectively combines the state transformations of both the original and the new monad.
func (s Stateful[S, T]) FlatMap(f func(T) State[S, T]) State[S, T] {
	return NewState[S, T](func(state S) (T, S) {
		val, newState := s.runner(state) // Run the original state
		newMonad := f(val)               // Create a new State monad via the function f
		if newMonad, ok := newMonad.(Stateful[S, T]); ok {
			return newMonad.runner(newState) // Run the new State monad and return its result
		}
		// handle error if the cast is not successful, return zero values or throw panic, etc.
		return *(new(T)), newState
	})
}
