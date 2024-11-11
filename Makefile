###############################################################################
# File        : Malefile
# Author      : hakkadaikon
###############################################################################
PROGRAM := unemp-tool 
BINROOT := ./bin
SRCROOT := ./src


.PHONY: install build clean format run test

install:
	cd $(SRCROOT) && \
	go mod tidy
build:
	cd $(SRCROOT) && \
	go fmt && \
	go build -o ../$(BINROOT) && \
	rm -f $(PROGRAM)

clean:
	rm -f $(BINROOT)/*
format:
	cd $(SRCROOT) && go fmt
run:
	$(BINROOT)/$(PROGRAM)
test:
	cd $(SRCROOT) && \
	go fmt && \
	go build && \
	rm -f $(PROGRAM) && \
	go test && \
	go test -v ./unemp
