NO_COLOR=\033[0m
OK_COLOR=\033[0;32m
TRACE_HOME=/tmp/test_trace_home

DEBUG?=0
ifeq ($(DEBUG), 1)
	VERBOSE="-v"
endif

all: test

format:
	@echo "$(OK_COLOR)==> Formatting the code $(NO_COLOR)"
	@gofmt -s -w *.go
	@goimports -w *.go

install:
	@echo "$(OK_COLOR)==> Installing test binaries $(NO_COLOR)"
	@`which go` install -v .

test:
	@echo "$(OK_COLOR)==> Preparing test environment $(NO_COLOR)"
	@echo "Cleaning $(TRACE_HOME) directory"
	@rm -rf $(TRACE_HOME)

	@echo "$(OK_COLOR)==> Building packages $(NO_COLOR)"
	@`which go` build -v ./...

	@echo "$(OK_COLOR)==> Testing packages $(NO_COLOR)"
	@$(HOME)/gopath/bin/goveralls -repotoken tHHcT2LfapnvCwCZ0ao3W883yBs4XzIS4

doc:
	@`which godoc` github.com/dizitart/trace | less

vet:
	@echo "$(OK_COLOR)==> Running go vet $(NO_COLOR)"
	@`which go` vet .

lint:
	@echo "$(OK_COLOR)==> Running golint $(NO_COLOR)"
	@`which golint` .

ctags:
	@ctags -R --languages=c,go

.PHONY: all install format test doc vet lint ctags
