build: lint ## Builds awesome-api
	go build -o awesome-api main.go
run: ## Executes awesome-api
	./awesome-api > ./awesome.log 2>&1 &
stop: ## Stops awesome-api
	pgrep awesome-api | xargs kill
clean: ## Cleans awesome-api
	rm -f awesome*
test: ## Tests awesome-api
	curl http://0.0.0.0:9999
	curl http://0.0.0.0:9999/health
lint: ## Lint code
	golangci-lint run -v
unit-tests:
	go test -v -short -coverprofile=coverage-units.out
help: ## Want to know more?
	@echo "help: to know whats going on"
	@echo "build: to compile the app"
	@echo "run: to run the app"
	@echo "stop: to stop the app"
	@echo "clean: clean the build"
	@echo "test: to test the app"
	@echo "lint: static analysis for the code"
	