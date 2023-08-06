package monad

// Predicate represents functions that perform a test on a value
type Predicate[T any] func(T) bool

// Func represents functions that transform a value into another (possibly with a type change)
type Func[V, T any] func(V) T

// Func0 represents a function that takes a value of type T and returns another value of the same type and an error
type Func0[T any] func(T) (T, error)

// Func1 represents a function that takes a value of type T and returns another value of the same type
type Func1[T any] func(T) T

// Func2 represents a function that takes a value of type T and returns a nullable of type T
type Func2[T any] func(T) *T

// ErrFunc represents a function that handles an error
type ErrFunc func(error)
