package monad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReaderLeftIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	f := func(x int) Reader[int, int] {
		return NewReader[int, int](func(env int) int {
			return x + env
		})
	}

	a := 3
	env := 2

	// Left Identity Law: return a >>= f is the same as f a
	leftHandSide := NewReader[int, int](func(e int) int { return a }).FlatMap(f).Run(env)
	rightHandSide := f(a).Run(env)

	is.Equal(leftHandSide, rightHandSide)
}

func TestReaderRightIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewReader[int, int](func(env int) int {
		return env + 3
	})
	env := 2

	// Right Identity Law: m >>= return is no different than just m
	leftHandSide := m.FlatMap(func(x int) Reader[int, int] {
		return NewReader[int, int](func(e int) int { return x })
	}).Run(env)

	rightHandSide := m.Run(env)

	is.Equal(leftHandSide, rightHandSide)
}

func TestReaderAssociativityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewReader[int, int](func(env int) int {
		return env + 3
	})
	f := func(x int) Reader[int, int] {
		return NewReader[int, int](func(env int) int {
			return x + env
		})
	}
	g := func(x int) Reader[int, int] {
		return NewReader[int, int](func(env int) int {
			return x * env
		})
	}
	env := 2

	// Associativity Law: (m >>= f) >>= g is just like doing m >>= (\x -> f x >>= g)
	leftHandSide := m.FlatMap(f).FlatMap(g).Run(env)
	rightHandSide := m.FlatMap(func(x int) Reader[int, int] {
		return f(x).FlatMap(g)
	}).Run(env)

	is.Equal(leftHandSide, rightHandSide)
}
