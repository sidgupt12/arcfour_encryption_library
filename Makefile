APP_NAME = arcfour
SRC = arcfour.go

all: build run

build:
	go build -o $(APP_NAME) $(SRC)

run:
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

.PHONY: all build run clean