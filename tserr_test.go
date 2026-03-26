// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tserr

// Import standard library packages
import (
	"encoding/json" // encoding/json
	"fmt"           // fmt
	"testing"       // testing
)

// errNil and errNNil hold standard error strings when tests fail
var (
	// a function under test returns nil, but an error is expected
	errNil string = "nil returned, but error expected"
	// a function under test returns an error, but nil is expected
	errNNil string = "error returned, but nil expected"
)

// TestErrorfNil tests the return value of errorf if nil is provided.
// The test fails if errorf returns nil.
func TestErrorfNil(t *testing.T) {
	err := errorf(nil)
	if err == nil {
		t.Error(errNil)
	}
}

// TestNil tests the return value of NilPtr():
// - checks if returned value is not nil
// - checks if returned error message is in valid JSON format
// - checks if returned error message matches expected error message
func TestNil(t *testing.T) {
	// call NilPtr and retrieve error
	err := NilPtr()
	// If error is nil, then test fails immediately
	if err == nil {
		t.Fatal(errNil)
	}
	// check returned error message if it is in valid JSON format
	testValidJson(t, err)
	// check returned error message if it equals the expected error message
	testEqualJson(t, err, &nilPtr)
}

// testValidJson tests if the error messag eis in valid JSON format
// and returns an error if not
func testValidJson(t *testing.T, e error) {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// convert error to []byte type and check for valid JSON format using
	// the encoding/json standard library package
	if !json.Valid([]byte(fmt.Sprintf("%v", e))) {
		// report error in case the error message is not in valid JSON format
		t.Error("not in valid json format")
	}
}

// testEqualJson tests if the error message in e equals the expected error message
// in emsg. The expected error message is constructed by using Marshal of the
// encoding/json standard library package. It returns an error if the expected
// error message does not equal the actual error message
func testEqualJson(t *testing.T, e error, emsg *errmsg) {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// if pointer emsg is nil then test fails immediately
	if emsg == nil {
		t.Fatal("nil pointer")
	}
	// generate expected error message by using Marshal from the
	// encoding/json standard library package
	j, errj := json.Marshal(&errwrap{*emsg})
	// if Marshal returns an error, test fails immediately
	if errj != nil {
		t.Fatal(errj)
	}
	// return an error, if error messages do not match
	if fmt.Sprintf("%v", e) != string(j) {
		t.Errorf("%v does not equal %v", e, string(j))
	}
}

// BenchmarkNil performs a benchmark calling NilPtr.
func BenchmarkNil(b *testing.B) {
	var err [2]error
	// Reset benchmark timer
	b.ResetTimer()
	// Run benchmark with NilPtr()
	for i := 0; i < b.N; i++ {
		err[i&0x1] = NilPtr()
	}
}
