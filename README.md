# monad

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/denisdubochevalier/monad)
[![GoDoc](https://godoc.org/github.com/denisdubochevalier/monad?status.svg)](https://pkg.go.dev/github.com/denisdubochevalier/monad)
![Build Status](https://github.com/denisdubochevalier/monad/actions/workflows/go.yml/badge.svg)
![Lint Status](https://github.com/denisdubochevalier/monad/actions/workflows/golangci-lint.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/denisdubochevalier/monad)](https://goreportcard.com/report/github.com/denisdubochevalier/monad)
[![Coverage](https://img.shields.io/codecov/c/github/denisdubochevalier/monad)](https://codecov.io/gh/denisdubochevalier/monad)
[![License](https://img.shields.io/github/license/denisdubochevalier/monad)](./LICENSE)

## Introduction

An expansive Go library that encapsulates a plethora of monadic concepts,
rigorously adhering to the three foundational monadic laws. Designed for both
pedagogical exploration and industrial-strength functional programming, this
package serves as a bedrock for complex, side-effect-free computations.

> Note: While the package has matured substantially, it remains a dynamic
> project. Use cautiously in mission-critical applications.

## Features

- Comprehensive suite of monads including but not limited to `Maybe`, `Either`,
  `Result`, `Identity`, `List`, `Reader`, `Writer`, `State`, and more.
- Functionally pure methods like `FlatMap`, `Map`, and other combinators for
  side-effect management.
- Developed with idiomatic Go patterns, optimizing for both readability and
  performance.
- Extensive test coverage to ensure adherence to monadic laws.

## What's Next

The current version of this monadic library serves as a foundational layer upon
which more advanced functionalities can be built. The roadmap ahead is exciting
and aims to elevate this project from a basic utility to a comprehensive toolkit
for functional programming in Go.

## Proper Examples for Each Monad

- **In-Situ Examples**: Each monad will be accompanied by in-situ examples to
  elucidate its practical applications.

- `example_test.go`: The examples will be implemented as `_test.go` files,
  serving dual roles as instructive code snippets and as integration tests.

## New Monads

- **Additional Monads**: Expanding the library's repertoire to include other
  essential monads like the following:
  - [ ] RWS (Reader-Writer-State) Monad
  - [ ] Promise Monad
  - [ ] Try Monad
  - [ ] Logic Monad
  - [ ] Event Monad
  - [ ] Transaction Monad
  - [ ] Parser Combinators as Monads
  - [ ] Probabilistic Monad
  - [ ] Co-Routine Monad
  - [ ] Lens Monad
  - [ ] Process Monad
- **Community Contributions**: We are open to contributions for implementing
  monads that are currently not part of the library but would offer significant
  value.

## Monad Composition and Monad Transformers

- **Composition**: Introducing methods to compose multiple monads into new, more
  powerful constructs.
- **Transformers**: Implementing monad transformers that enable more complex
  operations by stacking multiple monads.
- **Real-world Scenarios**: Accompanying the above with examples and
  documentation that demonstrate these advanced concepts in action, showing how
  they solve real-world problems.

## Installation

Due to the heavy use of on-the-edge generics, `monad` requires go version >
1.21.0

```bash
go get github.com/denisdubochevalier/monad
```

## Usage

Each monad comes with detailed documentation and example code. Please refer to
the GoDoc for individual guides.

```go
import "github.com/denisdubochevalier/monad"

// Example usage with Maybe monad
maybe := monad.Just(42)
result := maybe.FlatMap(func(x int) monad.Maybe {
  return monad.Just(x * 2)
})
```

## Contributing

Contributions are warmly welcomed. Please refer to the
[CONTRIBUTING.md](/CONTRIBUTING.md) file for guidelines.

## License

This project is licensed under the terms of the MIT license. See
[LICENSE](/LICENSE) for more details.
