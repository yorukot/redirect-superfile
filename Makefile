build:
	mkdir -p functions
	GOOS=linux GOARCH=amd64 go build -o functions/main main.go

clean:
	rm -rf functions

.PHONY: build clean