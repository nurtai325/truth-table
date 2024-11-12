src = $(shell find . -name '*.go')

truth-table: $(src)
	go mod tidy
	go build -o truth-table cmd/truth-table/main.go
test: cmd/test/main.go $(src)
	go mod tidy
	go build -o test cmd/test/main.go
clean:
	rm truth-table test
