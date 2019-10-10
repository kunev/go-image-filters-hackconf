package kernel

import "image"

// Kernel describes an image kernel
type Kernel struct {
	Width        int
	Height       int
	Coefficients [][]float32
}

// New returns a Kernel wrapping the given coefficients matrix
func New(coefficients [][]float32) Kernel {
	result := Kernel{}
	result.Height = len(coefficients)
	result.Width = len(coefficients[0])
	result.Coefficients = coefficients

	return result
}

// Apply applies a kernel to an image returning the resulting image
func (k Kernel) Apply(img image.Image) (image.Image, error) {
	imageBounds := img.Bounds()
	result := image.NewRGBA(imageBounds)

	for x := imageBounds.Min.X; x < imageBounds.Max.X; x++ {
		for y := imageBounds.Min.Y; y < imageBounds.Max.Y; y++ {
			result.Set(x, y, img.At(x, y))
		}
	}

	return result, nil
}
