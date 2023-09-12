package monad_test

import (
	"testing"

	"github.com/denisdubochevalier/monad"
	"github.com/stretchr/testify/require"
)

func TestValidationMap(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		v := monad.NewValid[string, int](5)
		newV := v.Map(func(x int) any { return x * 2 })
		is.True(newV.Valid())
		is.Equal(10, newV.Value())
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		inv := monad.NewInvalid[string, int]([]string{"error1", "error2"})
		newInv := inv.Map(func(x int) any { return x * 2 })
		is.False(newInv.Valid())
		is.Equal([]string{"error1", "error2"}, newInv.Errors())
	})
}

func TestValidationFlatMap(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		v := monad.NewValid[string, int](5)
		newV := v.FlatMap(func(x int) monad.Validation[string, int] {
			return monad.NewValid[string, int](x * 2)
		})
		is.True(newV.Valid())
		is.Equal(10, newV.Value())
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		inv := monad.NewInvalid[string, int]([]string{"error1", "error2"})
		newInv := inv.FlatMap(func(x int) monad.Validation[string, int] {
			return monad.NewValid[string, int](x * 2)
		})
		is.False(newInv.Valid())
		is.Equal([]string{"error1", "error2"}, newInv.Errors())
	})
}

func TestValidationMonadicLaws(t *testing.T) {
	t.Parallel()

	t.Run("Left Identity: return a >>= f is equivalent to f(a)", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		a := 5
		f := func(x int) monad.Validation[string, int] {
			return monad.NewValid[string, int](x * 2)
		}
		is.Equal(f(a).Value(), monad.NewValid[string, int](a).FlatMap(f).Value())
	})

	t.Run("Right Identity: m >>= return is equivalent to m", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		a := 5
		m := monad.NewValid[string, int](a)
		is.Equal(m.Value(), m.FlatMap(func(x int) monad.Validation[string, int] {
			return monad.NewValid[string, int](x)
		}).Value())
	})

	t.Run(
		"Associativity: (m >>= f) >>= g is equivalent to m >>= (\\x -> f x >>= g)",
		func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			f := func(x int) monad.Validation[string, int] {
				return monad.NewValid[string, int](x * 2)
			}
			g := func(x int) monad.Validation[string, int] {
				return monad.NewValid[string, int](x + 3)
			}

			a := 5
			m := monad.NewValid[string, int](a)

			is.Equal(
				m.FlatMap(f).FlatMap(g).Value(),
				m.FlatMap(func(x int) monad.Validation[string, int] {
					return f(x).FlatMap(g)
				}).Value(),
			)
		},
	)
}
