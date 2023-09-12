package monad

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResultLeftIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	f := func(x int) Result[int, error] { return Succeed[int, error](x + 1) }
	a := 3

	// Test for Success
	is.Equal(Succeed[int, error](a).FlatMap(f).Value(), f(a).Value())

	// Test for Failure
	is.Equal(
		Fail[int](errors.New("test")).FlatMap(f).Failure(),
		Fail[int](errors.New("test")).Failure(),
	)
}

func TestResultRightIdentityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := Succeed[int, error](3)

	// Test for Success
	is.Equal(m.FlatMap(Succeed).Value(), m.Value())

	// Test for Failure
	m = Fail[int](errors.New("test"))
	is.Equal(m.FlatMap(Succeed).Failure(), m.Failure())
}

func TestResultAssociativityLaw(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	m := Succeed[int, error](3)
	f := func(x int) Result[int, error] { return Succeed[int, error](x + 1) }
	g := func(x int) Result[int, error] { return Succeed[int, error](x * 2) }

	// Test for Success
	leftHandSide := m.FlatMap(f).FlatMap(g)
	rightHandSide := m.FlatMap(func(x int) Result[int, error] { return f(x).FlatMap(g) })
	is.Equal(leftHandSide.Value(), rightHandSide.Value())

	// Test for Failure
	m = Fail[int](errors.New("test"))
	leftHandSide = m.FlatMap(f).FlatMap(g)
	rightHandSide = m.FlatMap(func(x int) Result[int, error] { return f(x).FlatMap(g) })
	is.Equal(leftHandSide.Failure(), rightHandSide.Failure())
}
