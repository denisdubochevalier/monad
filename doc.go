// Package monad implements three commonly used monads: Maybe, Either, and
// Result.
//
// Monads are a category-theoretic concept that have found extensive
// applicability in functional programming. They serve as "programmable
// semicolons"; they dictate the order and side effects of function calls. While
// Go is not purely functional, the design patterns provided by monads can be
// useful for managing side-effectsand control flow.
//
// The monads defined in this package adhere strictly to the three monadic laws:
// - Left Identity:  return a >>= f  is equivalent to  f(a)
// - Right Identity: m >>= return  is equivalent to  m
// - Associativity:  (m >>= f) >>= g  is equivalent to  m >>= (\x -> f x >>= g)
//
// Maybe Monad:
// The Maybe monad is used for computations that may fail to return a value.
// The Maybe monad encapsulates an optional value, and allows you to compose
// functions that might not return a meaningful result.
//
// Either Monad:
// The Either monad extends Maybe to capture the reason for failure, often
// utilized for error handling. Either has two subclasses, Left and Right; by
// convention, Right is the "happy path" and Left captures an error or alternate
// value.
//
// Result Monad:
// Result is tailored to Go's idiomatic error handling. It represents the
// outcome of a function that could fail, wrapping either a successful result or
// an error. Unlike Maybe and Either, Result explicitly deals with Go's native
// error type.
//
// Each monad has a set of functional methods, primarily `FlatMap` and `Or`, to
// facilitate function composition, control flow, and side-effect management.
//
// Usage:
// Refer to individual monad sub-packages for example usage and further details.
package monad
