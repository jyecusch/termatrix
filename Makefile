ifeq ($(OS), Windows_NT)
	EXECUTABLE_EXT = .exe
else
	EXECUTABLE_EXT =
endif

.PHONY: build run

build:
	go build -o bin/matrix$(EXECUTABLE_EXT) main.go

run:
	go run main.go
