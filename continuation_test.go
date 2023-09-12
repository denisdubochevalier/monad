package monad

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestContinuationMonadLaws(t *testing.T) {
	t.Parallel()

	// Identity: return a >>= f ≡ f a
	t.Run("Left identity", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)
		val := 42
		f := func(x int) Continuation[int] {
			return NewContinuation[int](func(ctx context.Context) Result[int, error] {
				return Succeed[int, error](x * 2)
			})
		}

		cont1 := NewContinuation[int](func(ctx context.Context) Result[int, error] {
			return Succeed[int, error](val)
		}).FlatMap(f)

		cont2 := f(val)

		is.Equal(cont1.Run(context.Background()).Value(),
			cont2.Run(context.Background()).Value())
	})

	// Identity: m >>= return ≡ m
	t.Run("Right identity", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)
		cont := NewContinuation[int](func(ctx context.Context) Result[int, error] {
			return Succeed[int, error](42)
		})

		flatMapped := cont.FlatMap(func(x int) Continuation[int] {
			return NewContinuation[int](func(ctx context.Context) Result[int, error] {
				return Succeed[int, error](x)
			})
		})

		is.Equal(
			cont.Run(context.Background()).Value(),
			flatMapped.Run(context.Background()).Value(),
		)
	})

	// Associativity: (m >>= f) >>= g ≡ m >>= (\x -> f x >>= g)
	t.Run("Associativity", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)
		cont := NewContinuation[int](func(ctx context.Context) Result[int, error] {
			return Succeed[int, error](42)
		})

		f := func(x int) Continuation[int] {
			return NewContinuation[int](func(ctx context.Context) Result[int, error] {
				return Succeed[int, error](x * 2)
			})
		}

		g := func(x int) Continuation[int] {
			return NewContinuation[int](func(ctx context.Context) Result[int, error] {
				return Succeed[int, error](x + 1)
			})
		}

		left := cont.FlatMap(f).FlatMap(g)
		right := cont.FlatMap(func(x int) Continuation[int] {
			return f(x).FlatMap(g)
		})

		is.Equal(left.Run(context.Background()).Value(),
			right.Run(context.Background()).Value())
	})
}

func TestContextCancellation(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()
	is := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	cont := NewContinuation[int](func(ctx context.Context) Result[int, error] {
		// Simulating a long-running computation
		select {
		case <-time.After(5 * time.Second):
			return Succeed[int, error](42)
		case <-ctx.Done():
			return Fail[int, error](ctx.Err())
		}
	})

	go func() {
		// Cancel the context after a moment
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	result := cont.Run(ctx)
	is.True(result.Failure())
	is.Equal(context.Canceled, result.Error())
}
