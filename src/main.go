package main

import (
	"fmt"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	cInstance, err := wasmC()
	if err != nil {
		panic(err)
	}
	defer cInstance.Close()

	// Gets exported functions from the WebAssembly instance.
	sumC := cInstance.Exports["sum"]
	multiplyC := cInstance.Exports["multiply"]

	rustInstance, err := wasmRust()
	if err != nil {
		panic(err)
	}
	defer rustInstance.Close()

	// Gets exported functions from the WebAssembly instance.
	sumRust := rustInstance.Exports["sum"]
	multiplyRust := rustInstance.Exports["multiply"]

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	resultSumC, _ := sumRust(5, 37)
	resultSumRust, _ := sumC(5, 37)

	resultMultiplyC, _ := multiplyC(2, 5)
	resultMultiplyRust, _ := multiplyRust(2, 5)

	fmt.Println("Sums are equal: ", resultSumC == resultSumRust)                 // 42!
	fmt.Println("Multiplies are equal: ", resultMultiplyC == resultMultiplyRust) // 10!
}

// Assembly files are generated using ´https://webassembly.studio/´

func wasmC() (wasm.Instance, error) {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("./lib/sample-c.wasm")

	// Instantiates the WebAssembly module.
	return wasm.NewInstance(bytes)
}

func wasmRust() (wasm.Instance, error) {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("./lib/sample-rust.wasm")

	// Instantiates the WebAssembly module.
	return wasm.NewInstance(bytes)
}
