# Переменные для компилятора
RUN = go run
BUILD = go build

# Переменные для линтера
GOLANGCI = golangci-lint

# Переменные с именами и путями приложения
PROJECT_NAME = 	rest-placeholder
CMD = cmd/$(PROJECT_NAME).go
BUILD_PATH = bin/$(PROJECT_NAME)

run:
	$(RUN) $(CMD)

build:
	$(BUILD) -o $(BUILD_PATH) $(CMD)

lint:
	$(GOLANGCI) run