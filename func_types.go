package monad

// Predicate represents functions that perform a test on a value
type Predicate[T any] func(T) bool

// Failable represents a function that takes a value of type T and returns another value of the same type and an error
type Failable[T any] func(T) (T, error)

// Transformable represents a function that takes a value of type T and returns another value of the same type
type Transformable[T any] func(T) T

// Nilable represents a function that takes a value of type T and returns a nullable of type T
type Nilable[T any] func(T) *T

// ErrorHandler represents a function that handles an error
type ErrorHandler func(error)
