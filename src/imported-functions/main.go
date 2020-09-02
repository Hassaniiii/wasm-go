package main

// #include <stdlib.h>
//
// extern int32_t go_api_sum(void *context, int32_t x, int32_t y);
// extern int32_t go_api_multiply(void *context, int32_t x, int32_t y);
import "C"

import (
	"fmt"
	"unsafe"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {}

func InjectToWasm(x, y int, api_type string) int32 {
	var api interface{}
	var c_api unsafe.Pointer
	if api_type == "sum" {
		api = go_api_sum
		c_api = C.go_api_sum
	} else if api_type == "multiply" {
		api = go_api_multiply
		c_api = C.go_api_multiply
	}

	imports, _ := wasm.NewImports().AppendFunction("go_api", api, c_api)
	bytes, _ := wasm.ReadBytes("../../lib/sample-import.wasm")
	instance, _ := wasm.NewInstanceWithImports(bytes, imports)
	defer instance.Close()

	// Get exported function written in rust
	exported_function := instance.Exports["api"]

	// Call exported function
	resultSum, err := exported_function(x, y)
	if err != nil {
		fmt.Println("The error happens in suming with imported wasm", err)
		return -1
	}

	return resultSum.ToI32()
}

//export go_api_sum
func go_api_sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}

//export go_api_multiply
func go_api_multiply(context unsafe.Pointer, x int32, y int32) int32 {
	return x * y
}
