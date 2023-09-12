package monad

// Functor represents a higher-kinded type that encapsulates an endofunctor
// in the category of Go's types. In simpler terms, a Functor is a type-level
// function that takes a type and produces a new type, preserving both the
// structure and the behavior of operations (methods) defined on the original
// type.
//
// A Functor must satisfy two fundamental laws:
//
// 1. Identity:   fmap id  ==  id
// 2. Composition: fmap (f . g)  ==  fmap f . fmap g
//
// The Map method provides a mechanism to transform the encapsulated value(s)
// inside the Functor without altering its structural context. This allows
// for operations on the encapsulated type to be 'lifted' to operate on the
// Functor itself.
//
// Note: The Extract method is an extension that enables the retrieval of the
// encapsulated value. It is not part of the traditional Functor definition
// but serves a specific purpose within the Free Monad context in Go.
type Functor[F any] interface {
	Map(func(any) any) Functor[F]
	Extract() F
}
