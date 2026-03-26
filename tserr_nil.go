package tserr

// NilPtr error function is self-used internally
// in the package, so it is provided in this separate source file.
//
// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.

// nilPtr error message is id 0 with error code 500 and a simple
// error message without verbs.
var (
	nilPtr = errmsg{0, 500, "nil pointer"}
)

// NilPtr just provides the error message and does not have arguments.
func NilPtr() error {
	return errorf(&nilPtr)
}
