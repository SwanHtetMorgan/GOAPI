SERVER_PORT=8080
BLUE=\033[1;34m
NC=\033[0m  # No Color

export DB_USERNAME=root
export DB_PASSWORD=Swanhtet
export DB_HOST=localhost
export DB_PORT=3306
export DB_NAME=Banking

build:
	@echo "$(BLUE)Building the application .....$(NC)"
	@go build -o ./bin/API *.go

run: build
	@echo "$(BLUE)Running the application on the port: http://localhost:$(SERVER_PORT)$(NC)"
	./bin/API

clean:
	@echo "$(BLUE)Cleaning up...$(NC)"
	rm -rf ./bin/API
