package monad_test

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/denisdubochevalier/monad"
)

func ExampleMaybe() {
	spew.Println(
		"Computing example monad.Maybe values, showing what calling the .Value() method on them returns:\n",
	)
	spew.Println("Creation:\n=========")
	m1 := monad.Some(5)
	spew.Printf("m1 := monad.Some(5) -> %#v\n", m1.Value())

	m2 := monad.Nullable[int](nil)
	spew.Printf("m2 := monad.Nullable[int](new(int)) -> %#v\n", m2.Value())

	m3 := monad.None[any]()
	spew.Printf("m3 := monad.None[any]() -> %#v\n", m3.Value())

	spew.Println("\nFMap:\n=====")
	m4 := m1.FMap(func(x int) int { return x * 2 })
	spew.Printf("m4 := m1.FMap(func(x int) int { return x * 2 }) -> %#v\n", m4.Value())
	m5 := m2.FMap(func(x int) int { return x * 2 })
	spew.Printf("m5 := m2.FMap(func(x int) int { return x * 2 }) -> %#v\n", m5.Value())

	spew.Println("\nBind:\n=====")
	m6 := m1.Bind(func(x int) *int { x = x * 2; return &x })
	spew.Printf("m6 := m1.Bind(func(x int) *int { x = x * 2; return &x }) -> %#v\n", m6.Value())
	m7 := m1.Bind(func(x int) *int { return nil })
	spew.Printf("m7 := m1.Bind(func(x int) *int { return nil }) -> %#v\n", m7.Value())
	m8 := m2.Bind(func(x int) *int { x = x * 2; return &x })
	spew.Printf("m8 := m2.Bind(func(x int) *int { x = x * 2; return &x }) -> %#v\n", m8.Value())

	spew.Println("\nFilter:\n=======")
	m9 := monad.Some(5).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m9 := monad.Some(5).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#v\n",
		m9.Value(),
	)

	m10 := monad.Some(6).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m10 := monad.Some(6).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#v\n",
		m10.Value(),
	)

	m11 := monad.None[int]().Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m11 := monad.None[int].Filter(func(x int) bool { return x %% 2 == 0 }) -> %#v\n",
		m11.Value(),
	)

	m12 := monad.Nullable[int](nil).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m12 := monad.Nullable[int](nil).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#v\n",
		m12.Value(),
	)

	spew.Println("\nOrElse:\n=======")
	v1 := monad.Some(1).OrElse(2)
	spew.Printf("v1 := monad.Some(1).OrElse(2) -> %#v\n", v1)
	v2 := monad.None[int]().OrElse(2)
	spew.Printf("v2 := monad.None[int]().OrElse(2) -> %#v\n", v2)

	spew.Println("\nFBind:\n======")
	m13 := m1.FBind(func(x int) monad.Maybe[int] { return monad.Some(x * 2) })
	spew.Printf(
		"m13 := m1.FBind(func(x int) monad.Maybe[int] { return monad.Some(x * 2) }) -> %#v\n",
		m13,
	)
	m14 := monad.None[int]().FBind(func(x int) monad.Maybe[int] { return monad.Some(x * 2) })
	spew.Printf(
		"m14 := monad.None[int]().FBind(func(x int) monad.Maybe[int] { return monad.Some(x * 2) }) -> %#v\n",
		m14,
	)

	// Output:
	// Computing example monad.Maybe values, showing what calling the .Value() method on them returns:
	//
	// Creation:
	// =========
	// m1 := monad.Some(5) -> (int)5
	// m2 := monad.Nullable[int](new(int)) -> (int)0
	// m3 := monad.None[any]() -> (interface {})<nil>
	//
	// FMap:
	// =====
	// m4 := m1.FMap(func(x int) int { return x * 2 }) -> (int)10
	// m5 := m2.FMap(func(x int) int { return x * 2 }) -> (int)0
	//
	// Bind:
	// =====
	// m6 := m1.Bind(func(x int) *int { x = x * 2; return &x }) -> (int)10
	// m7 := m1.Bind(func(x int) *int { return nil }) -> (int)0
	// m8 := m2.Bind(func(x int) *int { x = x * 2; return &x }) -> (int)0
	//
	// Filter:
	// =======
	// m9 := monad.Some(5).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
	// m10 := monad.Some(6).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)6
	// m11 := monad.None[int].Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
	// m12 := monad.Nullable[int](nil).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
	//
	// OrElse:
	// =======
	// v1 := monad.Some(1).OrElse(2) -> (int)1
	// v2 := monad.None[int]().OrElse(2) -> (int)2
	//
	// FBind:
	// ======
	// m13 := m1.FBind(func(x int) monad.Maybe[int] { return monad.Some(x * 2) }) -> (monad.Just[int]){val:(int)10}
	// m14 := monad.None[int]().FBind(func(x int) monad.Maybe[int] { return monad.Some(x * 2) }) -> (monad.Nothing[int]){}
}
