package monad

// Writer is a generic interface for representing a writer monad.
// It wraps a value of type T and an output of type W, offering methods
// to perform transformations using Map and FlatMap.
type Writer[W, T any] interface {
	// Value returns the encapsulated value of type T.
	Value() T

	// Output returns the writer output of type W.
	Output() W

	// Run executes the writer and returns the encapsulated value and the writer output.
	Run() (T, W)

	// Map applies a transformation to the encapsulated value and returns
	// a new Writer monad with the transformed value. The output remains unchanged.
	Map(func(T) any) Writer[W, any]

	// FlatMap applies a transformation function that returns a new Writer monad.
	// It combines the value and the output of both the original and the new Writer monad.
	FlatMap(func(T) Writer[W, T]) Writer[W, T]
}

// writer is a concrete implementation of the Writer interface.
// It holds an encapsulated value and a writer output, along with a writer function
// to define the computation.
type writer[W, T any] struct {
	value  T                 // The encapsulated value
	output W                 // The writer output
	writer func(T, W) (T, W) // The writer function to define the computation
}

// NewWriter constructs a new Writer monad given an initial value and initial output.
func NewWriter[W, T any](initialValue T, initialOutput W) Writer[W, T] {
	return writer[W, T]{
		value:  initialValue,
		output: initialOutput,
		writer: func(value T, output W) (T, W) {
			return value, output
		},
	}
}

// Value returns the encapsulated value of the writer monad.
func (w writer[W, T]) Value() T {
	return w.value
}

// Output returns the writer output of the writer monad.
func (w writer[W, T]) Output() W {
	return w.output
}

// Run performs the writer computation and returns the encapsulated value and writer output.
func (w writer[W, T]) Run() (T, W) {
	return w.writer(w.value, w.output)
}

// Map applies a given function to transform the encapsulated value,
// while keeping the output unchanged. It returns a new Writer monad with the transformed value.
func (w writer[W, T]) Map(f func(T) any) Writer[W, any] {
	newValue := f(w.value)
	return writer[W, any]{
		value:  newValue,
		output: w.output,
		writer: func(anyValue any, anyOutput W) (any, W) {
			// Convert anyValue back to type T and run the original writer function
			convertedValue := anyValue.(T)
			_, newOutput := w.writer(convertedValue, anyOutput)
			return newValue, newOutput
		},
	}
}

// FlatMap applies a given function that returns a new Writer monad.
// It merges the value and the output of the original and new Writer monad
// into a new Writer monad.
func (w writer[W, T]) FlatMap(f func(T) Writer[W, T]) Writer[W, T] {
	newWriter := f(w.value)
	newValue, newOutput := newWriter.Run()
	return writer[W, T]{
		value:  newValue,
		output: newOutput,
		writer: w.writer,
	}
}
