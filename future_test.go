package monad

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFutureMonadicLaws(t *testing.T) {
	t.Parallel()

	t.Run("Left identity: return a >>= f == f a", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)
		a := 42
		f := func(x int) Future[int, error] {
			return NewFuture(func() Result[int, error] { return Succeed[int, error](x * 2) })
		}
		future1 := NewFuture(
			func() Result[int, error] { return Succeed[int, error](a) },
		).FlatMap(f).
			Await()
		future2 := f(a).Await()
		is.Equal(future1, future2)
	})

	t.Run("Right identity: m >>= return == m", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		a := 42
		m := NewFuture(func() Result[int, error] { return Succeed[int, error](a) })
		future3 := m.FlatMap(func(x int) Future[int, error] {
			return NewFuture(func() Result[int, error] { return Succeed[int, error](x) })
		}).Await()
		future4 := m.Await()
		is.Equal(future3, future4)
	})

	t.Run("Associativity: (m >>= f) >>= g == m >>= (\\x -> f x >>= g)", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		f := func(x int) Future[int, error] {
			return NewFuture(func() Result[int, error] { return Succeed[int, error](x * 2) })
		}
		g := func(x int) Future[int, error] {
			return NewFuture(func() Result[int, error] { return Succeed[int, error](x - 10) })
		}

		a := 42
		m := NewFuture(func() Result[int, error] { return Succeed[int, error](a) })
		future5 := m.FlatMap(f).FlatMap(g).Await()
		future6 := m.FlatMap(func(x int) Future[int, error] {
			return f(x).FlatMap(g)
		}).Await()
		is.Equal(future5, future6)
	})
}

func TestFutureFailurePropagation(t *testing.T) {
	t.Parallel()

	t.Run("Failure in the first Future propagates through FlatMap", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		failingFuture := NewFuture(
			func() Result[int, error] { return Fail[int, error](errors.New("Oops")) },
		)

		result := failingFuture.FlatMap(func(x int) Future[int, error] {
			return NewFuture(func() Result[int, error] { return Succeed[int, error](42) })
		}).Await()
		is.True(result.Failure())
		is.Equal("Oops", result.Error().Error())
	})

	t.Run("Map error propagation", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		failingFuture := NewFuture(
			func() Result[int, error] { return Fail[int, error](errors.New("Oops")) },
		)

		result2 := failingFuture.Map(func(x int) any {
			return x * 2
		}).Await()
		is.True(result2.Failure())
		is.Equal("Oops", result2.Error().Error())
	})
}
