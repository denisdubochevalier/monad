package monad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaybeLeftIdentityLaw(t *testing.T) {
	// The left identity law essentially posits that return a >>= f is the same as
	// f a, where return a produces a monad encapsulating a, and >>= denotes
	// monadic binding. Since Go doesn't have native monadic syntax, we'll use the
	// FlatMap method to represent this concept.
	t.Parallel()
	is := require.New(t)

	f := func(x int) Maybe[int] { return Some(x + 1) }
	a := 3

	// Test for Just
	is.Equal(Some(a).FlatMap(f).Value(), f(a).Value())

	// Test for Nothing
	is.Equal(None[int]().FlatMap(f).IsNothing(), None[int]().IsNothing())
}

func TestMaybeRightIdentityLaw(t *testing.T) {
	// The right identity law states that for any monadic value m, the expression
	// m >>= return should be equivalent to m. In the context of the Maybe monad,
	// this implies that mapping Some(a) through a function that returns a Maybe
	// encapsulating the original value should yield the original Maybe value
	// back.
	t.Parallel()
	is := require.New(t)

	m := Some(3)

	// Test for Just
	is.Equal(m.FlatMap(Some).Value(), m.Value())

	// Test for Nothing
	is.Equal(None[int]().FlatMap(Some).IsNothing(), None[int]().IsNothing())
}

func TestMaybeAssociativityLaw(t *testing.T) {
	// The associativity law for monads states that the order in which
	// computations are nested should not affect the final result. Specifically,
	// when we have a monadic value m and two monadic functions f and g, the law
	// states that:
	//
	// First binding m with f and then binding the result with g should be the
	// same as binding m with a function that is the result of first binding f
	// with g.
	//
	// In more concrete terms, we can express it as:
	//
	// (m >>= f) >>= g should be equal to m >>= (x -> f(x) >>= g)
	t.Parallel()
	is := require.New(t)

	m := Some(3)
	f := func(x int) Maybe[int] { return Some(x + 1) }
	g := func(x int) Maybe[int] { return Some(x * 2) }

	// Test for Just
	leftHandSide := m.FlatMap(f).FlatMap(g)
	rightHandSide := m.FlatMap(func(x int) Maybe[int] { return f(x).FlatMap(g) })
	is.Equal(leftHandSide.Value(), rightHandSide.Value())

	// Test for Nothing
	m = None[int]()
	leftHandSide = m.FlatMap(f).FlatMap(g)
	rightHandSide = m.FlatMap(func(x int) Maybe[int] { return f(x).FlatMap(g) })
	is.Equal(leftHandSide.IsNothing(), rightHandSide.IsNothing())
}

// Additional tests for methods like OrElse, IsJust, IsNothing etc.
func TestMaybeMethods(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	just := Some(5)
	nothing := None[int]()

	// Test IsJust and IsNothing
	is.True(just.IsJust())
	is.False(just.IsNothing())
	is.False(nothing.IsJust())
	is.True(nothing.IsNothing())

	// Test OrElse
	is.Equal(5, just.OrElse(10))
	is.Equal(10, nothing.OrElse(10))

	// Test Filter
	is.Equal(just, just.Filter(Predicate[int](func(a int) bool { return a == 5 })))
	is.Equal(nothing, just.Filter(Predicate[int](func(a int) bool { return a != 5 })))
	is.Equal(nothing, nothing.Filter(Predicate[int](func(a int) bool { return a == 5 })))

	// Test Value
	is.Equal(5, just.Value())
	is.Equal(0, nothing.Value()) // should be zero value for type int
}
