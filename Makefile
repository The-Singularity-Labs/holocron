
GO_ROOT := $(shell go env GOROOT)
TINY_GO_ROOT := $(shell tinygo env TINYGOROOT)

all: run

dev:
	cd ui && yarn run dev --host 127.0.0.1 --port 41119

build: build-wasm build-linux build-windows build-osx

build-wasm:
	mkdir -p ui/src/assets/wasm
	cp $(TINY_GO_ROOT)/targets/wasm_exec.js ui/src/wasm_exec.js
	cd lib/wasm && tinygo build -o ../../ui/src/assets/wasm/golib.wasm -target wasm -no-debug  ./main.go
	cd ui/src/assets/wasm/ && wasm-opt -Oz golib.wasm -o golib.wasm

app:
	guark run

build-linux:
	guark build  --target linux --rm

build-windows:
	guark build  --target windows --rm

build-osx:
	guark build  --target darwin --rm