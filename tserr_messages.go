// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tserr

// Import Go standard package http
import "net/http" // http

// Error ids, error codes and error messages with their potential verbs.
var (
	errmsgCheck           = errmsg{1, http.StatusPreconditionFailed, "check %v failed: %w"}
	errmsgNotExistent     = errmsg{2, http.StatusNotFound, "%v does not exist"}
	errmsgAlreadyExistent = errmsg{3, http.StatusConflict, "%v already exists"}
	errmsgOp              = errmsg{4, http.StatusUnprocessableEntity, "%v %v failed: %w"}
	errmsgNilFailed       = errmsg{5, http.StatusInternalServerError, "%v returned nil, but error expected"}
	errmsgNilExpected     = errmsg{6, http.StatusInternalServerError, "%v did not return nil, but nil expected"}
	errmsgEmpty           = errmsg{7, http.StatusBadRequest, "%v cannot be empty"}
	errmsgEqualStr        = errmsg{8, http.StatusInternalServerError, "value of %v is %v, but expected to be equal to %v"}
	errmsgTypeNotMatching = errmsg{9, http.StatusMethodNotAllowed, "%v does not match type %v"}
	errmsgForbidden       = errmsg{10, http.StatusForbidden, "operation on %v forbidden"}
	errmsgReturn          = errmsg{11, http.StatusInternalServerError, "%v returned %v, but %v expected"}
	errmsgHigher          = errmsg{12, http.StatusInternalServerError, "value of %v is %d, but expected to be at least equal to or higher than %d"}
	errmsgEqual           = errmsg{14, http.StatusInternalServerError, "value of %v is %d, but expected to be equal to %d"}
	errmsgLower           = errmsg{15, http.StatusInternalServerError, "value of %v is %d, but expected to be lower than %d"}
	errmsgNotSet          = errmsg{16, http.StatusNotFound, "%v not set"}
	errmsgNotAvailable    = errmsg{17, http.StatusServiceUnavailable, "%v not available: %w"}
	errmsgEqualf          = errmsg{18, http.StatusInternalServerError, "value of %v is %f, but expected to be equal to %f"}
	errmsgNonPrintable    = errmsg{19, http.StatusBadRequest, "%v contains non-printable runes, but only printable runes are allowed"}
	errmsgNotEqual        = errmsg{20, http.StatusInternalServerError, "variable %v equals variable %v, but not allowed to equal"}
	errmsgDuplicate       = errmsg{21, http.StatusForbidden, "%v is a duplicate and already exists"}
	errmsgLocked          = errmsg{22, http.StatusLocked, "%v is locked"}
)
