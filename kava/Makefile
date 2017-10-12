SOURCE_DIR=.
BIN_NAME='kava'
PACKAGE_NAME=$(shell go list $(SOURCE_DIR))
PACKAGES=$(shell go list ./... | grep -v '^$(PACKAGE)/vendor/')
SOURCES=$(shell find $(SOURCE_DIR) -name '*.go')

# default target, compile the code
all: build

# get tools needed for running other targets
tools: go-lint
	go get -u github.com/kardianos/govendor
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/kisielk/godepgraph

go-lint:
	$(eval GOLINT_INSTALLED := $(shell which golint))

	@if [ "$(GOLINT_INSTALLED)" = "" ] ; then \
		go get github.com/golang/lint/golint; \
	fi;

# friendly name for building the binary
build: $(BIN_NAME)


# Compile/build targets
# ----------------------------------------------------------------------------

# compile binary from sources
$(BIN_NAME): $(SOURCES)
	go-bindata config
	go build -x -o $(BIN_NAME) -ldflags "\
		-X $(PACKAGE_NAME)/version.BuildTime=$(shell date -u +%FT%T%z)\
		-X $(PACKAGE_NAME)/version.GitCommit=$(shell git rev-parse --short HEAD)\
		-X $(PACKAGE_NAME)/version.Version=$(shell git describe --abbrev=0 --tags 2> /dev/null || echo v0.0.1)"

# Test targets
# ----------------------------------------------------------------------------

test:
	go test -v -race $(PACKAGES)

# Lint/format targets
# ----------------------------------------------------------------------------

# check that all go files have been formated
check-fmt:
	$(eval NEEDS_FORMAT := $(shell gofmt -l . | grep -v '^vendor'))

	@if [ "$(NEEDS_FORMAT)" != "" ] ; then \
		echo "$(NEEDS_FORMAT)"; \
		exit 1; \
	fi;

# format all files, for those that don't have editors configured auto-format
format:
	go fmt $(PACKAGES)

# check for lint issues
lint: go-lint
	go vet $(PACKAGES)
	golint -set_exit_status $(PACKAGES)

# Friendly target to build dependancy graph
dep_graph: dep-graph.png

# Build the dependancy graph
dep-graph.png: $(SOURCES)
	godepgraph -s -p $(PACKAGE_NAME)/vendor $(PACKAGE_NAME) | dot -Tpng -o dep-graph.png

# Cleanup targets
# ----------------------------------------------------------------------------

clean:
	rm -f dep-graph.png

.PHONY: all tools go-lint build gen_proto gen_migrations gen check-fmt format lint test dep_graph clean
