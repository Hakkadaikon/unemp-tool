###############################################################################
# File        : Malefile
# Author      : hakkadaikon
###############################################################################
PROGRAM := unemp-tool 
BINROOT := ./bin
SRCROOT := ./src


.PHONY: build clean format

build:
	cd $(SRCROOT) && go build -o ../$(BINROOT)

clean:
	rm -f $(BINROOT)/*
format:
	cd $(SRCROOT) && go fmt
