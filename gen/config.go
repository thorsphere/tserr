// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package gen

// tserrconfig defines the structure of the JSON configuration file for generating Go code
// for the tserr package.
type tserrconfig struct {
	Root errorsconfig `json:"tserr"` // Root element
}

// errorsconfig defines the structure of the "tserr" element in the JSON configuration file. It contains
// the path to the tserr package, the version of the tserr package, and a list of error definitions.
type errorsconfig struct {
	Path   string   `json:"path"`    // Path to tserr package
	Ver    string   `json:"version"` // Version of tserr package
	Errors []errmsg `json:"errors"`  // Errors
}

// errmsg defines the structure of each error definition in the JSON configuration file. It contains
// the name of the error, a comment describing the error, the HTTP status code associated with the error,
// the error message (which may contain verbs), and a list of parameters for the error.
type errmsg struct {
	Name    string  `json:"name"`    // Error name
	Comment string  `json:"comment"` // Comment
	Code    string  `json:"code"`    // Error code (HTTP status code from Go standard package http)
	Msg     string  `json:"message"` // Error messages (may contain verbs)
	Param   []param `json:"param"`   // Error parameters
}

// param defines the structure of each parameter in the error definitions in the JSON configuration file.
// It contains the name of the parameter, a comment describing the parameter, and the type of the parameter. The
type param struct {
	Name    string `json:"name"`    // Parameter name
	Comment string `json:"comment"` // Parameter comment
	Type    string `json:"type"`    // Parameter type
}
