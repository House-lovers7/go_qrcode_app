package qrgen

import (
	"bytes"
	"image"
	"image/png"
)

// CreateImage generates image from the given byte format of the image
func CreateImage(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}
