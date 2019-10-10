package kernel

import "testing"

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
