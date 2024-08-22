package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/bytecodealliance/wasmtime-go/v23"
)

func main() {
	config := wasmtime.NewConfig()
	engine := wasmtime.NewEngineWithConfig(config)

	//read the wasmfile
	wasm, err := os.ReadFile("main.wasm")
	if err != nil {
		log.Fatal(err)
	}

	module, err := wasmtime.NewModule(engine, wasm)
	if err != nil {
		log.Fatal(err)
	}

	// Create a linker with WASI functions defined within it
	linker := wasmtime.NewLinker(engine)
	err = linker.DefineWasi()
	if err != nil {
		log.Fatal(err)
	}

	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritArgv()
	wasiConfig.InheritStdout()
	wasiConfig.InheritStderr()
	store := wasmtime.NewStore(engine)
	store.SetWasi(wasiConfig)

	instance, err := linker.Instantiate(store, module)
	if err != nil {
		log.Fatal(err)
	}

	// Run the function
	nom := instance.GetFunc(store, "elewis787:example/bar@0.0.1#baz")
	res, err := nom.Call(store, 0, 0)

	if err != nil {
		log.Fatal(err)
	}

	mem := instance.GetExport(store, "memory")
	data := mem.Memory().UnsafeData(store)

	ret := res.(int32)
	start := int32(binary.LittleEndian.Uint32(data[ret : ret+4]))
	len := int32(binary.LittleEndian.Uint32(data[ret+4 : ret+8]))

	fmt.Println("Output:\n", string(data[start:start+len]))

}
