package monad_test

import (
	"errors"
	"strconv"

	"github.com/davecgh/go-spew/spew"

	"github.com/denisdubochevalier/monad"
)

func ExampleMaybe() {
	spew.Println(
		"Computing example monad.Maybe values, showing what calling the .Value() method on them returns:\n",
	)
	spew.Println("Creation:\n=========")
	m1 := monad.New(5)
	spew.Printf("m1 := monad.New(5) -> %#+v\n", m1.Value())

	m2 := monad.OfNullable[int](new(int))
	spew.Printf("m2 := monad.OfNullable[int](new(int)) -> %#+v\n", m2.Value())

	m3 := monad.Empty[any]()
	spew.Printf("m3 := monad.Empty[any]() -> %#+v\n", m3.Value())

	spew.Println("\nMap:\n====")
	m4 := monad.Map(monad.New[int](5), strconv.Itoa)
	spew.Printf(
		"m4 := monad.Map(monad.New(5), strconv.Itoa) -> %#+v\n",
		m4.Value(),
	)

	m5 := monad.Map(divideBy(6, 3), func(x int) int { return x * 2 })
	spew.Printf(
		"m5 := monad.Map(divideBy(6, 3), func(x int) int { return x * 2 }) -> %#+v\n",
		m5.Value(),
	)

	m6 := monad.Map(divideBy(6, 0), func(x int) int { return x * 2 })
	spew.Printf(
		"m6 := monad.Map(divideBy(6, 0), func(x int) int { return x * 2 }) -> %#+v\n",
		m6.Value(),
	)

	m7 := monad.FlatMap[int, int](
		divideBy(6, 0),
		func(x int) monad.Maybe[int] { return divideBy(0, 7) },
	)
	spew.Printf(
		"m7 := monad.Flatmap[int, int](divideBy(6, 0), func(x int) monad.Maybe[int] { return divideBy(0, 7) }) -> %#+v\n",
		m7.Value(),
	)

	spew.Println("\nFilter:\n=======")
	m8 := monad.New(5).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m8 := monad.New(5).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#+v\n",
		m8.Value(),
	)

	m9 := monad.New(6).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m9 := monad.New(6).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#+v\n",
		m9.Value(),
	)

	m10 := monad.Empty[int]().Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m10 := monad.Empty[int].Filter(func(x int) bool { return x %% 2 == 0 }) -> %#+v\n",
		m10.Value(),
	)

	m11 := monad.OfNullable[int](nil).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m11 := monad.OfNullable[int](nil).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#+v\n",
		m11.Value(),
	)

	// Output:
	// Computing example monad.Maybe values, showing what calling the .Value() method on them returns:
	//
	// Creation:
	// =========
	// m1 := monad.New(5) -> (int)5
	// m2 := monad.OfNullable[int](new(int)) -> (int)0
	// m3 := monad.Empty[any]() -> (interface {})<nil>
	//
	// Map:
	// ====
	// m4 := monad.Map(monad.New(5), strconv.Itoa) -> (string)5
	// m5 := monad.Map(divideBy(6, 3), func(x int) int { return x * 2 }) -> (int)4
	// m6 := monad.Map(divideBy(6, 0), func(x int) int { return x * 2 }) -> (int)0
	// m7 := monad.Flatmap[int, int](divideBy(6, 0), func(x int) monad.Maybe[int] { return divideBy(0, 7) }) -> (int)0
	//
	// Filter:
	// =======
	// m8 := monad.New(5).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
	// m9 := monad.New(6).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)6
	// m10 := monad.Empty[int].Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
	// m11 := monad.OfNullable[int](nil).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
}

func divideBy(x, y int) monad.Maybe[int] {
	if y == 0 {
		return monad.Empty[int]()
	}
	return monad.New[int](x / y)
}

func ExampleErrorHandler() {
	spew.Println(
		"Computing example monad.ErrorHandler values, showing what calling the .Value() and .Error() method on them returns",
	)
	spew.Println("\nCreation:")
	spew.Println("=========")

	m1 := monad.Succeed(1)
	spew.Printf("m1 := monad.Succeed(1) -> Value: %#v\n", m1.Value())
	spew.Printf("m1 := monad.Succeed(1) -> Error: %#v\n", m1.Error())

	m2 := monad.Fail[int](errors.New("test"))
	spew.Printf("m2 := monad.Fail[int](errors.New(\"test\")) -> Value: %#v\n", m2.Value())
	spew.Printf("m2 := monad.Fail[int](errors.New(\"test\")) -> Error: %#v\n", m2.Error())

	m3 := monad.FromTuple(1, nil)
	spew.Printf("m3 := monad.FromTuple(1, nil) -> Value: %#v\n", m3.Value())
	spew.Printf("m3 := monad.FromTuple(1, nil) -> Error: %#v\n", m3.Error())

	spew.Println("\nBind:")
	spew.Println("=====")

	m4 := m1.Bind(func(x int) (int, error) { return x * 2, nil })
	spew.Printf(
		"m4 := m1.Bind(func(x int) (int, error) { return x * 2, nil }) -> Value: %#v\n",
		m4.Value(),
	)
	spew.Printf(
		"m4 := m1.Bind(func(x int) (int, error) { return x * 2, nil }) -> Error: %#v\n",
		m4.Error(),
	)

	m5 := m4.Bind(func(_ int) (int, error) { return 0, errors.New("test") })
	spew.Printf(
		"m5 := m4.Bind(func(_ int) (int, error) { return 0, errors.New(\"test\") }) -> Value: %#v\n",
		m5.Value(),
	)
	spew.Printf(
		"m5 := m4.Bind(func(_ int) (int, error) { return 0, errors.New(\"test\") }) -> Error: %#v\n",
		m5.Error(),
	)

	m6 := m2.Bind(func(x int) (int, error) { return x * 2, nil })
	spew.Printf(
		"m6 := m2.Bind(func(x int) (int, error) { return x * 2, nil }) -> Value: %#v\n",
		m6.Value(),
	)
	spew.Printf(
		"m6 := m2.Bind(func(x int) (int, error) { return x * 2, nil }) -> Error: %#v\n",
		m6.Error(),
	)

	// Output:
	// Computing example monad.ErrorHandler values, showing what calling the .Value() and .Error() method on them returns
	//
	// Creation:
	// =========
	// m1 := monad.Succeed(1) -> Value: (int)1
	// m1 := monad.Succeed(1) -> Error: (interface {})<nil>
	// m2 := monad.Fail[int](errors.New("test")) -> Value: (int)0
	// m2 := monad.Fail[int](errors.New("test")) -> Error: (*errors.errorString)test
	// m3 := monad.FromTuple(1, nil) -> Value: (int)1
	// m3 := monad.FromTuple(1, nil) -> Error: (interface {})<nil>
	//
	// Bind:
	// =====
	// m4 := m1.Bind(func(x int) (int, error) { return x * 2, nil }) -> Value: (int)2
	// m4 := m1.Bind(func(x int) (int, error) { return x * 2, nil }) -> Error: (interface {})<nil>
	// m5 := m4.Bind(func(_ int) (int, error) { return 0, errors.New("test") }) -> Value: (int)0
	// m5 := m4.Bind(func(_ int) (int, error) { return 0, errors.New("test") }) -> Error: (*errors.errorString)test
	// m6 := m2.Bind(func(x int) (int, error) { return x * 2, nil }) -> Value: (int)0
	// m6 := m2.Bind(func(x int) (int, error) { return x * 2, nil }) -> Error: (*errors.errorString)test
}
