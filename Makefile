.PHONY: build
build:
	cd ./cmd/go-sudoku-solver && go build

.PHONY: run
run: build
	./cmd/go-sudoku-solver/go-sudoku-solver

.PHONY: verify
verify:
	./scripts/check_gofmt.sh
	./scripts/check_golint.sh

.PHONY: test
test:
	cd ./cmd/go-sudoku-solver && go test
	go test ./pkg/gss/

.PHONY: test-unit
test-unit:
	./scripts/test_unit.sh
