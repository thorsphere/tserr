# tserr/gen

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thorsphere/tserr/gen)](https://pkg.go.dev/github.com/thorsphere/tserr/gen)

This is the code generation component of the [tserr](../README.md) monorepo. The `gen` package automates the creation of error handling code for the core `tserr` library. It reads error definitions from a JSON configuration file and generates the corresponding error types, API functions, and comprehensive test functions. 

The generator is tightly integrated with the `tserr` library in the root directory to heavily minimize boilerplate code for custom error types.

- **Automated**: Generates complete error handling implementation from simple JSON configuration
- **Tested**: Automatically generates comprehensive test functions for all error cases
- **Configuration-driven**: Simple JSON schema makes it easy to define and extend errors
- **Dependencies**: Minimal dependencies, relying on [Go Standard Library](https://pkg.go.dev/std), [lpcode](https://pkg.go.dev/github.com/thorsphere/lpcode), [tsfio](https://pkg.go.dev/github.com/thorsphere/tsfio) and the root `tserr` package.

## Usage

In your Go code, import the package as follows:

```go
import "github.com/thorsphere/tserr/gen"
```

The package provides automated code generation for error handling based on JSON configuration files. The primary entry point is the Generate function, which orchestrates the entire code generation process.

### Basic Workflow

The generation process follows these steps:

1. Define your errors in a JSON configuration file (e.g., tserr.json)
2. Call Generate with the path to your configuration file
3. The generator produces three Go files:
- `tserr_messages.go` - Error message constants
- `tserr_api.go` - Error constructor functions
- `tserr_api_test.go` - Comprehensive test functions

```go
func main() {
    if err := gen.Generate(tsfio.Filename("tserr.json")); err != nil {
        log.Fatal(err)
    }
}
```

### Configuration File

The JSON configuration file defines error definitions with their names, comments, HTTP status codes, message templates, and parameters:

```json
{
  "tserr": {
    "path": "../tserr",
    "version": "1.0.0",
    "errors": [
      {
        "name": "NotFound",
        "comment": "NotFound is returned when a resource is not found.",
        "code": "404",
        "message": "resource %v not found",
        "param": [
          {
            "name": "resource",
            "comment": "the resource type that was not found",
            "type": "string"
          }
        ]
      }
    ]
  }
}
```

### Generated code

For each error definition, the generator creates:

- Error message structs in `tserr_messages.go` - Constants holding error metadata
- Constructor functions in `tserr_api.go` - Type-safe functions for creating errors
- Test functions in `tserr_api_test.go` - Comprehensive tests validating error generation

### Core Function

```go
func Generate(fn tsfio.Filename) error
```

The `Generate` function reads the JSON configuration file and generates the complete error handling implementation. It validates the configuration and returns an error if any step fails.

### Requirements

Generated files require header and footer template files with suffixes `.header` and `.footer` to be present in the target directory. These files contain package declarations, imports, and other boilerplate that should appear at the beginning or end of generated files:

- `tserr_messages.go.header` and `tserr_messages.go.footer`
- `tserr_api.go.header` and `tserr_api.go.footer`
- `tserr_api_test.go.header` and `tserr_api_test.go.footer`

## Example

```go
package main

import (
    "github.com/thorsphere/tserr/gen"
)

func main() {
    if e := gen.Generate("tserr.json"); e != nil {
        panic(e)
    }
}
```

## Limitations

The generator does not validate or check for non-printable characters in the `tserr.json` configuration file. When defining error names, comments, or message templates, ensure that your JSON configuration contains only printable characters. Non-printable characters (control characters, invalid UTF-8 sequences, etc.) may result in corrupted or invalid generated error messages.

**Recommendation**: Use only standard ASCII printable characters and valid UTF-8 sequences in your configuration file.

## Links

[Godoc](https://pkg.go.dev/github.com/thorsphere/tserr/gen)

