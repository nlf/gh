SRC = $(wildcard src/*.go)
BIN = $(patsubst src/%.go,bin/%,$(SRC))

bin/%: src/%.go lib/github.go
	@mkdir -p bin
	go build -o $@ $<

all: $(BIN) | bin

.phony: all
