.PHONY: run build clean

TARGET=bin/app

build:
	go build -o $(TARGET) ./cmd/app

run: build
	./bin/app

clean:
	rm -rf $(TARGET)
