all: sqlc-gen-test sqlc-gen-test.wasm

install: all
	cp sqlc-gen-test ~/bin

sqlc-gen-test: main.go go.mod go.sum
	go build .

sqlc-gen-test.wasm: main.go go.mod go.sum
	GOOS=wasip1 GOARCH=wasm go build -o sqlc-gen-test.wasm main.go

sha256: sqlc-gen-test.wasm
	openssl sha256 sqlc-gen-test.wasm