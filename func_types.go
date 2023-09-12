package monad

// Predicate represents functions that perform a test on a value
type Predicate[T any] func(T) bool

// Failable represents a function that takes a value of type T and returns another value of the same type and an error
type Failable[T any] func(T) (T, error)

// Nilable represents a function that takes a value of type T and returns a nullable of type T
type Nilable[T any] func(T) *T
