###############################################################################
# File        : Malefile
# Author      : hakkadaikon
###############################################################################
PROGRAM := unemp-tool 
BINROOT := ./bin
SRCROOT := ./src


.PHONY: build clean format run test

build:
	cd $(SRCROOT) && go fmt && go build -o ../$(BINROOT) && rm -f $(PROGRAM)

clean:
	rm -f $(BINROOT)/*
format:
	cd $(SRCROOT) && go fmt
run:
	$(BINROOT)/$(PROGRAM)
test:
	cd $(SRCROOT) && go fmt && go build && go test
