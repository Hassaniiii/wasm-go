package main

import "testing"

func TestWasmC(t *testing.T) {
	c_results_sum, c_results_multiply := WasmC(5, 10)
	if c_results_sum == -1 || c_results_multiply == -1 {
		t.Error("The C wasm loaded with error")
	}

	if c_results_sum != 15 || c_results_multiply != 50 {
		t.Error("The C wasm calculated wrongly")
	}
}

func TestWasmRust(t *testing.T) {
	rust_results_sum, rust_results_multiply := WasmRust(3, 8)
	if rust_results_sum == -1 || rust_results_multiply == -1 {
		t.Error("The Rust wasm loaded with error")
	}

	if rust_results_sum != 11 || rust_results_multiply != 24 {
		t.Error("The Rust wasm calculated wrongly")
	}
}

func TestWasmCompare(t *testing.T) {
	c_results_sum, c_results_multiply := WasmC(4, 3)
	rust_results_sum, rust_results_multiply := WasmRust(4, 3)

	if c_results_sum != rust_results_sum || c_results_multiply != rust_results_multiply {
		t.Error("Rust and C wasms return different resutls!")
	}
}
