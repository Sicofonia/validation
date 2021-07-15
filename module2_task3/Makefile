build: lint ## Builds awesome-api
	go build -o awesome-api main.go
run: ## Executes awesome-api
	./awesome-api > ./awesome.log 2>&1 &
stop: ## Stops awesome-api
	pgrep awesome-api | xargs kill
clean: ## Cleans awesome-api
	rm -f awesome*
	rm -f coverage*
test: unit-tests integration-tests ## Run all tests
lint: ## Lint code
	golangci-lint run -v
unit-tests:
	go test -v -short -coverprofile=coverage-units.out
integration-tests: ## Integration tests
	go test -v -coverprofile=coverage-integrations.out
help: ## Want to know more?
	@echo "help: to know whats going on"
	@echo "build: to compile the app"
	@echo "clean: clean the build"
	@echo "run: to run the app"
	@echo "stop: to stop the app"
	@echo "test: Runs both integration and unit tests"
	@echo "lint: static analysis for the code"
	@echo "unit-tests: Unit tests"
	@echo "integration-tests: Integration tests"
	