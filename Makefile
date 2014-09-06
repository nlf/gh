SRC = $(wildcard src/*.go)
BIN = $(patsubst src/%.go,bin/%,$(SRC))
DEPS = $(wildcard *.go)

bin/%: src/%.go $(DEPS)
	@mkdir -p bin
	go build -o $@ $<

all: $(BIN) | bin

clean:
	rm -rf bin

.phony: all clean
