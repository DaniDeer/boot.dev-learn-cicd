# learn-cicd-starter (Notely)

![CI status](https://github.com/DaniDeer/boot.dev-learn-cicd/actions/workflows/ci.yml/badge.svg)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

_This starts the server in non-database mode._ It will serve a simple webpage at `http://localhost:8080`.

You do _not_ need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!

# Daniel's version of Boot.dev's Notely app

## Learnings

- [Course: Boot.dev - Breakdown of GitHub Actions](https://www.boot.dev/lessons/de167f69-46ed-474a-9a10-c9749b8040b6)
- [Course: Boot.dev - GitHub Actions](https://www.boot.dev/lessons/6df20cf9-d8b3-49be-a299-5ecaaa40b71d)

### Others

Testing if the code is properly formatted:

```bash
test -z $(go fmt ./...)
echo $?
```

`echo $?` will print the exit code of the previous command. If the code is properly formatted, `go fmt` will not make any changes and will exit with a status code of `0`, which means success. If there are formatting issues, `go fmt` will make changes to the code and exit with a non-zero status code (`1`), indicating that there were formatting issues that need to be addressed.

### Testing in GO

#### Unit Testing in Go

- [Blog: Dave Cheney - Prefer Table Driven Tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

In GO tests are just regular GO functions with a few rules:

1. The name of the test function must start with `Test`.
2. The test function must take one argument of type `*testing.T`. A `*testing.T` is a type injected by the `testing` package itself, to provide ways to print, skip, and fail the test.

**Table-Driven Tests**

- Use table-driven tests to test multiple cases in a single test function and reducing boilerplate code.
- Use the `t.Run` method to run subtests for each case in the table-driven test, which allows for better organization and reporting of test results.
- Use the `cmp.Diff` function from the `github.com/google/go-cmp/cmp` package to compare expected and actual values in tests, which provides a more detailed and human-readable output when tests fail.

Example see here: `./internal/auth/get_api_key_test.go`

**Code Coverage** is a measure of how much of your code is executed when you run your tests.

Show coverage of tests:

```GO
go test ./... -cover
```

#### Integration Testing in Go

- [Blog: Boot.dev - Writing Good Unit Tests: Don't Mock Database Connections](https://www.boot.dev/blog/backend/writing-good-unit-tests-dont-mock-database-connections/)
- [Blog: CircleCI - Unit Testing vs Integration Testing](https://circleci.com/blog/unit-testing-vs-integration-testing/)

### Formatting and Linting in GO

#### Formatting

Enforces whitespace, indentation, and line length rules.

```BASH
go fmt ./...
```

#### Linting

Detects potential errors, code smells, and style issues in your code.

The most popular linter for GO is [Staticccheck](https://staticcheck.dev/docs/) ([github.com/dominikh/go-tools](https://github.com/dominikh/go-tools)). It comes with sane defaults and it is easy to configure.

```BASH
# Install Staticcheck
go install honnef.co/go/tools/cmd/staticcheck@latest

# Run Staticcheck
staticcheck ./...
```

Find the `staticcheck` action in the [GitHub Marketplace](https://github.com/marketplace/actions/staticcheck)

See `staticcheck` checks here: [Staticcheck Checks](https://staticcheck.dev/docs/checks)
