package monad

// Free represents a Free Monad interface.
type Free[F any, A any] interface {
	// FlatMap composes this Free operation with another, yielding a new Free.
	FlatMap(func(A) Free[F, A]) Free[F, A]

	// Map applies a function to the result of the Free operation.
	Map(func(A) any) Free[F, any]

	// RunFree interprets the Free monad using a given interpreter function.
	RunFree(func(F) A) A
}

// pure represents a concrete implementation for a pure value in Free.
type pure[F any, A any] struct {
	value A
}

// freeOp represents a concrete implementation for a functor operation in Free.
type freeOp[F any, A any] struct {
	functor Functor[F]
	cont    func(F) Free[F, A]
}

// NewPure creates a new Pure wrapped in Free.
func NewPure[F any, A any](value A) Free[F, A] {
	return pure[F, A]{value: value}
}

// NewFreeOp creates a new FreeOp wrapped in Free.
func NewFreeOp[F any, A any](functor Functor[F], cont func(F) Free[F, A]) Free[F, A] {
	return freeOp[F, A]{functor: functor, cont: cont}
}

// FlatMap unwraps the pure value and applies the given function to it, yielding another Free Monad.
func (p pure[F, A]) FlatMap(fn func(A) Free[F, A]) Free[F, A] {
	return fn(p.value)
}

// Map unwraps the pure value and applies the given function to it, yielding another Free Monad.
func (p pure[F, A]) Map(fn func(A) any) Free[F, any] {
	return NewPure[F, any](fn(p.value))
}

// RunFree returns the contained pure value.
func (p pure[F, A]) RunFree(_ func(F) A) A {
	return p.value
}

// FlatMap composes the functor with another Free Monad, effectively chaining operations.
func (op freeOp[F, A]) FlatMap(fn func(A) Free[F, A]) Free[F, A] {
	newCont := func(f F) Free[F, A] {
		return op.cont(f).FlatMap(fn)
	}
	return NewFreeOp(op.functor, newCont)
}

// Map transforms the value inside the Free Monad using a given function, yielding a new Free Monad.
func (op freeOp[F, A]) Map(fn func(A) any) Free[F, any] {
	newCont := func(f F) Free[F, any] {
		return op.cont(f).Map(fn)
	}
	return NewFreeOp(op.functor, newCont)
}

// RunFree interprets the Free Monad using the provided interpreter function to yield a result.
func (op freeOp[F, A]) RunFree(interpreter func(F) A) A {
	next := op.cont(op.functor.Extract())
	return next.RunFree(interpreter)
}
