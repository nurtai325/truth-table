src = $(shell find . -name '*.go')

truth-table: $(src)
	go build -o truth-table cmd/truth-table/main.go
test: cmd/test/main.go $(src)
	go build -o test cmd/test/main.go
	./test
push:
	make
	make test
	go test ./...
	make clean
	git push
clean:
	rm truth-table test
