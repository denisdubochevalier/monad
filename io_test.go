package monad

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIOMonadicLaws(t *testing.T) {
	t.Parallel()

	t.Run("Left identity: return a >>= f == f a", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)
		a := 42
		f := func(x int) IO[int, error] {
			return NewIO(func() Result[int, error] { return Succeed[int, error](x * 2) })
		}
		io1 := NewIO(
			func() Result[int, error] { return Succeed[int, error](a) },
		).FlatMap(f).
			Perform()
		io2 := f(a).Perform()
		is.Equal(io1, io2)
	})

	t.Run("Right identity: m >>= return == m", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		a := 42
		m := NewIO(func() Result[int, error] { return Succeed[int, error](a) })
		io3 := m.FlatMap(func(x int) IO[int, error] {
			return NewIO(func() Result[int, error] { return Succeed[int, error](x) })
		}).Perform()
		io4 := m.Perform()
		is.Equal(io3, io4)
	})

	t.Run("Associativity: (m >>= f) >>= g == m >>= (\\x -> f x >>= g)", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		f := func(x int) IO[int, error] {
			return NewIO(func() Result[int, error] { return Succeed[int, error](x * 2) })
		}
		g := func(x int) IO[int, error] {
			return NewIO(func() Result[int, error] { return Succeed[int, error](x - 10) })
		}

		a := 42
		m := NewIO(func() Result[int, error] { return Succeed[int, error](a) })
		io5 := m.FlatMap(f).FlatMap(g).Perform()
		io6 := m.FlatMap(func(x int) IO[int, error] {
			return f(x).FlatMap(g)
		}).Perform()
		is.Equal(io5, io6)
	})
}

func TestIOFailurePropagation(t *testing.T) {
	t.Parallel()

	failingIO := NewIO(func() Result[int, error] { return Fail[int, error](errors.New("Oops")) })

	t.Run("Failure in the first IO propagates through FlatMap", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		result := failingIO.FlatMap(func(x int) IO[int, error] {
			return NewIO(func() Result[int, error] { return Succeed[int, error](42) })
		}).Perform()
		is.True(result.Failure())
		is.Equal("Oops", result.Error().Error())
	})

	t.Run("Map error propagation", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		result2 := failingIO.Map(func(x int) any {
			return x * 2
		}).Perform()
		is.True(result2.Failure())
		is.Equal("Oops", result2.Error().Error())
	})
}
