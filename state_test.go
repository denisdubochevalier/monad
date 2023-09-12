package monad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStateLeftIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	f := func(x int) State[int, int] {
		return NewState[int, int](func(s int) (int, int) {
			return x + 1, s + x
		})
	}
	a := 3
	initialState := 0

	// Test left identity: return a >>= f is the same thing as f applied to a
	lhsValue, lhsState := NewState[int, int](
		func(s int) (int, int) { return a, s },
	).FlatMap(f).
		Run(initialState)
	rhsValue, rhsState := f(a).Run(initialState)

	is.Equal(lhsState, rhsState)
	is.Equal(lhsValue, rhsValue)
}

func TestStateRightIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewState[int, int](func(s int) (int, int) {
		return 3, s + 3
	})
	initialState := 0

	// Test right identity: m >>= return is no different than just m
	lhsValue, lhsState := m.FlatMap(func(x int) State[int, int] {
		return NewState[int, int](func(s int) (int, int) {
			return x, s
		})
	}).Run(initialState)
	rhsValue, rhsState := m.Run(initialState)

	is.Equal(lhsState, rhsState)
	is.Equal(lhsValue, rhsValue)
}

func TestStateAssociativityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewState[int, int](func(s int) (int, int) {
		return 3, s + 3
	})
	f := func(x int) State[int, int] {
		return NewState[int, int](func(s int) (int, int) {
			return x + 1, s + x
		})
	}
	g := func(x int) State[int, int] {
		return NewState[int, int](func(s int) (int, int) {
			return x * 2, s + x
		})
	}
	initialState := 0

	// Test associativity: (m >>= f) >>= g is just like doing m >>= (\x -> f x >>= g)
	lhsValue, lhsState := m.FlatMap(f).FlatMap(g).Run(initialState)
	rhsValue, rhsState := m.FlatMap(func(x int) State[int, int] {
		return f(x).FlatMap(g)
	}).Run(initialState)

	is.Equal(lhsState, rhsState)
	is.Equal(lhsValue, rhsValue)
}
