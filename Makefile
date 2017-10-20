SOURCE_DIR=.
BIN_NAME='lights'
PACKAGE_NAME=$(shell go list $(SOURCE_DIR))
PACKAGES=$(shell go list ./... | grep -v '^$(PACKAGE)/vendor/')
SOURCES=$(shell find $(SOURCE_DIR) -name '*.go')
BRANCH_NAME=$(shell git rev-parse --abbrev-ref HEAD)

# default target, compile the code
all: build

# get tools needed for running other targets
tools: go-lint
	go get -u github.com/kardianos/govendor
	go get -u github.com/kisielk/godepgraph
	go get -u github.com/whytheplatypus/godoc-templates

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

docs:
	@for pkg in $(PACKAGES); do \
		godoc -templates $$GOPATH/src/github.com/whytheplatypus/godoc-templates $$pkg | sed 's/\/target\///' | sed 's/(\\/(/' > $$GOPATH/src/$$pkg/README.md ; \
	done


notes:
	@for pkg in $(PACKAGES); do \
		godoc -notes ".*" -templates $$GOPATH/src/github.com/whytheplatypus/godoc-templates/notes-only cmd/$$pkg | sed 's/\/src\/github.com\/whytheplatypus\/$(BIN_NAME)/\/whytheplatypus\/$(BIN_NAME)\/$(BRANCH_NAME)/' | sed 's/(\\/(/' >> $$GOPATH/src/$(PACKAGE_NAME)/NOTES.md ; \
	done

# sed 's/\/src\/$(PACKAGE_NAME)\//\/whytheplatypus\/$(BIN_NAME)\/blob\/$(BRANCH_NAME)\//'
# Cleanup targets
# ----------------------------------------------------------------------------

clean:
	rm -f dep-graph.png

.PHONY: all tools go-lint build check-fmt format lint test dep_graph clean
