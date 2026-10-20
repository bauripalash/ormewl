PROJECT:=mewl
CMD_DIR:=cmd/$(PROJECT)
BUILD_DIR:=build

GOCMD:=go
GOFMT:=gofmt

FILERUN:=a.mwl

.PHONY: all
all: run

## build: build mewl interpreter
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	@$(GOCMD) build -o $(BUILD_DIR)/$(PROJECT) $(CMD_DIR)/main.go

## run: run the test file
.PHONY: run
run: build
	@./$(BUILD_DIR)/$(PROJECT) $(FILERUN) $(ARGS)


## repl: build and run mewl interpreter
.PHONY: repl
repl: build
	@./$(BUILD_DIR)/$(PROJECT) $(ARGS)

## fmt: format source code
.PHONY: fmt
fmt:
	@$(GOFMT) -s -w .

## clean: clean built artifacts
.PHONY: clean
clean:
	@$(GOCMD) clean
	@rm -rf $(BUILD_DIR)

## help: show this help
.PHONY: help
help:
	@echo "Mewl Dispatcher Help :"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/  /'
