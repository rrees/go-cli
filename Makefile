
.PHONY: run
run-dice: dice
	./bin/dice

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

dice:
	go build -o bin/dice cmd/dice/main.go

wc: test fmt
	go build -o bin/wc cmd/wc/main.go


build: dice