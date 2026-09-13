// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package gen_test

// Import packages for testing.
import (
	"encoding/json" // for json.Unmarshal
	"fmt"           // for fmt.Sprintf
	"testing"       // for testing.T

	"github.com/thorsphere/tserr"     // for error handling
	"github.com/thorsphere/tserr/gen" // to test the gen package
	"github.com/thorsphere/tsfio"     // for tsfio.Filename and tsfio.ReadFile
)

// TestGenerate tests the Generate function of the tserrgen package.
// It also generates tserr package. The generated code is not checked in this test,
// but it can be manually inspected after running the test. It fails if there is an error during generation.
func TestGenerate(t *testing.T) {
	// Call the Generate function with the path to the JSON file containing error definitions.
	if e := gen.Generate("tserr.json"); e != nil {
		// If there is an error, fail the test and print the error message.
		t.Error(e)
	}
}

// TestValidate tests the Validate function of the tserrgen package.
// It reads the JSON file containing error definitions and validates the configuration.
// It fails if there is an error during validation.
func TestValidate(t *testing.T) {
	// Read the JSON file containing error definitions.
	fn := tsfio.Filename("tserr.json")
	b, err := tsfio.ReadFile(fn)
	// If there is an error, return an error with details about the operation that failed.
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	// Unmarshal the JSON data into a tserrconfig struct.
	var m gen.Tserrconfig
	if e := json.Unmarshal(b, &m); e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "Unmarshal", Fn: string(fn), Err: e}))
	}
	// The real configuration must pass validation.
	if e := gen.Validate(&m); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "validate", Fn: string(fn), Err: e}))
	}
}

// TestValidateErr tests the Validate function with invalid error definitions.
// It fails if validation does not return an error for each case.
func TestValidateErr(t *testing.T) {
	// Test cases: too many verbs, too many params, malformed verb.
	m := []gen.Errmsg{
		{Name: "Check", Msg: "check %v failed: %w",
			Param: []gen.Param{{Name: "F", Type: "string"}}},
		{Name: "NotExistent", Msg: "%v does not exist",
			Param: []gen.Param{{Name: "F", Type: "string"}, {Name: "Err", Type: "error"}}},
		{Name: "NotExistent", Msg: "broken %",
			Param: []gen.Param{{Name: "F", Type: "string"}}},
	}
	// Test each case.
	for i, v := range m {
		// Create a tserrconfig struct with the error definition.
		var c gen.Tserrconfig
		// Set the root error.
		c.Root.Errors = []gen.Errmsg{v}
		// The configuration should fail validation.
		if e := gen.Validate(&c); e == nil {
			t.Error(tserr.NilFailed(fmt.Sprintf("validate case %d", i)))
		}
	}
}
