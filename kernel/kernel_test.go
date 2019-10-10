package kernel

import (
	"fmt"
	"image"
	"os"
	"testing"

	_ "image/jpeg"
)

func TestNewKernel(t *testing.T) {
	kernel := New([][]float32{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
	if kernel.Width != 3 {
		t.Fatal("Width is wrong")
	}
	if kernel.Height != 3 {
		t.Fatal("Height is wrong")
	}
}

func TestKernel_Apply(t *testing.T) {
	kernel := New([][]float32{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
	file, err := os.Open("../input.jpg")
	image, _, err := image.Decode(file)
	progressChannel := make(chan int)
	go func() {
		for progress := range progressChannel {
			if progress >= 100 || progress < 0 {
				t.Error(fmt.Sprintf("incorrect progress percent reported: %d", progress))
			}
		}
	}()
	result, err := kernel.Apply(image, progressChannel)
	close(progressChannel)
	if err != nil {
		t.Fatal("kernel.Apply returner an error: ", err)
	}
	if result.Bounds() != image.Bounds() {
		t.Fatal("result of kernel.Apply has different size")
	}
}
