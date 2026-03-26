// Package tserr provides a small, opinionated toolkit for structured error
// messages in JSON format, with an HTTP-status-code-aligned code.
//
// Each function returns an error whose string is JSON like:
//
//	{"error":{"id":<int>,"code":<int>,"message":"<string>"}}
//
// - `id` is a package-level error identifier (incremental, `0` = nil pointer).
// - `code` is a related HTTP status code (e.g. 400, 404, 409, 500).
// - `message` is a pre-defined message template (supports fmt-style verbs).
//
// Two patterns exist:
// 1. Single-arg errors via direct function params, e.g. `tserr.Empty("path")`.
// 2. Multi-arg errors via struct pointer args, e.g.:
//
//	err := tserr.EqualStr(&tserr.EqualStrArgs{
//	    Var: "name", Actual: "foo", Want: "bar",
//	})
//
// If a multi-arg struct pointer is nil, `tserr.NilPtr()` is returned.
// Otherwise the template is formatted and wrapped as JSON.
//
// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tserr

// Import standard library packages
import "fmt" // fmt

// Struct errmsg contains content of the error message.
//   - Id: consecutively numbered error id as integer; JSON element "id"
//     NilPtr() always returns id 0.
//   - C: relating HTTP status code as integer; JSON element "code"
//   - M: error message as string, which may contain verbs; JSON element "message"
type errmsg struct {
	Id int    `json:"id"`      // id
	C  int    `json:"code"`    // error code (HTTP status code)
	M  string `json:"message"` // error message
}

// Struct errwrap is the root element holding the error message.
type errwrap struct {
	E errmsg `json:"error"` // root element
}

// errformat holds the JSON format of the error message with id, code and
// message as verb.
var (
	errformat string = "{" +
		"\"error\":{" +
		"\"id\":%d," +
		"\"code\":%d," +
		"\"message\":\"%w\"" +
		"}" +
		"}"
)

// errorf returns the JSON formatted error based on the provided pointer to
// the error message provided as struct errmsg. The error message may contain
// verbs. The contents of the verbs is provided by optional additional arguments.
func errorf(e *errmsg, a ...any) error {
	// If the pointer to struct errmsg is nil, then return nilPtr error
	if e == nil {
		// Note: does not call Nilptr(), because NilPtr() calls errorf,
		// in worst case ending up in an infinite loop calling NilPtr().
		return fmt.Errorf(errformat, nilPtr.Id, nilPtr.C, nilPtr.M)
	}
	// Return error in JSON format with id, code and error message.
	return fmt.Errorf(errformat, e.Id, e.C, fmt.Errorf(e.M, a...))
}
