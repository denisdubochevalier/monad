package monad

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListLeftIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	f := func(x int) List[int] {
		return NewList([]int{x + 1})
	}
	a := 3

	// Left identity: return a >>= f is the same thing as f applied to a
	leftHandSide := NewList([]int{a}).FlatMap(f).Values()
	rightHandSide := f(a).Values()

	is.True(reflect.DeepEqual(leftHandSide, rightHandSide))
}

func TestListRightIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewList([]int{3})

	// Right identity: m >>= return is no different than just m
	leftHandSide := m.FlatMap(func(x int) List[int] {
		return NewList([]int{x})
	}).Values()
	rightHandSide := m.Values()

	is.True(reflect.DeepEqual(leftHandSide, rightHandSide))
}

func TestListAssociativityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewList([]int{3})
	f := func(x int) List[int] {
		return NewList([]int{x + 1})
	}
	g := func(x int) List[int] {
		return NewList([]int{x * 2})
	}

	// Associativity: (m >>= f) >>= g is just like doing m >>= (\x -> f x >>= g)
	leftHandSide := m.FlatMap(f).FlatMap(g).Values()
	rightHandSide := m.FlatMap(func(x int) List[int] {
		return f(x).FlatMap(g)
	}).Values()

	is.True(reflect.DeepEqual(leftHandSide, rightHandSide))
}
