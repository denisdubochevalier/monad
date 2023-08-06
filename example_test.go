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
	m1 := monad.OfValue(5)
	spew.Printf("m1 := monad.OfValue(5) -> %#v\n", m1.Value())

	m2 := monad.OfNullable[int](new(int))
	spew.Printf("m2 := monad.OfNullable[int](new(int)) -> %#v\n", m2.Value())

	m3 := monad.Empty[any]()
	spew.Printf("m3 := monad.Empty[any]() -> %#v\n", m3.Value())

	spew.Println("\nMap:\n====")
	m4 := monad.Map(monad.OfValue[int](5), strconv.Itoa)
	spew.Printf(
		"m4 := monad.Map(monad.OfValue(5), strconv.Itoa) -> %#v\n",
		m4.Value(),
	)

	m5 := monad.Map(divideBy(6, 3), func(x int) int { return x * 2 })
	spew.Printf(
		"m5 := monad.Map(divideBy(6, 3), func(x int) int { return x * 2 }) -> %#v\n",
		m5.Value(),
	)

	m6 := monad.Map(divideBy(6, 0), func(x int) int { return x * 2 })
	spew.Printf(
		"m6 := monad.Map(divideBy(6, 0), func(x int) int { return x * 2 }) -> %#v\n",
		m6.Value(),
	)

	m7 := monad.FlatMap[int, int](
		divideBy(6, 0),
		func(x int) monad.Maybe[int] { return divideBy(0, 7) },
	)
	spew.Printf(
		"m7 := monad.FlatMap[int, int](divideBy(6, 0), func(x int) monad.Maybe[int] { return divideBy(0, 7) }) -> %#v\n",
		m7.Value(),
	)

	m8 := monad.FlatMap(divideBy(6, 5), func(x int) monad.Maybe[int] { return divideBy(5, 3) })
	spew.Printf(
		"m8 := monad.FlatMap(divideBy(6, 5), func(x int) monad.Maybe[int] { return divideBy(5, 3) }) -> %#v\n",
		m8.Value(),
	)

	spew.Println("\nFilter:\n=======")
	m9 := monad.OfValue(5).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m9 := monad.OfValue(5).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#v\n",
		m9.Value(),
	)

	m10 := monad.OfValue(6).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m10 := monad.OfValue(6).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#v\n",
		m10.Value(),
	)

	m11 := monad.Empty[int]().Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m11 := monad.Empty[int].Filter(func(x int) bool { return x %% 2 == 0 }) -> %#v\n",
		m11.Value(),
	)

	m12 := monad.OfNullable[int](nil).Filter(func(x int) bool { return x%2 == 0 })
	spew.Printf(
		"m12 := monad.OfNullable[int](nil).Filter(func(x int) bool { return x %% 2 == 0 }) -> %#v\n",
		m12.Value(),
	)

	spew.Println("\nOrElse:\n=======")
	v1 := monad.OfValue(1).OrElse(2)
	spew.Printf("v1 := monad.OfValue(1).OrElse(2) -> %#v\n", v1)
	v2 := monad.Empty[int]().OrElse(2)
	spew.Printf("v2 := monad.Empty[int]().OrElse(2) -> %#v\n", v2)

	// Output:
	// Computing example monad.Maybe values, showing what calling the .Value() method on them returns:
	//
	// Creation:
	// =========
	// m1 := monad.OfValue(5) -> (int)5
	// m2 := monad.OfNullable[int](new(int)) -> (int)0
	// m3 := monad.Empty[any]() -> (interface {})<nil>
	//
	// Map:
	// ====
	// m4 := monad.Map(monad.OfValue(5), strconv.Itoa) -> (string)5
	// m5 := monad.Map(divideBy(6, 3), func(x int) int { return x * 2 }) -> (int)4
	// m6 := monad.Map(divideBy(6, 0), func(x int) int { return x * 2 }) -> (int)0
	// m7 := monad.FlatMap[int, int](divideBy(6, 0), func(x int) monad.Maybe[int] { return divideBy(0, 7) }) -> (int)0
	// m8 := monad.FlatMap(divideBy(6, 5), func(x int) monad.Maybe[int] { return divideBy(5, 3) }) -> (int)1
	//
	// Filter:
	// =======
	// m9 := monad.OfValue(5).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
	// m10 := monad.OfValue(6).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)6
	// m11 := monad.Empty[int].Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
	// m12 := monad.OfNullable[int](nil).Filter(func(x int) bool { return x % 2 == 0 }) -> (int)0
	//
	// OrElse:
	// =======
	// v1 := monad.OfValue(1).OrElse(2) -> (int)1
	// v2 := monad.Empty[int]().OrElse(2) -> (int)2
}

func divideBy(x, y int) monad.Maybe[int] {
	if y == 0 {
		return monad.Empty[int]()
	}
	return monad.OfValue[int](x / y)
}

func ExampleResult() {
	spew.Println(
		"Computing example monad.Result values, showing what calling the .Value() and .Error() method on them returns",
	)
	spew.Println("\nCreation:")
	spew.Println("=========")

	m1 := monad.Succeed(1)
	spew.Printf("m1 := monad.Succeed(1) -> Value: %#v\n", m1.Value())
	spew.Printf("m1 := monad.Succeed(1) -> Error: %#v\n", m1.Error())
	spew.Printf("m1 := monad.Succeed(1) -> Failure: %#v\n", m1.Failure())
	spew.Printf("m1 := monad.Succeed(1) -> Success: %#v\n", m1.Success())

	m2 := monad.Fail[int](errors.New("test"))
	spew.Printf("m2 := monad.Fail[int](errors.New(\"test\")) -> Value: %#v\n", m2.Value())
	spew.Printf("m2 := monad.Fail[int](errors.New(\"test\")) -> Error: %#v\n", m2.Error())
	spew.Printf("m2 := monad.Fail[int](errors.New(\"test\")) -> Failure: %#v\n", m2.Failure())
	spew.Printf("m2 := monad.Fail[int](errors.New(\"test\")) -> Success: %#v\n", m2.Success())

	m3 := monad.FromTuple(1, nil)
	spew.Printf("m3 := monad.FromTuple(1, nil) -> Value: %#v\n", m3.Value())
	spew.Printf("m3 := monad.FromTuple(1, nil) -> Error: %#v\n", m3.Error())

	m4 := monad.FromTuple(1, errors.New("test"))
	spew.Printf("m4 := monad.FromTuple(1, errors.New(\"test\")) -> Value: %#v\n", m4.Value())
	spew.Printf("m4 := monad.FromTuple(1, errors.New(\"test\")) -> Error: %#v\n", m4.Error())

	spew.Println("\nBind:")
	spew.Println("=====")

	m5 := m1.Bind(func(x int) (int, error) { return x * 2, nil })
	spew.Printf(
		"m5 := m1.Bind(func(x int) (int, error) { return x * 2, nil }) -> Value: %#v\n",
		m5.Value(),
	)
	spew.Printf(
		"m5 := m1.Bind(func(x int) (int, error) { return x * 2, nil }) -> Error: %#v\n",
		m5.Error(),
	)

	m6 := m5.Bind(func(_ int) (int, error) { return 0, errors.New("test") })
	spew.Printf(
		"m6 := m5.Bind(func(_ int) (int, error) { return 0, errors.New(\"test\") }) -> Value: %#v\n",
		m6.Value(),
	)
	spew.Printf(
		"m6 := m5.Bind(func(_ int) (int, error) { return 0, errors.New(\"test\") }) -> Error: %#v\n",
		m6.Error(),
	)

	m7 := m2.Bind(func(x int) (int, error) { return x * 2, nil })
	spew.Printf(
		"m7 := m2.Bind(func(x int) (int, error) { return x * 2, nil }) -> Value: %#v\n",
		m7.Value(),
	)
	spew.Printf(
		"m7 := m2.Bind(func(x int) (int, error) { return x * 2, nil }) -> Error: %#v\n",
		m7.Error(),
	)

	// Output:
	// Computing example monad.Result values, showing what calling the .Value() and .Error() method on them returns
	//
	// Creation:
	// =========
	// m1 := monad.Succeed(1) -> Value: (int)1
	// m1 := monad.Succeed(1) -> Error: (interface {})<nil>
	// m1 := monad.Succeed(1) -> Failure: (bool)false
	// m1 := monad.Succeed(1) -> Success: (bool)true
	// m2 := monad.Fail[int](errors.New("test")) -> Value: (int)0
	// m2 := monad.Fail[int](errors.New("test")) -> Error: (*errors.errorString)test
	// m2 := monad.Fail[int](errors.New("test")) -> Failure: (bool)true
	// m2 := monad.Fail[int](errors.New("test")) -> Success: (bool)false
	// m3 := monad.FromTuple(1, nil) -> Value: (int)1
	// m3 := monad.FromTuple(1, nil) -> Error: (interface {})<nil>
	// m4 := monad.FromTuple(1, errors.New("test")) -> Value: (int)0
	// m4 := monad.FromTuple(1, errors.New("test")) -> Error: (*errors.errorString)test
	//
	// Bind:
	// =====
	// m5 := m1.Bind(func(x int) (int, error) { return x * 2, nil }) -> Value: (int)2
	// m5 := m1.Bind(func(x int) (int, error) { return x * 2, nil }) -> Error: (interface {})<nil>
	// m6 := m5.Bind(func(_ int) (int, error) { return 0, errors.New("test") }) -> Value: (int)0
	// m6 := m5.Bind(func(_ int) (int, error) { return 0, errors.New("test") }) -> Error: (*errors.errorString)test
	// m7 := m2.Bind(func(x int) (int, error) { return x * 2, nil }) -> Value: (int)0
	// m7 := m2.Bind(func(x int) (int, error) { return x * 2, nil }) -> Error: (*errors.errorString)test
}
