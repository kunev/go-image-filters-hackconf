package kernel

import (
	"image"
	"image/color"
)

// Kernel describes an image kernel
type Kernel struct {
	Width        int
	Height       int
	Coefficients [][]float32
}

type neighbour struct {
	xOffset int
	yOffset int
	clr     color.Color
}

// New returns a Kernel wrapping the given coefficients matrix
func New(coefficients [][]float32) Kernel {
	result := Kernel{}
	result.Height = len(coefficients)
	result.Width = len(coefficients[0])
	result.Coefficients = coefficients

	return result
}

// NewBlur returns a blur Kernel with the specified size
func NewBlur(size int) Kernel {
	coefficients := [][]float32{}

	for x := 0; x < size; x++ {
		row := []float32{}
		for y := 0; y < size; y++ {
			row = append(row, 1/float32(size*size))
		}
		coefficients = append(coefficients, row)
	}

	return New(coefficients)
}

func (k Kernel) getNeighbourhood(x, y int, img image.Image) []neighbour {
	bounds := img.Bounds()
	neighbourhood := []neighbour{}
	for i := -k.Width / 2; i <= k.Width/2; i++ {
		if x+i < bounds.Min.X || x+i > bounds.Max.X {
			continue
		}
		for j := -k.Height / 2; j <= k.Height/2; j++ {
			if y+j < bounds.Min.Y || y+j > bounds.Max.Y {
				continue
			}
			neighbourhood = append(neighbourhood, neighbour{
				xOffset: i,
				yOffset: j,
				clr:     img.At(x+i, y+j),
			})
		}
	}
	return neighbourhood
}

func (k Kernel) pixelValueFromNeighbourhood(neighbourhood []neighbour) color.Color {
	result := color.RGBA64{}
	for _, n := range neighbourhood {
		coef := k.Coefficients[n.xOffset+k.Width/2][n.yOffset+k.Height/2]
		r, g, b, a := n.clr.RGBA()
		result.R += uint16(float32(r) * coef)
		result.G += uint16(float32(g) * coef)
		result.B += uint16(float32(b) * coef)
		result.A = uint16(a)
	}
	return result
}

// Apply applies a kernel to an image returning the resulting image
func (k Kernel) Apply(img image.Image, progressChannel chan<- int) (image.Image, error) {
	imageBounds := img.Bounds()
	result := image.NewRGBA(imageBounds)
	totalPixels := imageBounds.Max.X * imageBounds.Max.Y
	currentlyProcessedPixels := 0

	for x := imageBounds.Min.X; x < imageBounds.Max.X; x++ {
		for y := imageBounds.Min.Y; y < imageBounds.Max.Y; y++ {
			neighbourhood := k.getNeighbourhood(x, y, img)
			result.Set(x, y, k.pixelValueFromNeighbourhood(neighbourhood))
			currentlyProcessedPixels++
		}

		progressChannel <- 100 * currentlyProcessedPixels / totalPixels
	}

	return result, nil
}
