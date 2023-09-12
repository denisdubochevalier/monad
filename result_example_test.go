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

	m1 := monad.Succeed[int, error](1)
	spew.Printf("m1 := monad.Succeed[int, error](1) -> Value: %#v\n", m1.Value())
	spew.Printf("m1 := monad.Succeed[int, error](1) -> Error: %#v\n", m1.Error())
	spew.Printf("m1 := monad.Succeed[int, error](1) -> failure: %#v\n", m1.Failure())
	spew.Printf("m1 := monad.Succeed[int, error](1) -> success: %#v\n", m1.Success())

	m2 := monad.Fail[int, error](errors.New("test"))
	spew.Printf("m2 := monad.Fail[int, error](errors.New(\"test\")) -> Value: %#v\n", m2.Value())
	spew.Printf("m2 := monad.Fail[int, error](errors.New(\"test\")) -> Error: %#v\n", m2.Error())
	spew.Printf(
		"m2 := monad.Fail[int, error](errors.New(\"test\")) -> failure: %#v\n",
		m2.Failure(),
	)
	spew.Printf(
		"m2 := monad.Fail[int, error](errors.New(\"test\")) -> success: %#v\n",
		m2.Success(),
	)

	m3 := monad.FromTuple[int, error](1, nil)
	spew.Printf("m3 := monad.FromTuple[int, error](1, nil) -> Value: %#v\n", m3.Value())
	spew.Printf("m3 := monad.FromTuple[int, error](1, nil) -> Error: %#v\n", m3.Error())

	m4 := monad.FromTuple(1, errors.New("test"))
	spew.Printf("m4 := monad.FromTuple(1, errors.New(\"test\")) -> Value: %#v\n", m4.Value())
	spew.Printf("m4 := monad.FromTuple(1, errors.New(\"test\")) -> Error: %#v\n", m4.Error())

	spew.Println("\nFlatMap:\n=====")
	m8 := m1.FlatMap(
		func(x int) monad.Result[int, error] { return monad.Succeed[int, error](x * 2) },
	)
	spew.Printf(
		"m8 := m1.FlatMap(func(x int) monad.Result[int, error] { return monad.Succeed[int, error](x * 2) }) -> %#v\n",
		m8,
	)
	m9 := m2.FlatMap(
		func(x int) monad.Result[int, error] { return monad.Succeed[int, error](x * 2) },
	)
	spew.Printf(
		"m9 := m2.FlatMap(func(x int) monad.Result[int, error] { return monad.Succeed[int, error](x * 2) }) -> %#v\n",
		m9,
	)

	spew.Println("\nOr:\n===")
	m10 := m1.Or(func(_ error) monad.Result[int, error] { return monad.Succeed[int, error](1) })
	spew.Printf(
		"m10 := m1.Or(func(_ error) monad.Result[int, error]{return monad.Succeed[int, error](1)}) -> %#v\n",
		m10,
	)
	// Output:
	// Computing example monad.Result values, showing what calling the .Value() and .Error() method on them returns
	//
	// Creation:
	// =========
	// m1 := monad.Succeed[int, error](1) -> Value: (int)1
	// m1 := monad.Succeed[int, error](1) -> Error: (interface {})<nil>
	// m1 := monad.Succeed[int, error](1) -> failure: (bool)false
	// m1 := monad.Succeed[int, error](1) -> success: (bool)true
	// m2 := monad.Fail[int, error](errors.New("test")) -> Value: (int)0
	// m2 := monad.Fail[int, error](errors.New("test")) -> Error: (*errors.errorString)test
	// m2 := monad.Fail[int, error](errors.New("test")) -> failure: (bool)true
	// m2 := monad.Fail[int, error](errors.New("test")) -> success: (bool)false
	// m3 := monad.FromTuple[int, error](1, nil) -> Value: (int)1
	// m3 := monad.FromTuple[int, error](1, nil) -> Error: (interface {})<nil>
	// m4 := monad.FromTuple(1, errors.New("test")) -> Value: (int)0
	// m4 := monad.FromTuple(1, errors.New("test")) -> Error: (*errors.errorString)test
	//
	// FlatMap:
	// =====
	// m8 := m1.FlatMap(func(x int) monad.Result[int, error] { return monad.Succeed[int, error](x * 2) }) -> (monad.success[int,error]){val:(int)2}
	// m9 := m2.FlatMap(func(x int) monad.Result[int, error] { return monad.Succeed[int, error](x * 2) }) -> (monad.failure[int,error]){err:(*errors.errorString)test}
	//
	// Or:
	// ===
	// m10 := m1.Or(func(_ error) monad.Result[int, error]{return monad.Succeed[int, error](1)}) -> (monad.success[int,error]){val:(int)1}
}
