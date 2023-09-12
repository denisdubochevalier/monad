package monad

// List represents a generic interface for the List monad. It is a container
// that holds a slice of type []T and provides monadic methods to perform
// transformations on the contained elements.
type List[T any] interface {
	// Values returns the slice of encapsulated values.
	Values() []T

	// Map applies a transformation to each element in the List.
	Map(func(T) any) List[any]

	// FlatMap applies a transformation that returns a new List and concatenates
	// all resulting Lists into a single List.
	FlatMap(func(T) List[T]) List[T]
}

// list is a concrete implementation of the List interface.
type list[T any] struct {
	values []T
}

// NewList creates a new List given a slice of initial values.
func NewList[T any](values []T) List[T] {
	return list[T]{values: values}
}

// Values retrieves the encapsulated slice from a list.
func (l list[T]) Values() []T {
	return l.values
}

// Map applies a transformation to each element in the list and returns a new List.
func (l list[T]) Map(f func(T) any) List[any] {
	var newValues []any
	for _, v := range l.values {
		newValues = append(newValues, f(v))
	}
	return NewList[any](newValues)
}

// FlatMap transforms each element in the list to a new List and flattens the result.
func (l list[T]) FlatMap(f func(T) List[T]) List[T] {
	var newValues []T
	for _, v := range l.values {
		newValues = append(newValues, f(v).Values()...)
	}
	return NewList[T](newValues)
}
