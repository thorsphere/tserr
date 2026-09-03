// Package tserr exports error functions following three calling patterns:
//
//  1. No arguments: Simple errors requiring no parameters,
//     e.g., NilPtr() (defined in tserr_nil.go)
//
//  2. Single argument: Errors with one parameter passed directly,
//     e.g., NotExistent("config.json")
//
//  3. Multiple arguments: Errors with multiple parameters passed via struct pointer,
//     e.g., EqualStr(&tserr.EqualStrArgs{Var: "port", Actual: "9000", Want: "3000"})
//
// For pattern 3 (struct pointer arguments), each error function checks if the struct
// pointer is nil. If nil, it returns NilPtr(). Otherwise, it formats the error message:
//
//	if a == nil {
//	    return NilPtr()
//	}
//	return errorf(&errmsgEqualStr, a.Var, a.Actual, a.Want)
//
// All exported error functions except NilPtr are defined here. NilPtr is in tserr_nil.go.
//
// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package tserr

// CheckArgs holds the required arguments for the error function Check
type CheckArgs struct {
	// F is the name of the object causing the failed check, for example, a filename
	F string
	// Err is the error causing the failed check, for example, 'file is a directory'
	Err error
}

// Check can be used for negative validations on an object.
func Check(a *CheckArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgCheck, a.F, a.Err)
}

// NotExistent can be used if a required object does not exist, for example, a file.
// F is the name of the object, for example, a filename
func NotExistent(F string) error {
	return errorf(&errmsgNotExistent, F)
}

// AlreadyExistent can be used if a required object already exists, for example, a file, but is not expected to exist.
// F is the name of the object, for example, a filename
func AlreadyExistent(F string) error {
	return errorf(&errmsgAlreadyExistent, F)
}

// OpArgs holds the required arguments for the error function Op
type OpArgs struct {
	// Op is the name of the failed operation, for example, 'WriteStr'
	Op string
	// Fn is the name of the object passed to the operation, for example, a filename
	Fn string
	// Err is the error returned by the failed operation, for example, 'file does not exist'
	Err error
}

// Op can be used for failed operations on an object.
func Op(a *OpArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgOp, a.Op, a.Fn, a.Err)
}

// NilFailed can be used if a function implementing an operation returns nil when an error was expected (e.g., in unit tests).
// Op is the name of the operation, for example, 'ExistsFile'
func NilFailed(Op string) error {
	return errorf(&errmsgNilFailed, Op)
}

// NilExpected can be used if a function implementing an operation returns an error when nil was expected (e.g., in unit tests).
// Op is the name of the operation, for example, 'ExistsFile'
func NilExpected(Op string) error {
	return errorf(&errmsgNilExpected, Op)
}

// Empty can be used if a required object is empty but not allowed to be empty, for example, an input argument of type string.
// F is the name of the empty object, for example, a parameter or filename
func Empty(F string) error {
	return errorf(&errmsgEmpty, F)
}

// EqualStrArgs holds the required arguments for the error function EqualStr
type EqualStrArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual string
	// Want is the expected value of Var
	Want string
}

// EqualStr can be used if a string is not equal to an expected string.
func EqualStr(a *EqualStrArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgEqualStr, a.Var, a.Actual, a.Want)
}

// TypeNotMatchingArgs holds the required arguments for the error function TypeNotMatching
type TypeNotMatchingArgs struct {
	// Actual is the actual type of the object, for example, 'file'
	Actual string
	// Want is the expected or required type of the object, for example, 'directory'
	Want string
}

// TypeNotMatching can be used if the type of an object does not match the expected type.
func TypeNotMatching(a *TypeNotMatchingArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgTypeNotMatching, a.Actual, a.Want)
}

// Forbidden can be used if an operation on an object is forbidden.
// F is the name of the forbidden object, for example, a directory or filename
func Forbidden(F string) error {
	return errorf(&errmsgForbidden, F)
}

// ReturnArgs holds the required arguments for the error function Return
type ReturnArgs struct {
	// Op is the operation name
	Op string
	// Actual is the actual return value returned by Op
	Actual string
	// Want is the expected return value from Op
	Want string
}

// Return can be used if an operation returns an actual value, but another return value was expected.
func Return(a *ReturnArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgReturn, a.Op, a.Actual, a.Want)
}

// HigherArgs holds the required arguments for the error function Higher
type HigherArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual int64
	// LowerBound is the lower bound. Actual is expected to be equal to or higher than LowerBound
	LowerBound int64
}

// Higher can be used if an integer fails to be at least equal to or higher than a defined lower bound.
func Higher(a *HigherArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgHigher, a.Var, a.Actual, a.LowerBound)
}

// EqualIntArgs holds the required arguments for the error function EqualInt
type EqualIntArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual int64
	// Want is the expected value of Var
	Want int64
}

// EqualInt can be used if an integer fails to be equal to an expected value.
func EqualInt(a *EqualIntArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgEqualInt, a.Var, a.Actual, a.Want)
}

// EqualArgs holds the required arguments for the error function Equal
type EqualArgs struct {
	// Name of the first variable
	X string
	// Name of the second variable
	Y string
}

// Equal can be used if two variables are not equal when they are expected to be equal.
func Equal(a *EqualArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgEqual, a.X, a.Y)
}

// LowerArgs holds the required arguments for the error function Lower
type LowerArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual int64
	// HigherBound is the upper bound. Actual is expected to be lower than HigherBound
	HigherBound int64
}

// Lower can be used if an integer fails to be lower than a defined upper bound.
func Lower(a *LowerArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgLower, a.Var, a.Actual, a.HigherBound)
}

// NotSet can be used if a required object or configuration is not set, for example, an environment variable.
// F is the name of the unset object, for example, an environment variable name
func NotSet(F string) error {
	return errorf(&errmsgNotSet, F)
}

// NotAvailableArgs holds the required arguments for the error function NotAvailable
type NotAvailableArgs struct {
	// S is the name of the unavailable service
	S string
	// Err is the underlying error returned by the service
	Err error
}

// NotAvailable can be used if a service is not available.
func NotAvailable(a *NotAvailableArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgNotAvailable, a.S, a.Err)
}

// EqualfArgs holds the required arguments for the error function Equalf
type EqualfArgs struct {
	// Var is the name of the variable
	Var string
	// Actual is the actual value of Var
	Actual float64
	// Want is the expected value of Var
	Want float64
}

// Equalf can be used if a float value is not equal to an expected value.
func Equalf(a *EqualfArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgEqualf, a.Var, a.Actual, a.Want)
}

// NonPrintable can be used if a string contains non-printable runes when only printable runes are permitted.
// F is the name of the string that should only contain printable runes
func NonPrintable(F string) error {
	return errorf(&errmsgNonPrintable, F)
}

// NotEqualArgs holds the required arguments for the error function NotEqual
type NotEqualArgs struct {
	// Name of the first variable
	X string
	// Name of the second variable
	Y string
}

// NotEqual can be used if two variables are equal when they are not permitted to be equal.
func NotEqual(a *NotEqualArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgNotEqual, a.X, a.Y)
}

// Duplicate can be used if an object already exists but must be unique, for example, a key or identifier.
// F is the name of the duplicate object, for example, a key name
func Duplicate(F string) error {
	return errorf(&errmsgDuplicate, F)
}

// Locked can be used if a resource or service is locked, for example, because an operation is still in progress.
// S is the name of the locked service or resource
func Locked(S string) error {
	return errorf(&errmsgLocked, S)
}

// MethodNotAllowedArgs holds the required arguments for the error function MethodNotAllowed
type MethodNotAllowedArgs struct {
	// Method is the name of the disallowed method
	Method string
	// Resource is the name of the target resource
	Resource string
}

// MethodNotAllowed can be used if an operation or HTTP method is not allowed on a given resource.
func MethodNotAllowed(a *MethodNotAllowedArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgMethodNotAllowed, a.Method, a.Resource)
}

// InvalidJson can be used if a JSON payload is invalid and cannot be parsed.
// Err is the error that occurred while parsing the JSON
func InvalidJson(Err error) error {
	return errorf(&errmsgInvalidJson, Err)
}

// InvalidFormat can be used if a string is invalid and cannot be parsed according to the expected format.
// S is the string that has an invalid format
func InvalidFormat(S string) error {
	return errorf(&errmsgInvalidFormat, S)
}

// InvalidTimestampFormat can be used if a timestamp string cannot be parsed.
// Err is the error that occurred while parsing the timestamp
func InvalidTimestampFormat(Err error) error {
	return errorf(&errmsgInvalidTimestampFormat, Err)
}

// StatusNotMatchingArgs holds the required arguments for the error function StatusNotMatching
type StatusNotMatchingArgs struct {
	// Expected is the expected status
	Expected int
	// Actual is the actual status received
	Actual int
}

// StatusNotMatching can be used if an actual status does not match the expected status.
func StatusNotMatching(a *StatusNotMatchingArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgStatusNotMatching, a.Expected, a.Actual)
}

// NotFound can be used if a required object cannot be found, for example, a file or a database record.
// F is the name of the missing object, for example, a filename or key
func NotFound(F string) error {
	return errorf(&errmsgNotFound, F)
}

// UnexpectedField can be used if an unexpected field is encountered in parsed input.
// F is the name of the unexpected field
func UnexpectedField(F string) error {
	return errorf(&errmsgUnexpectedField, F)
}

// UnexpectedErrorArgs holds the required arguments for the error function UnexpectedError
type UnexpectedErrorArgs struct {
	// Expected is the expected error
	Expected error
	// Actual is the received error
	Actual error
}

// UnexpectedError can be used if a received error does not match the expected error (e.g., in unit tests).
func UnexpectedError(a *UnexpectedErrorArgs) error {
	if a == nil {
		return NilPtr()
	}
	return errorf(&errmsgUnexpectedError, a.Expected, a.Actual)
}

// NoChanges can be used if no changes are found (for example, no staged Git changes).
// F is the name of the object without changes
func NoChanges(F string) error {
	return errorf(&errmsgNoChanges, F)
}

// Aborted can be used if an operation was aborted by the user (e.g., rejecting a confirmation prompt).
// Op is the name of the operation that was aborted
func Aborted(Op string) error {
	return errorf(&errmsgAborted, Op)
}
