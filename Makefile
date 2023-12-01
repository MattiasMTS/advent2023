year := $(shell date +'%Y')
day := $(shell date +'%-d')
part := 1

.PHONY: test
test:
	@echo "Testing year=$(year) day=$(day) part=$(part)"
	@if [ $(day) -lt 10 ]; then \
		go test ./day0$(day) -v -run=solvePart$(part) ; \
	else \
		go test ./day$(day) -v -run=solvePart$(part) ; \
	fi

.PHONY: gen
gen:
	@echo "Generating year=$(year) day=$(day)"
	@go run ./cli/main.go input -d $(day) -y $(year)
	@go run ./templates/generate_aoc.go day=$(day)

.PHONY: submit
submit:
	@echo "Submitting year=$(year) day=$(day) part=$(part)"
	@if [ $(day) -lt 10 ]; then \
		_SUBMIT=1 go test ./day0$(day) -v -run=solvePart$(part) | go run ./cli/main.go submit -d $(day) -y $(year) -p $(part) ; \
	else \
		_SUBMIT=1 go test ./day$(day) -v -run=solvePart$(part) | go run ./cli/main.go submit -d $(day) -y $(year) -p $(part) ; \
	fi
