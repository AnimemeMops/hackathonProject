package vectorcalculator

import (
	"image"
)

const (
	flot64_255 = float64(255)
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
			rgba = append(rgba, 
				float64(r)/flot64_255,
				float64(g)/flot64_255,
				float64(b)/flot64_255,
				float64(a)/flot64_255,
			)
        }
    }
	
	return rgba
}