GO=go 
BINFOLDER=bin
BINARY=binary-tree

.PHONY:
build: clean
	$(GO) build -o $(BINFOLDER)/$(BINARY)

.PHONY:
test:
	$(GO) test -v -cover ./...

.PHONY:
clean:
	rm -rf $(BINFOLDER)