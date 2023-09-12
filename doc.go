// Package monad provides an extensive implementation of various monadic
// structures, encapsulating a wide range of computational design patterns.
// These include: Maybe, Either, Result, Identity, List, Reader, Writer, and
// State.
//
// Originating from category theory and serving as computational building blocks
// in functional programming, monads enable fine-grained control over
// programmatic side-effects, sequencing, and value context.
//
// Every monad implemented in this package rigorously adheres to the three
// foundational monadic laws:
// - Left Identity:  return a >>= f  is equivalent to  f(a)
// - Right Identity: m >>= return  is equivalent to  m
// - Associativity:  (m >>= f) >>= g  is equivalent to  m >>= (\x -> f x >>= g)
//
// The following monads are implemented:
//   - Maybe: Represents optional values and handles the absence of a value.
//   - Either: Extends Maybe by encapsulating an error or alternate reason for
//     the absence of a value.
//   - Result: Tailored to Go's idiomatic error handling, encapsulating either a
//     successful result or an error.
//   - Identity: The simplest monad, acting as a container for a single value.
//   - List: Represents a collection of values in a monadic context.
//   - Reader: Encapsulates a shared environment required by various
//     computations.
//   - Writer: Captures additional output during the computation, useful for
//     logging or state tracing.
//   - State: Allows the embedding of stateful computations within the monadic
//     structure.
//   - Validation: To accumulate all errors rather than failing fast,
//     useful in form validation.
//   - Continuation: Suited for asynchronous or nested computations.
//   - IO Monad: To encapsulate effectful computations in a functional style.
//   - Future: To encapsulate future asynchronous computations.
//   - Free Monad: An advanced construct to build interpreters for embedded
//     DSLs.
//
// Planned monads:
//   - RWS (Reader-Writer-State) Monad: An amalgam of the Reader, Writer, and
//     State monads, this monad is particularly useful when you need to carry an
//     environment, write logs, and maintain state all in one.
//   - Promise Monad: Similar to the Future Monad but offers more flexibility in
//     terms of chaining asynchronous operations and handling failures.
//   - Try Monad: A derivative of the Either Monad that's specialized for
//     capturing exceptions, making it easier to interface with code that might
//     throw exceptions rather than return explicit error values.
//   - Logic Monad: Useful for computations that involve non-determinism or
//     search, akin to logic programming in languages like Prolog.
//   - Event Monad: Aimed at event-driven architectures, this could encapsulate
//     event sources as monadic values, allowing you to manipulate and combine
//     them functionally.
//   - Transaction Monad: Useful for operations that could be rolled back,
//     allowing you to capture the essence of transactions in a functional
//     context.
//   - Parser Combinators as Monads: While not a singular monad, parser
//     combinators often make extensive use of monadic interfaces to create
//     highly compositional parsing strategies.
//   - Probabilistic Monad: For stochastic or probabilistic computations, allows
//     you to build up complex models while maintaining the ability to extract
//     probability distributions.
//   - Co-Routine Monad: Provides a way to encapsulate generators and other
//     co-routines in a functional way.
//   - Lens Monad: Though lenses are more commonly associated with functional
//     programming languages like Haskell, a Lens monad could offer a similar
//     capability of focusing on and manipulating nested immutable data
//     structures in a clean way.
//   - Process Monad: For encapsulating potentially long-running, stateful
//     processes that might be paused, resumed, or terminated.
//
// Each monad is endowed with functional methods like `FlatMap` and `Map` to
// facilitate composability and side-effect management.
//
// Usage:
// Consult the associated documentation for each individual monad to explore
// example usage and further details.
package monad
