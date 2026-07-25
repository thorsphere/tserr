# tserr

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorsphere/tserr)](https://pkg.go.dev/mod/github.com/thorsphere/tserr)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorsphere/tserr)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorsphere/tserr)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorsphere/tserr)
[![CodeFactor](https://www.codefactor.io/repository/github/thorsphere/tserr/badge)](https://www.codefactor.io/repository/github/thorsphere/tserr)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorsphere/tserr)

**tserr** is a Go monorepo providing a lightweight, zero-dependency package for generating standardized, structured error messages in JSON format, along with a powerful code generator companion.

## Monorepo Structure

This repository contains two main components:

| Component | Path | Description |
| --------- | ---- | ----------- |
| **tserr Library** | `/` | The core error handling library. Provides consistent, HTTP-aligned JSON errors. |
| **tserr Generator**| `/gen` | The code generator tool (formerly `tserrgen`) used to automate error creation. |

---

## Part 1: The `tserr` Library

The main library provides a simple, consistent approach to error handling without any external dependencies.

### Key Features

- **Structured Output**: All errors are formatted as JSON for easy parsing and logging
- **Zero Dependencies**: Only uses the [Go Standard Library](https://pkg.go.dev/std)
- **Simple API**: Just function calls, no configuration needed
- **Tested**: High code coverage with comprehensive unit tests
- **HTTP-Aligned**: Error codes correspond to HTTP status codes for consistency

### Installation

Add the library to your Go project:

```bash
go get github.com/thorsphere/tserr
```

### Usage Styles

Import the package in your Go code (`import "github.com/thorsphere/tserr"`). `tserr` supports different calling patterns:

#### 1. Simple Errors (No Arguments)
For straightforward errors without parameters:
```go
err := tserr.NilPtr()
```

#### 2. Single-Argument Errors
For errors with one parameter:
```go
f := "foo.txt"
err := tserr.NotExistent(f)
```

#### 3. Multi-Argument Errors
For errors requiring multiple parameters, pass a pointer to a struct:
```go
err := tserr.EqualStr(&tserr.EqualStrArgs{
    Var:    "username",
    Actual: "alice",
    Want:   "bob",
})
```
*Note: Multi-argument error functions check if the struct pointer is nil before processing. If nil, they return `tserr.NilPtr()`.*

### Output Format

Every error is formatted as a JSON string with consistent structure:

```json
{
  "error": {
    "id": 8,
    "code": 500,
    "message": "value of username is alice, but expected to be equal to bob"
  }
}
```

- **`id`**: A unique, incrementally-numbered error identifier.
- **`code`**: An HTTP status code corresponding to the error category.
- **`message`**: The error message (can contain formatted values from arguments).

### Example Code

```go
package main

import (
	"fmt"
	"github.com/thorsphere/tserr"
)

func main() {
	// Simple & Single-argument
	fmt.Println(tserr.NilPtr())
	fmt.Println(tserr.NotExistent("config.json"))

	// Multi-argument error
	fmt.Println(tserr.EqualStr(&tserr.EqualStrArgs{
		Var:    "port",
		Actual: "8000",
		Want:   "3000",
}
```

[Run in Go Playground](https://go.dev/play/p/s9IF9NUVA-y)

---

## ⚡ Part 2: The `tserr` Generator (`/gen`)

The generator automates the creation of custom error wrappers and types, streamlining the usage of `tserr` across large applications.

### Installation

You can install the generator globally via `go install`:

```bash
go install github.com/thorsphere/tserr/gen@latest
```

### Usage

Please see the [Generator Documentation](./gen/README.md) for detailed instructions, configuration options, and usage examples.

---

## Documentation & Resources

- [Go Package Documentation](https://pkg.go.dev/github.com/thorsphere/tserr) — Complete API reference
- [Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Ftserr) — Dependency analysis

---

## ⚖️ License & Commercial Usage

Copyright (c) 2023-2026 thorsphere. All rights reserved.

This project is licensed under the **Functional Source License v1.1 (FSL-1.1-ALv2)**. 

* The use, modification, and redistribution of this Go package is completely free for private, educational, non-commercial, and internal purposes. 
* If you are a company or institution looking to use this package in a commercial product, service, or business environment, you must secure a commercial license.
* Each version of this software automatically converts to the fully open-source Apache License, Version 2.0 on the second anniversary of its release.

For full details, please see the [LICENSE](LICENSE.md) file.

### 💼 Commercial Licensing & Inquiries

To purchase a commercial license or discuss support options, please reach out directly:

* 📩 **Contact:** business at thorsphere dot com
* 💬 **Response Time:** Usually within a couple of business days.

*Please include your company name and a brief overview of your use case so I can provide the right licensing details.*
