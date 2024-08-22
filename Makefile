.PHONY: clean gen-go

default: gen-go-wasip1

gen-go-wasip1: ; wit-bindgen-go generate --target wasip1 --world foo --out ./gen/go/wasip1 ./wit

gen-go-wasip2: ; wit-bindgen-go generate --world foo --out ./gen/go/wasip2 ./wit

build-wasip1: ; tinygo build -tags purego -o main.wasm -target=wasip1

build-wasip2: ; tinygo build -tags purego -o main.wasm -target=wasip2 -wit-package ./wit --wit-world foo 

wasip1: gen-go-wasip1 build-wasip1 ; go run ./cmd/wasmtime/main.go 

wasip2: gen-go-wasip2 build-wasip2 ; wasmtime run main.wasm 

wat: ; wasm-tools parse -t  main.wasm -o main.wat

clean: ; rm -rf ./gen/go/wasip1/ & rm -rf ./gen/go/wasip2/

