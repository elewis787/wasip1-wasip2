# wasip1 vs wasip2 

[WebAssembly System Interface (WASI)](https://wasi.dev/) is a group of standard API specifications for software compiled to the W3C WebAssembly (Wasm) standard. WASI is designed to provide a secure standard interface for applications that can be compiled to Wasm from any language, and that may run anywhere—from browsers to clouds to embedded devices.

To date, WASI has seen two milestone releases known as 0.1 and 0.2. (Sometimes you will see these referred to as Preview 1 and Preview 2, or P1 and P2). We will reference these are wasip1 and wasip2 per the compiler targets defined. 

WASIP2 is typically linked with the [Component Model](https://component-model.bytecodealliance.org/). However, there are useful tools emergeing such as WIT or [WebAssembly Interface Type](https://component-model.bytecodealliance.org/design/wit.html). Using WIT to outlined the contracts of a Wasm Module or Component is highly desirable. 

This repo outlines how bindings generated via WIT and tools such as [wit-bindgen-go](https://github.com/ydnar/wasm-tools-go). 

All bindings in this repo are generated with wit-bindgen-go and require the latest(dev) branch of tinygo. 

## WASIP1 

`make wasip1` 

This produces and runs a core wasm module that uses the generated bindings for wasip1. 

A sample runtime is provided to execution the exported function. 

This can manually be executed by running `go run ./cmd/wasmtime/main.go` 

The golang bindings of wasmtime ( wasmtime-go) do not support wasip2 components. 

### Observations 

Tinygo optimizes out the wasip2 imports for wasi components such as poll and streams. This is useful for hosts that do not intend on implement wasip2 or it is not supported. 

## WASIP2 

`make wasip2` 

This produces and runs a wasm component that is generated for wasip2. 

This uses wasmtime cli to execute the component. 

### Observations 

Wasmtime cli is the only runtime this has been tested with. In theory the rust wasmtime library could be used to execute the component but it has not been tested. 