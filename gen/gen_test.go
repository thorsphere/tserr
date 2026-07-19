// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package gen_test

// Import Go standard library packages and the tserr/gen package.
import (
	"testing" // testing

	"github.com/thorsphere/tserr/gen" // gen
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
