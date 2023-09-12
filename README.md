# monad

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.20-%23007d9c)
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

## Planned Additions

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

## Installation

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
