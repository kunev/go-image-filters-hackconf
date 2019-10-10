package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	fmt.Printf("Attempting to read image from %s\n", filePath)
	imageFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	imageData, format, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read a %s image \n", format)
	fmt.Printf("The size of the image is %s\n", imageData.Bounds().Size())

	return
}
