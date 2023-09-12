package monad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEitherLeftIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	f := func(x int) Either[int] { return NewRVal(x + 1) }
	a := 3

	// Test for Right
	is.Equal(NewRVal(a).FMap(f).Value(), f(a).Value())

	// Test for Left
	is.Equal(NewLVal(a).FMap(f).Left(), NewLVal(a).Left())
}

func TestEitherRightIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewRVal(3)

	// Test for Right
	is.Equal(m.FMap(NewRVal).Value(), m.Value())

	// Test for Left
	m = NewLVal(3)
	is.Equal(m.FMap(NewRVal).Left(), m.Left())
}

func TestEitherAssociativityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := NewRVal(3)
	f := func(x int) Either[int] { return NewRVal(x + 1) }
	g := func(x int) Either[int] { return NewRVal(x * 2) }

	// Test for Right
	leftHandSide := m.FMap(f).FMap(g)
	rightHandSide := m.FMap(func(x int) Either[int] { return f(x).FMap(g) })
	is.Equal(leftHandSide.Value(), rightHandSide.Value())

	// Test for Left
	m = NewLVal(3)
	leftHandSide = m.FMap(f).FMap(g)
	rightHandSide = m.FMap(func(x int) Either[int] { return f(x).FMap(g) })
	is.Equal(leftHandSide.Left(), rightHandSide.Left())
}
