SERVER_PORT=8080

build:
	@echo "Building the application ....."
	@go build -o ./bin/API main.go

run: build
	@echo  "Running the application on the port: http://localhost:$(SERVER_PORT)"
	./bin/API

clean:
	@echo  "Cleaning the Binary file"
	rm -rf bin/API
