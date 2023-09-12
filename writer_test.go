package monad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriterLeftIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	f := func(x int) Writer[int, int] {
		return NewWriter(x+1, x)
	}
	initialValue, initialOutput := 3, 0

	// Left identity: return a >>= f is the same as f a
	leftVal, leftOut := NewWriter(initialValue, initialOutput).FlatMap(f).Run()
	rightVal, rightOut := f(initialValue).Run()

	is.Equal(leftVal, rightVal)
	is.Equal(leftOut, rightOut)
}

func TestWriterRightIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	// Initial Writer Monad
	m := NewWriter(3, 0)

	// Right identity: m >>= return is no different than just m
	leftVal, leftOut := m.FlatMap(func(x int) Writer[int, int] {
		return NewWriter(x, 0)
	}).Run()

	rightVal, rightOut := m.Run()

	is.Equal(leftVal, rightVal)
	is.Equal(leftOut, rightOut)
}

func TestWriterAssociativityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	// Initial Writer Monad
	m := NewWriter(3, 0)

	f := func(x int) Writer[int, int] {
		return NewWriter(x+1, x)
	}

	g := func(x int) Writer[int, int] {
		return NewWriter(x*2, x)
	}

	// Associativity: (m >>= f) >>= g is just like doing m >>= (\x -> f x >>= g)
	leftVal, leftOut := m.FlatMap(f).FlatMap(g).Run()
	rightVal, rightOut := m.FlatMap(func(x int) Writer[int, int] {
		return f(x).FlatMap(g)
	}).Run()

	is.Equal(leftVal, rightVal)
	is.Equal(leftOut, rightOut)
}
