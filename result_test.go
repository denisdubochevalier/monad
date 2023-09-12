package monad

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResultLeftIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	f := func(x int) Result[int] { return Succeed(x + 1) }
	a := 3

	// Test for Success
	is.Equal(Succeed(a).FMap(f).Value(), f(a).Value())

	// Test for Failure
	is.Equal(
		Fail[int](errors.New("test")).FMap(f).Failure(),
		Fail[int](errors.New("test")).Failure(),
	)
}

func TestResultRightIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := Succeed(3)

	// Test for Success
	is.Equal(m.FMap(Succeed).Value(), m.Value())

	// Test for Failure
	m = Fail[int](errors.New("test"))
	is.Equal(m.FMap(Succeed).Failure(), m.Failure())
}

func TestResultAssociativityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := Succeed(3)
	f := func(x int) Result[int] { return Succeed(x + 1) }
	g := func(x int) Result[int] { return Succeed(x * 2) }

	// Test for Success
	leftHandSide := m.FMap(f).FMap(g)
	rightHandSide := m.FMap(func(x int) Result[int] { return f(x).FMap(g) })
	is.Equal(leftHandSide.Value(), rightHandSide.Value())

	// Test for Failure
	m = Fail[int](errors.New("test"))
	leftHandSide = m.FMap(f).FMap(g)
	rightHandSide = m.FMap(func(x int) Result[int] { return f(x).FMap(g) })
	is.Equal(leftHandSide.Failure(), rightHandSide.Failure())
}
