all: 
	go clean
	go build -ldflags "-s -w"