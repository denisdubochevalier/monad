# Contributing to monad

Thank you for considering contributing to monad. Your contributions, in the form
of issues, pull requests, and insights, are highly valued.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Forking the Repository](#forking-the-repository)
- [Development Environment](#development-environment)
- [Submitting a Pull Request](#submitting-a-pull-request)
- [Testing](#testing)
- [Documentation](#documentation)
- [Issue Tracking](#issue-tracking)

## Code of Conduct

We adhere to a code of conduct adapted from the
[Contributor Covenant](https://www.contributor-covenant.org/version/2/0/code_of_conduct/),
and we expect all contributors to respect it. Harassment will not be tolerated.

## Forking the Repository

If you're not a direct contributor, you'll need to fork the repository. This is
essential for you to make your contributions. Simply click on the "Fork" button
at the top right of this repository.

## Development Environment

Before diving into the intricacies of monad, ensure your development environment
is adequately configured:

- Go version 1.20 or higher
- A properly configured `$GOPATH`

## Submitting a Pull Request

1. Ensure you are working on a forked repository and are up to date with the
   latest code from the `main` branch.
2. Create a new branch descriptive of the feature or issue
   (`git checkout -b feature/my-new-feature`).
3. Implement your changes, adhering to the Go programming idioms and the
   established code style of the repository.
4. Run tests to ensure your code changes are both functional and well-tested.
5. Commit your changes, following a
   [conventional commit message](https://www.conventionalcommits.org/en/v1.0.0/).
6. Push your changes to your fork (`git push origin feature/my-new-feature`).
7. Open a pull request against the `main` branch.
8. Describe the changes in detail and link any relevant issues.

Your pull request will undergo code reviews, and you may be asked to make
changes. Once your pull request is approved, it will be merged into `main`.

## Testing

We believe in test-driven development and request that you write tests to cover
all new functionality, as well as any changes to existing functionality. These
should adhere to the existing test patterns within the repository.

## Documentation

Accurate documentation is indispensable. Please update the documentation to
reflect your changes or new features, ensuring your documentation follows the
current layout.

## Issue Tracking

We use GitHub issues to track public bugs and features. Please ensure your bug
description is clear and has sufficient instructions to be able to reproduce the
issue.
