package kernel

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
