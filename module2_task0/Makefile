build:
	go build -o awesome-api main.go
run:
	./awesome-api > ./awesome.log 2>&1 &
stop:
	pgrep awesome-api | xargs kill
clean:
	rm -f awesome*
test:
	curl http://0.0.0.0:9999
	curl http://0.0.0.0:9999/health
help:
	@echo "help: to know whats going on"
	@echo "build: to compile the app"
	@echo "run: to run the app"
	@echo "stop: to stop the app"
	@echo "clean: clean the build"
	@echo "test: to test the app"