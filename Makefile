###############################################################################
# File        : Malefile
# Author      : hakkadaikon
###############################################################################
PROGRAM := unemp-tool 
BINROOT := ./bin
SRCROOT := ./src


.PHONY: build clean format run

build:
	cd $(SRCROOT) && go fmt && go build -o ../$(BINROOT)

clean:
	rm -f $(BINROOT)/*
format:
	cd $(SRCROOT) && go fmt
run:
	$(BINROOT)/$(PROGRAM)
