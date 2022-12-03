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
func (fr *FileReader) ReadBmpImages(paths ...string) []image.Image {
	images := make([]image.Image, 0, len(paths))
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
func (fr *FileReader) ReadBmpImage(path string) (image.Image, error) {
	buff, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	image, err := bmp.Decode(bytes.NewReader(buff))
	return image, err
}
