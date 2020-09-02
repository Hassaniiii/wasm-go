package main

import (
	"fmt"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {}

// Assembly files are generated using ´https://webassembly.studio/´

func WasmC(x, y int) (int32, int32) {
	// Reads the WebAssembly module as bytes.
	bytes, err := wasm.ReadBytes("../../lib/sample-c.wasm")
	if err != nil {
		fmt.Println("Error in reading c wasm file")
		return -1, -1
	}

	// Instantiates the WebAssembly module.
	instance, err := wasm.NewInstance(bytes)
	if err != nil {
		return -1, -1
	}
	defer instance.Close()

	// Gets exported functions from the WebAssembly instance.
	sum := instance.Exports["sum"]
	multiply := instance.Exports["multiply"]

	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	resultSum, err := sum(x, y)
	if err != nil {
		fmt.Println("The error happens in suming with c wasm", err)
		return -1, -1
	}

	resultMultiply, err := multiply(x, y)
	if err != nil {
		fmt.Println("The error happens in multiplying with c wasm", err)
		return -1, -1
	}

	return resultSum.ToI32(), resultMultiply.ToI32()
}

func WasmRust(x, y int) (int32, int32) {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("../../lib/sample-rust.wasm")

	// Instantiates the WebAssembly module.
	instance, err := wasm.NewInstance(bytes)
	if err != nil {
		return -1, -1
	}
	defer instance.Close()

	// Gets exported functions from the WebAssembly instance.
	sum := instance.Exports["sum"]
	multiply := instance.Exports["multiply"]

	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	resultSum, err := sum(x, y)
	if err != nil {
		fmt.Println("The error happens in suming with rust wasm", err)
		return -1, -1
	}

	resultMultiply, err := multiply(x, y)
	if err != nil {
		fmt.Println("The error happens in multiplying with rust wasm", err)
		return -1, -1
	}

	return resultSum.ToI32(), resultMultiply.ToI32()
}
