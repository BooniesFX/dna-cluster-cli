GOFMT=gofmt
GC=go build
VERSION := $(shell git describe --abbrev=4 --dirty --always --tags)
Minversion := $(shell date)
BUILD_CLI_PARA = -ldflags "-X dna-cluster/common/version.Version=$(VERSION)" #-race

all:
	$(GC)  $(BUILD_CLI_PARA) -o dna-cluster-cli main.go

format:
	$(GOFMT) -w main.go

clean:
	rm -rf *.8 *.o *.out *.6