package vectorcalculator

import "image"

type Img interface {
	Path() string
	Img() image.Image
}

type IV struct {
	Path   string
	Vector []float64
}
