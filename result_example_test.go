package monad_test

import (
	"errors"

	"github.com/davecgh/go-spew/spew"

	"github.com/denisdubochevalier/monad"
)

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

	spew.Println("\nFMap:\n=====")
	m8 := m1.FMap(func(x int) int { return x * 2 })
	spew.Printf("m8 := m1.FMap(func(x int) int { return x * 2 }) -> %#v\n", m8)
	m9 := m2.FMap(func(x int) int { return x * 2 })
	spew.Printf("m9 := m2.FMap(func(x int) int { return x * 2 }) -> %#v\n", m9)

	spew.Println("\nOr:\n===")
	m10 := m1.Or(func(_ error) {})
	spew.Printf("m10 := m1.Or(func(_ error) {}) -> %#v\n", m10)
	m11 := m2.Or(func(_ error) {})
	spew.Printf("m11 := m2.Or(func(_ error) {}) -> %#v\n", m11)

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
	//
	// FMap:
	// =====
	// m8 := m1.FMap(func(x int) int { return x * 2 }) -> (monad.Success[int]){val:(int)2}
	// m9 := m2.FMap(func(x int) int { return x * 2 }) -> (monad.Failure[int]){err:(*errors.errorString)test}
	//
	// Or:
	// ===
	// m10 := m1.Or(func(_ error) {}) -> (monad.Success[int]){val:(int)1}
	// m11 := m2.Or(func(_ error) {}) -> (monad.Failure[int]){err:(*errors.errorString)test}
}
