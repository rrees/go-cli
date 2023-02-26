
.PHONY: run
run-dice: dice
	./bin/dice

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test: fmt
	go test ./...

dice:
	go build -o bin/dice cmd/dice/main.go

wc: test fmt
	go build -o bin/wc cmd/wc/main.go

dgstat: test fmt
	go build -o bin/dgstat cmd/dgstat/main.go

minibx: test fmt
	go build -o bin/minibx cmd/minibx/main.go

build: wc dice dgstat minibx