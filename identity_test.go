package monad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIdentityLeftIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	f := func(x int) Identity[int] {
		return NewIdentity(x + 1)
	}
	a := 3

	// Test left identity: return a >>= f is the same as f applied to a
	leftHandSide := NewIdentity(a).FlatMap(f).Value()
	rightHandSide := f(a).Value()

	is.Equal(leftHandSide, rightHandSide)
}

func TestIdentityRightIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewIdentity(3)

	// Test right identity: m >>= return is no different than just m
	leftHandSide := m.FlatMap(func(x int) Identity[int] {
		return NewIdentity(x)
	}).Value()
	rightHandSide := m.Value()

	is.Equal(leftHandSide, rightHandSide)
}

func TestIdentityAssociativityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewIdentity(3)
	f := func(x int) Identity[int] {
		return NewIdentity(x + 1)
	}
	g := func(x int) Identity[int] {
		return NewIdentity(x * 2)
	}

	// Test associativity: (m >>= f) >>= g is just like doing m >>= (\x -> f x >>= g)
	leftHandSide := m.FlatMap(f).FlatMap(g).Value()
	rightHandSide := m.FlatMap(func(x int) Identity[int] {
		return f(x).FlatMap(g)
	}).Value()

	is.Equal(leftHandSide, rightHandSide)
}
