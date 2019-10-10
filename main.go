package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/kunev/go-image-filters-hackconf/kernel"
)

func loadImage(filePath string) (image.Image, string, error) {
	imageFile, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	imageData, format, err := image.Decode(imageFile)
	if err != nil {
		return nil, "", err
	}
	return imageData, format, nil
}

func writeImage(imageData image.Image, format string) error {
	writer, err := os.Create(fmt.Sprintf("output.%s", format))
	if err != nil {
		log.Fatal(err)
	}

	switch format {
	case "jpeg":
		return jpeg.Encode(writer, imageData, nil)
	case "png":
		return png.Encode(writer, imageData)
	default:
		return errors.New("Unknown format")
	}
}

func main() {
	filePath := os.Args[1]
	fmt.Printf("Attempting to read image from %s\n", filePath)
	imageData, format, err := loadImage(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read a %s image \n", format)
	fmt.Printf("The size of the image is %s\n", imageData.Bounds().Size())

	k := kernel.New([][]float32{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
	resultImage, err := k.Apply(imageData)

	if err != nil {
		log.Fatal(err)
	}

	if err := writeImage(resultImage, format); err != nil {
		log.Fatal(err)
	}

	return
}
