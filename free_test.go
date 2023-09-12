package monad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Our dummy Functor implementation for testing purposes.
type TestFunctor struct {
	Value int
}

func (tf TestFunctor) Map(f func(any) any) Functor[TestFunctor] {
	return TestFunctor{Value: f(tf.Value).(int)}
}

func (tf TestFunctor) Extract() TestFunctor {
	return tf
}

// Functions to interpret TestFunctor
func interpreterAny(tf TestFunctor) any {
	return tf.Value
}

func interpreterInt(tf TestFunctor) int {
	return tf.Value
}

func TestFreeMonad(t *testing.T) {
	t.Parallel()

	t.Run("TestPure", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		m := NewPure[TestFunctor, any](42)
		is.Equal(42, m.RunFree(interpreterAny))
	})

	t.Run("TestMap", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		m := NewPure[TestFunctor, int](21).Map(func(x int) any {
			return x * 2
		})
		is.Equal(42, m.RunFree(interpreterAny))
	})

	t.Run("TestFlatMap", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		m := NewPure[TestFunctor, int](21).FlatMap(func(x int) Free[TestFunctor, int] {
			return NewPure[TestFunctor, int](x * 2)
		})
		is.Equal(42, m.RunFree(interpreterInt))
	})

	t.Run("TestFreeOp", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		op := NewFreeOp[TestFunctor, int](
			TestFunctor{Value: 21},
			func(f TestFunctor) Free[TestFunctor, int] {
				return NewPure[TestFunctor, int](f.Value * 2)
			},
		)
		is.Equal(42, op.RunFree(interpreterInt))
	})

	t.Run("TestMonadLaws", func(t *testing.T) {
		t.Parallel()

		t.Run("LeftIdentity", func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			a := 21
			f := func(x int) Free[TestFunctor, int] {
				return NewPure[TestFunctor, int](x * 2)
			}

			left := NewPure[TestFunctor, int](a).FlatMap(f)
			right := f(a)
			is.Equal(left.RunFree(interpreterInt), right.RunFree(interpreterInt))
		})

		t.Run("RightIdentity", func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			a := 21
			m := NewPure[TestFunctor, int](a)
			right := m.FlatMap(NewPure[TestFunctor, int])
			is.Equal(m.RunFree(interpreterInt), right.RunFree(interpreterInt))
		})

		t.Run("Associativity", func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			m := NewPure[TestFunctor, int](21)
			f := func(x int) Free[TestFunctor, int] {
				return NewPure[TestFunctor, int](x * 2)
			}
			g := func(x int) Free[TestFunctor, int] {
				return NewPure[TestFunctor, int](x + 1)
			}

			left := m.FlatMap(f).FlatMap(g)
			right := m.FlatMap(func(x int) Free[TestFunctor, int] {
				return f(x).FlatMap(g)
			})
			is.Equal(left.RunFree(interpreterInt), right.RunFree(interpreterInt))
		})
	})
}
