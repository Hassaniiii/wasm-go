package main

import "testing"

func TestWasmImportedAPI_1(t *testing.T) {
	results_sum := InjectToWasm(3, 5, "sum")
	if results_sum == -1 {
		t.Error("The imported wasm loaded with error")
	}

	if results_sum != 8 {
		t.Error("The imported wasm calculated wrongly")
	}
}

func TestWasmImportedAPI_2(t *testing.T) {
	results_sum := InjectToWasm(3, 5, "multiply")
	if results_sum == -1 {
		t.Error("The imported wasm loaded with error")
	}

	if results_sum != 15 {
		t.Error("The imported wasm calculated wrongly")
	}
}
