// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build wasip1

// Package bar represents the exported interface "elewis787:example/bar@0.0.1".
package bar

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

//go:wasmexport elewis787:example/bar@0.0.1#baz
//export elewis787:example/bar@0.0.1#baz
func wasmexport_Baz(key0 *uint8, key1 uint32) (result *string) {
	key := cm.LiftString[string]((*uint8)(key0), (uint32)(key1))
	result_ := Exports.Baz(key)
	result = &result_
	return
}
