BIN_DIR = bin
OBSERVER = observer
ARTIFACTS = .artifacts
CONFIG = config

.PHONY: build
build: $(BIN_DIR) $(OBSERVER)

.PHONY: artifacts
artifacts: $(ARTIFACTS) $(CONFIG)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

.PHONY: $(OBSERVER)
$(OBSERVER):
	go build -o $(BIN_DIR)/$(OBSERVER) cmd/observer/*.go

$(ARTIFACTS):
	mkdir -p $(ARTIFACTS)

.PHONY: $(CONFIG)
$(CONFIG):
	go run configuration/gen/gen.go
	mv ./observer.yaml $(ARTIFACTS)/observer.yaml

.PHONY: ensure
ensure:
	dep ensure -v

.PHONY: test
test:
	go test -json -v -count 10 -timeout 20m --coverprofile=converage.txt --covermode=atomic ./... | tee ci_test_with_coverage.json

.PHONY: all
all: ensure build artifacts
