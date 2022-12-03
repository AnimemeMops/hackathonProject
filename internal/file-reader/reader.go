package filereader

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

type FileReader struct {
}

func NewFileReader() *FileReader {
	return &FileReader{}
}

type BmpImage struct {
	path string
	img  image.Image
}

func (bi *BmpImage) Path() string {
	return bi.path
}

func (bi *BmpImage) Img() image.Image {
	return bi.img
}

func newBmpImage(path string, img image.Image) *BmpImage {
	return &BmpImage{
		path: path,
		img:  img,
	}
}
func (fr *FileReader) ReadBmpImages(paths ...string) []*BmpImage {
	images := make([]*BmpImage, 0, len(paths))
	for i := range paths {
		img, err := fr.ReadBmpImage(paths[i])
		if err != nil {
			log.Print(fmt.Errorf("read bmp images error: %w", err))
			continue
		}
		images = append(images, img)
	}
	return images

}
func (fr *FileReader) ReadBmpImage(path string) (*BmpImage, error) {
	buff, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	image, err := bmp.Decode(bytes.NewReader(buff))
	return newBmpImage(path, image), err
}
