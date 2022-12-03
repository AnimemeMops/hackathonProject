package vectorcalculator

import (
	"image"
)

func NewManager() *VC {
	return &VC{}
}

func (vc *VC) RGBAVector(img image.Image) []float64 {
	bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y
	rgba := make([]float64, 0, width*height*4)

    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rgba = append(rgba, float64(r),float64(g),float64(b),float64(a))
        }
    }
	
	return rgba
}