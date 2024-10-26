src = $(shell find . -name '*.go')

truth-table: $(src)
	go build -o truth-table cmd/truth-table/main.go
clean:
	rm truth-table
