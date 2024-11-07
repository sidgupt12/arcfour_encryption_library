APP_NAME = arcfour
SRC = arcfour.go

all: build run

build:
	go build -o $(APP_NAME) $(SRC)
	

run:
	@read -p "Enter text to encrypt: " text; \
	./$(APP_NAME) "$$text"


clean:
	rm -f $(APP_NAME)

.PHONY: all build run clean