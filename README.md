# tserr

[![Go Report Card](https://goreportcard.com/badge/github.com/thorsphere/tserr)](https://goreportcard.com/report/github.com/thorsphere/tserr)
[![CodeFactor](https://www.codefactor.io/repository/github/thorsphere/tserr/badge)](https://www.codefactor.io/repository/github/thorsphere/tserr)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorsphere/tserr)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorsphere/tserr)](https://pkg.go.dev/mod/github.com/thorsphere/tserr)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorsphere/tserr)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorsphere/tserr)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorsphere/tserr)
![GitHub last commit](https://img.shields.io/github/last-commit/thorsphere/tserr)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorsphere/tserr)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorsphere/tserr)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorsphere/tserr)
![GitHub](https://img.shields.io/github/license/thorsphere/tserr)

tserr is a lightweight Go package for generating standardized, structured error messages in JSON format. It provides a simple, consistent approach to error handling without any external dependencies.

## Key Features

- **Structured Output**: All errors are formatted as JSON for easy parsing and logging
- **Zero Dependencies**: Only uses the [Go Standard Library](https://pkg.go.dev/std)
- **Simple API**: Just function calls, no configuration needed
- **Tested**: High code coverage with comprehensive unit tests
- **HTTP-Aligned**: Error codes correspond to HTTP status codes for consistency

## Installation

Add tserr to your Go project:

```bash
go get github.com/thorsphere/tserr
```

## Usage

Import the package in your Go code:

```go
import "github.com/thorsphere/tserr"
```

### Two Patterns for Error Functions

tserr supports different calling patterns:

#### Pattern 1: Simple Errors (No Arguments)

For straightforward errors without parameters:

```go
err := tserr.NilPtr()
```

#### Pattern 2: Single-Argument Errors

For errors with one parameter:

```go
f := "foo.txt"
err := tserr.NotExistent(f)
```

#### Pattern 3: Multi-Argument Errors

For errors requiring multiple parameters, pass a pointer to a struct:

```go
err := tserr.EqualStr(&tserr.EqualStrArgs{
    Var:    "username",
    Actual: "alice",
    Want:   "bob",
})
```

**Important**: All multi-argument error functions check if the struct pointer is nil before processing. If nil, they return `tserr.NilPtr()`.

### Output Format

Every error is formatted as a JSON string with consistent structure:

```json
{"error":{"id":8,"code":500,"message":"value of username is alice, but expected to be equal to bob"}}
```

## JSON Format Details

Each error message contains three components:

- **`id`**: A unique, incrementally-numbered error identifier (e.g., 0 for nil pointer, 2 for not existent)
- **`code`**: An HTTP status code corresponding to the error category (e.g., 404 for not found, 400 for bad request)
- **`message`**: The error message (may contain formatted values from function arguments)

Structure:

```json
{
  "error": {
    "id": <int>,
    "code": <int>,
    "message": "<string>"
  }
}
```

## Example

```go
package main

import (
	"fmt"

	"github.com/thorsphere/tserr"
)

func main() {
	// Simple error
	err1 := tserr.NilPtr()
	fmt.Println(err1)

	// Single-argument error
	filename := "config.json"
	err2 := tserr.NotExistent(filename)
	fmt.Println(err2)

	// Multi-argument error
	err3 := tserr.EqualStr(&tserr.EqualStrArgs{
		Var:    "port",
		Actual: "8000",
		Want:   "3000",
	})
	fmt.Println(err3)
}
```

[Run in Go Playground](https://go.dev/play/p/s9IF9NUVA-y)

Output:
```json
{"error":{"id":0,"code":500,"message":"nil pointer"}}
{"error":{"id":2,"code":404,"message":"config.json does not exist"}}
{"error":{"id":8,"code":500,"message":"value of port is 8000, but expected to be equal to 3000"}}
```

## Documentation & Resources

- [Go Package Documentation](https://pkg.go.dev/github.com/thorsphere/tserr) — Complete API reference
- [Go Report Card](https://goreportcard.com/report/github.com/thorsphere/tserr) — Code quality metrics
- [Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Ftserr) — Dependency analysis

## License

Copyright (c) 2023-2026 thorsphere. Licensed under the GNU Affero General Public License v3.0. See [LICENSE](LICENSE) for details.
