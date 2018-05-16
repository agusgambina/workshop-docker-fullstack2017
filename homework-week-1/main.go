package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

// docker run --rm -v "$PWD":/usr/src/hello -w /usr/src/hello -e GOOS=darwin -e GOARCH=amd64 golang:latest go build -v

// docker run --rm -v "$PWD":/usr/src/hello -w /usr/src/hello golang:latest go build -v
