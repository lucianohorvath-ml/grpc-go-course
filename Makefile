### Required tools
GOTOOLS_CHECK = go goverreport
export GO111MODULE=on

all: check_tools fmt test linter

### Tools & dependencies
check_tools:
	@# https://stackoverflow.com/a/25668869
	@echo "Found tools: $(foreach tool,$(GOTOOLS_CHECK),\
        $(if $(shell which $(tool)),$(tool),$(error "No $(tool) in PATH")))"

### Testing
test:
	@echo "==> Running tests..."
	go test -v ./... -covermode=atomic -coverpkg=./... -count=1 -race || (echo -e "****Test failed****"; exit 1)

### Formatting
fmt:
	@echo "==> Running format..."
	go fmt ./...

### Run binaries
run-server:
	@echo "==> Running gRPC server (greet)..."
	cd greet && go run greet_server/server.go

run-client:
	@echo "==> Running gRPC client (greet)..."
	cd greet && go run greet_client/client.go

run-server-calc:
	@echo "==> Running gRPC server (calculator)..."
	cd calculator && go run calculator_server/server.go

run-client-calc:
	@echo "==> Running gRPC client (calculator)..."
	cd calculator && go run calculator_client/client.go

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: all check_tools test test-cover fmt linter go-modules-init-12 go-modules-tidy-12 run-api
