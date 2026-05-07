// Package tserr/gen provides automated Go code generation for the tserr package.
// It reads error definitions from a JSON configuration file and generates the
// corresponding error types, API functions, and test functions.
//
// The main component is:
//
//   - Generate: The primary entry point that orchestrates the entire code generation
//     process. It reads a JSON configuration file, parses the error definitions,
//     and invokes the appropriate generators for messages, API functions, and tests.
//
// The JSON configuration file defines errors with their names, comments, HTTP status
// codes, message templates, and parameter specifications. The generator creates type-safe,
// testable error functions that integrate seamlessly with the tserr error handling package.
//
// The package leverages lpcode for programmatic code generation, tsfio for file
// operations, and tserr for consistent error handling throughout the generation process.
//
// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package gen

// Import Go standard library packages, lpcode, tserr and tsfio
import (
	"encoding/json" // json
	"fmt"           // fmt

	"github.com/thorsphere/lpcode" // lpcode
	"github.com/thorsphere/tserr"  // tserr
	"github.com/thorsphere/tsfio"  // tsfio
)

// Define variables for filenames and test variables used in the code generation process.
var (
	tserr_messages_go = tsfio.Filename("tserr_messages.go")
	tserr_api_go      = tsfio.Filename("tserr_api.go")
	tserr_api_test_go = tsfio.Filename("tserr_api_test.go")
	tserr_testvars    = []lpcode.TestVar{
		{T: "string", N: "strFoo", V: "\"foo\""},
		{T: "error", N: "errFoo", V: "fmt.Errorf(\"foo error\")"},
		{T: "int64", N: "int64Foo", V: "42"},
		{T: "int", N: "intFoo", V: "7"},
		{T: "float64", N: "floatFoo", V: "314"},
	}
)

// Generate generates Go code for the tserr package based on the error definitions
// provided in the JSON file specified by fn. It reads the JSON file,
// unmarshals it into a tserrconfig struct, and then generates the necessary Go code
// for error messages and API functions. If any error occurs during this process,
// it returns an error with details about the operation that failed.
func Generate(fn tsfio.Filename) error {
	// Read the JSON file specified by fn.
	b, err := tsfio.ReadFile(fn)
	// If there is an error, return an error with details about the operation that failed.
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err})
	}

	// Unmarshal the JSON data into a tserrconfig struct.
	var m tserrconfig
	if e := json.Unmarshal(b, &m); e != nil {
		// If there is an error,return an error with details about the operation that failed.
		return tserr.Op(&tserr.OpArgs{Op: "Unmarshal", Fn: string(fn), Err: err})
	}

	// Generate the Go code for error messagesbased on the unmarshaled configuration.
	if e := genMessages(&m); e != nil {
		// If there is an error, return an error with details about the operation that failed.
		return e
	}

	// Generate the Go code for API functions based on the unmarshaled configuration.
	if e := genApi(&m, tserr_api_go, genApiFunc, nil); e != nil {
		// If there is an error, return an error with details about the operation that failed.
		return e
	}

	// Generate the Go code for API test functions based on the unmarshaled configuration.
	if e := genApi(&m, tserr_api_test_go, genApiTestFunc, tserr_testvars); e != nil {
		// If there is an error, return an error with details about the operation that failed.
		return e
	}

	// If everything is successful, return nil to indicate that the generation process completed without errors.
	return nil
}

// genMessages generates Go code for error messages based on the provided tserrconfig struct.
// It creates a new code file for error messages, iterates through the error definitions
// in the configuration, and writes the corresponding Go code for each error message.
// If any error occurs during this process, it returns an error with details about
// the operation that failed.
func genMessages(m *tserrconfig) error {
	// Return an error if the input tserrconfig pointer is nil
	if m == nil {
		return tserr.NilPtr()
	}
	// Create a new code file for error messages using the path specified in the configuration
	// and the filename for error messages.
	cf, err := lpcode.NewCodefile(tsfio.Directory(m.Root.Path), tserr_messages_go)
	// If there is an error, return an error with details about the operation that failed.
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "NewCodefile", Fn: m.Root.Path + string(tserr_messages_go), Err: err})
	}
	// Start writing to the code file.
	if e := cf.StartFile(); e != nil {
		// If there is an error, return an error with details about the operation that failed.
		return tserr.Op(&tserr.OpArgs{Op: "StartFile", Fn: string(cf.Filepath()), Err: e})
	}
	// Iterate through the error definitions in the configuration and write the corresponding Go code for each error message.
	for i, v := range m.Root.Errors {
		// Skip error code 13.
		if i++; i >= 13 {
			i++
		}
		// Create a new code snippet for the error message using the lpcode package.
		// The code snippet includes the error code and the error message.
		c := lpcode.NewCode().Assignment(&lpcode.AssignmentArgs{ExprLeft: "errmsg" + v.Name}).CompositeLit("errmsg").Ident(fmt.Sprint(i)).List().Ident(v.Code).List().Ident("\"" + v.Msg + "\"").BlockEnd()
		// Write the generated code snippet to the code file.
		if e := cf.WriteCode(c); e != nil {
			// If there is an error, return an error with details about the operation that failed.
			return tserr.Op(&tserr.OpArgs{Op: "WriteCode", Fn: string(cf.Filepath()), Err: e})
		}
	}
	// Finish writing to the code file.
	if e := cf.FinishFile(); e != nil {
		// If there is an error, return an error with details about the operation that failed.
		return tserr.Op(&tserr.OpArgs{Op: "FinishFile", Fn: string(cf.Filepath()), Err: e})
	}
	// If everything is successful, return nil to indicate that the generation process completed without errors.
	return nil
}

// genApi generates Go code for API functions based on the provided tserrconfig struct, filename, code generation function, and test variables.
// It creates a new code file for API functions, iterates through the error definitions in the configuration,
// and writes the corresponding Go code for each API function using the provided code generation function.
// If any error occurs during this process, it returns an error with details about the operation that failed.
func genApi(m *tserrconfig, fn tsfio.Filename, genf func(*errmsg) (*lpcode.Code, error), tc []lpcode.TestVar) error {
	// Return an error if the input tserrconfig pointer is nil or if the code generation function is nil
	if (m == nil) || (genf == nil) {
		return tserr.NilPtr()
	}
	// Create a new code file for API functions using the path specified in the configuration and
	// the provided filename.
	cf, err := lpcode.NewCodefile(tsfio.Directory(m.Root.Path), fn)
	// If there is an error, return an error with details about the operation that failed.
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "NewCodefile", Fn: m.Root.Path + string(fn), Err: err})
	}
	// Start writing to the code file.
	if e := cf.StartFile(); e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "StartFile", Fn: string(cf.Filepath()), Err: e})
	}
	// If test variables are provided, write the corresponding Go code for the test variables to the code file.
	if tc != nil {
		if e := cf.WriteCode(lpcode.NewCode().TestVarDecl(tc)); e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "WriteCode", Fn: string(cf.Filepath()), Err: e})
		}
	}
	// Iterate through the error definitions in the configuration and generate the corresponding Go code
	// for each API function using the provided code generation function.
	for _, v := range m.Root.Errors {
		// Return an error if the parameters in the error message definition are nil
		if v.Param == nil {
			return tserr.NilPtr()
		}
		// Return an error if the number of parameters in the error message definition is less than or equal to 0
		if len(v.Param) == 0 {
			return tserr.Empty("Param")
		}
		// Generate the Go code for the API function using the provided code generation function and
		// write it to the code file.
		c, e := genf(&v)
		// If there is an error during the code generation, return an error with details about the operation that failed.
		if e != nil {
			return e
		}
		// Write the generated code snippet to the code file.
		if err := cf.WriteCode(c); err != nil {
			// If there is an error, return an error with details about the operation that failed.
			return tserr.Op(&tserr.OpArgs{Op: "WriteCode", Fn: string(cf.Filepath()), Err: err})
		}
	}
	// Finish writing to the code file.
	if e := cf.FinishFile(); e != nil {
		// If there is an error, return an error with details about the operation that failed.
		return tserr.Op(&tserr.OpArgs{Op: "FinishFile", Fn: string(cf.Filepath()), Err: e})
	}
	// If everything is successful, return nil to indicate that the generation process completed without errors.
	return nil
}

// genApiTestFunc generates Go code for an API test function based on the provided error message definition in m.
// If any error occurs during this process, it returns an error with details about the operation that failed.
func genApiTestFunc(m *errmsg) (*lpcode.Code, error) {
	// Return an error if the input error message definition is nil
	if m == nil {
		return nil, tserr.NilPtr()
	}
	// Return an error if the parameters in the error message definition are nil
	if m.Param == nil {
		return nil, tserr.NilPtr()
	}
	// Retrieve the number of parameters from the error message definition
	l := len(m.Param)
	// If the number of parameters is equal to 1, generate code for an API test function that takes a single parameter.
	if l == 1 {
		return genApiTestFunc1(m)
	} else if l > 1 { // If the number of parameters is greater than 1, generate code for an API test function that takes multiple parameters.
		return genApiTestFuncM(m)
	}
	// If the number of parameters is less than or equal to 0, return an error indicating that the expected number of parameters is greater than 0.
	return nil, tserr.Higher(&tserr.HigherArgs{Var: "number of parameters", Actual: int64(l), LowerBound: 1})
}

// genApiFunc generates Go code for an API function based on the provided error message definition in m.
// If any error occurs during this process, it returns an error with details about the operation that failed.
func genApiFunc(m *errmsg) (*lpcode.Code, error) {
	// Return an error if the input error message definition is nil
	if m == nil {
		return nil, tserr.NilPtr()
	}
	// Return an error if the parameters in the error message definition are nil
	if m.Param == nil {
		return nil, tserr.NilPtr()
	}
	// Retrieve the number of parameters from the error message definition
	l := len(m.Param)
	// If the number of parameters is equal to 1, generate code for an API function that takes a single parameter.
	if l == 1 {
		return genApiFunc1(m)
	} else if l > 1 { // If the number of parameters is greater than 1, generate code for an API function that takes multiple parameters.
		return genApiFuncM(m)
	}
	// If the number of parameters is less than or equal to 0, return an error indicating that the expected number of parameters is greater than 0.
	return nil, tserr.Higher(&tserr.HigherArgs{Var: "number of parameters", Actual: int64(l), LowerBound: 1})
}

// genApiTestFunc1 generates Go code for an API test function that takes a single parameter based on the provided error message definition in m.
// If any error occurs during this process, it returns an error with details about the operation that failed.
func genApiTestFunc1(m *errmsg) (*lpcode.Code, error) {
	// Return an error if the input error message definition is nil
	if m == nil {
		return nil, tserr.NilPtr()
	}
	// Return an error if the parameters in the error message definition are nil
	if m.Param == nil {
		return nil, tserr.NilPtr()
	}
	// Retrieve the number of parameters from the error message definition
	l := len(m.Param)
	// If the number of parameters is not equal to 1, return an error indicating that the expected number of parameters is 1.
	if l != 1 {
		return nil, tserr.Equal(&tserr.EqualArgs{Var: "number of parameters", Actual: int64(l), Want: 1})

	}
	// Create a new code snippet for the API test function using the lpcode package.
	c := lpcode.NewCode().Func1(&lpcode.Func1Args{Name: "Test" + m.Name, Var: "t", Type: "*testing.T", Return: ""})
	t, e := lpcode.FindTestVar(m.Param[0].Type, tserr_testvars)
	if e != nil {
		return nil, e
	}
	if t == nil {
		return nil, tserr.NilPtr()
	}
	c.ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: "a", Expr: t.N})
	c.ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: "em", Expr: "&errmsg" + m.Name})
	c.ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: "err", Expr: m.Name + "(a)"})
	c.If(&lpcode.IfArgs{ExprLeft: "err", ExprRight: "nil", Operator: "=="})
	c.SelMethod(&lpcode.SelArgs{Val: "t", Sel: "Fatal"}).Ident("errNil").ParamEndln().BlockEnd()
	c.Call("testValidJson").Ident("t").List().Ident("err").ParamEndln()
	// Generate the code for testing the error message and compare it with the actual error message returned by the API function.
	c, e = genEmsgTest(c, m)
	// If there is an error during the generation of the error message test code, return an error with details about the operation that failed.
	if e != nil {
		return nil, e
	}
	// Call the function to test the equality of the expected and actual error messages.
	c.Call("testEqualJson").Ident("t").List().Ident("err").List().Ident("&emsg").ParamEndln()
	c.FuncEnd()
	// Return the generated code snippet and nil to indicate that the generation process completed without errors.
	return c, nil
}

// genEmsgTest generates Go code for testing the error message of an API function based on the provided error message definition.
// It creates a new code snippet that constructs an error message using the parameters defined in the error message definition and
// compares it with the actual error message returned by the API function.
// If any error occurs during this process, it returns an error with details about the operation that failed.
func genEmsgTest(c *lpcode.Code, m *errmsg) (*lpcode.Code, error) {
	// Return an error if the input code snippet or error message definition is nil
	if (c == nil) || (m == nil) {
		return nil, tserr.NilPtr()
	}
	// Return an error if the parameters in the error message definition are nil
	if m.Param == nil {
		return nil, tserr.NilPtr()
	}
	// Retrieve the number of parameters from the error message definition
	l := len(m.Param)
	// If the number of parameters is less than or equal to 0, return an error indicating that the expected number of parameters is greater than 0
	if l == 0 {
		return nil, tserr.Higher(&tserr.HigherArgs{Var: "number of parameters", Actual: int64(l), LowerBound: 1})
	}
	// Create a new code snippet for testing the error message using the lpcode package
	c.ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: "emsg", Expr: "errmsg{"})
	c.SelField(&lpcode.SelArgs{Val: "em", Sel: "Id"}).Listln()
	c.SelField(&lpcode.SelArgs{Val: "em", Sel: "C"}).Listln()
	c.SelMethod(&lpcode.SelArgs{Val: "fmt", Sel: "Sprintf"}).Ident("\"%v\"").List()
	c.SelMethod(&lpcode.SelArgs{Val: "fmt", Sel: "Errorf"}).SelField(&lpcode.SelArgs{Val: "em", Sel: "M"}).List()
	if l == 1 {
		c.Ident("a")
	} else if l > 1 {
		for _, v := range m.Param {
			c.SelField(&lpcode.SelArgs{Val: "a", Sel: v.Name}).List()
		}
	}
	c.ParamEnd().ParamEnd().Listln().BlockEnd()
	// Return the generated code snippet and nil to indicate that the generation process completed without errors.
	return c, nil
}

// genApiFunc1 generates Go code for an API function that takes a single parameter
// based on the provided error message definition. If any error occurs during this process,
// it returns an error with details about the operation that failed.
func genApiFunc1(m *errmsg) (*lpcode.Code, error) {
	// Return an error if the input error message definition is nil
	if m == nil {
		return nil, tserr.NilPtr()
	}
	// Retrieve the number of parameters from the error message definition
	l := len(m.Param)
	// If the number of parameters is not equal to 1, return an error indicating that the expected number of parameters is 1
	if l != 1 {
		return nil, tserr.Equal(&tserr.EqualArgs{Var: "number of parameters", Actual: int64(l), Want: 1})
	}
	// Create a new code snippet for the API function using the lpcode package.
	c := lpcode.NewCode().LineComment(m.Comment).LineComment(m.Param[0].Comment).Func1(&lpcode.Func1Args{Name: m.Name, Var: m.Param[0].Name, Type: m.Param[0].Type, Return: "error"})
	c.Return().Call("errorf").Addr().Ident("errmsg" + m.Name).List().Ident(m.Param[0].Name).ParamEndln().FuncEnd()
	// Return the generated code snippet and nil
	// to indicate that the generation process completed without errors.
	return c, nil
}

// genApiFuncM generates Go code for an API function that takes multiple parameters
// based on the provided error message definition in m. If any error occurs
// during this process, it returns an error with details about the operation that failed.
func genApiTestFuncM(m *errmsg) (*lpcode.Code, error) {
	// Return an error if the input error message definition is nil
	if m == nil {
		return nil, tserr.NilPtr()
	}
	// Return an error if the parameters in the error message definition are nil
	if m.Param == nil {
		return nil, tserr.NilPtr()
	}
	// Retrieve the number of parameters from the error message definition
	l := len(m.Param)
	// If the number of parameters is less than or equal to 1, return an error indicating that the expected number of parameters is greater than 1
	if l <= 1 {
		return nil, tserr.Higher(&tserr.HigherArgs{Var: "number of parameters", Actual: int64(l), LowerBound: 2})
	}
	// Create a new code snippet for the API test function using the lpcode package
	c := lpcode.NewCode().Func1(&lpcode.Func1Args{Name: "Test" + m.Name + "Nil", Var: "t", Type: "*testing.T", Return: ""})
	c.IfErr(&lpcode.IfErrArgs{Method: m.Name + "(nil)", Operator: "=="}).SelMethod(&lpcode.SelArgs{Val: "t", Sel: "Errorf"})
	c.Ident("\"%s\",errNil").ParamEndln().BlockEnd().FuncEnd()
	c.Func1(&lpcode.Func1Args{Name: "Test" + m.Name, Var: "t", Type: "*testing.T", Return: ""})
	c.ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: "a", Expr: m.Name + "Args{"})
	for _, v := range m.Param {
		t, e := lpcode.FindTestVar(v.Type, tserr_testvars)
		if e != nil {
			return nil, e
		}
		if t == nil {
			return nil, tserr.NilPtr()
		}
		c.KeyedElement(&lpcode.KeyedElementArgs{Key: v.Name, Elem: t.N})
	}
	c.BlockEnd()
	c.ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: "em", Expr: "&errmsg" + m.Name})
	c.ShortVarDecl(&lpcode.ShortVarDeclArgs{Ident: "err", Expr: m.Name + "(&a)"})
	c.If(&lpcode.IfArgs{ExprLeft: "err", ExprRight: "nil", Operator: "=="})
	c.SelMethod(&lpcode.SelArgs{Val: "t", Sel: "Fatal"}).Ident("errNil").ParamEndln().BlockEnd()
	c.Call("testValidJson").Ident("t").List().Ident("err").ParamEndln()
	// Generate the code for testing the error message and compare it with the actual error message returned by the API function.
	c, e := genEmsgTest(c, m)
	// If there is an error during the generation of the error message test code, return an error with details about the operation that failed.
	if e != nil {
		return nil, e
	}
	// Call the function to test the equality of the expected and actual error messages.
	c.Call("testEqualJson").Ident("t").List().Ident("err").List().Addr().Ident("emsg").ParamEndln()
	c.FuncEnd()
	// Return the generated code snippet and nil to indicate that the generation process completed without errors.
	return c, nil
}

// genApiFuncM generates Go code for an API function that takes multiple parameters
// based on the provided error message definition in m. If any error occurs
// during this process, it returns an error with details about the operation that failed.
func genApiFuncM(m *errmsg) (*lpcode.Code, error) {
	// Return an error if the input error message definition is nil
	if m == nil {
		return nil, tserr.NilPtr()
	}
	// Return an error if the parameters in the error message definition are nil
	if m.Param == nil {
		return nil, tserr.NilPtr()
	}
	// Retrieve the number of parameters from the error message definition
	l := len(m.Param)
	// If the number of parameters is less than or equal to 1,
	// return an error indicating that the expected number of parameters is greater than 1
	if l <= 1 {
		return nil, tserr.Higher(&tserr.HigherArgs{Var: "number of parameters", Actual: int64(l), LowerBound: 2})
	}
	// Create a new code snippet for the API function using the lpcode package.
	c := lpcode.NewCode().LineComment(m.Name + "Args holds the required arguments for the error function " + m.Name).TypeStruct(m.Name + "Args")
	for _, v := range m.Param {
		c.LineComment(v.Comment).VarSpec(&lpcode.VarSpecArgs{Ident: v.Name, Type: v.Type})
	}
	c.FuncEnd().LineComment(m.Comment).Func1(&lpcode.Func1Args{Name: m.Name, Var: "a", Type: " *" + m.Name + "Args", Return: "error"})
	c.If(&lpcode.IfArgs{ExprLeft: "a", ExprRight: "nil", Operator: "=="}).Return().Call("NilPtr").ParamEndln().BlockEnd()
	c.Return().Call("errorf").Addr().Ident("errmsg" + m.Name)
	for _, v := range m.Param {
		c.List().SelField(&lpcode.SelArgs{Val: "a", Sel: v.Name})
	}
	c.ParamEndln().FuncEnd()
	// Return the generated code snippet and nil to indicate that
	// the generation process completed without errors.
	return c, nil
}
