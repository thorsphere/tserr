// Package tserr provides a simple API for standardized error messages
// in the JSON format.
//
// The tserr package provides easy-to-use functions to get standardized
// error messages for typical errors. The error messages are fomatted in
// the JSON format:
//
//	{"error":{"id":<int>,"code":<int>,"message":"<string>"}}
//
// The root element is named "error". "id" is a consecutively numbered
// id. "code" is a relating HTTP status code. "message" contains the actual
// pre-defined error message.
//
// The error message may contain verbs to be filled by arguments. The
// arguments for the verbs are provided as function arguments. A function
// may hold one argument used as one verb in the error message. A function
// may hold multiple arguments used for more than one verb in the error
// message. Multiple arguments are passed to a function as a pointer to
// a struct, e.g.,
//
//	err := tserr.EqualStr(&tserr.EqualStrArgs{X: "test1", Y: "test2"})
//
// Output with fmt.Println(err):
//
//	{"error":{"id":6,"code":500,"message":"test1 does not equal test2"}}
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
