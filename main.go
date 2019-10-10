package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[1]
	fmt.Printf("Attempting to read image from %s\n", filePath)
	return
}
