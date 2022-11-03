package qrgen

import (
	"bytes"
	"image"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// CreateImage generates image from the given byte format of the image
func CreateImage(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}

// GenQRCode generates QR code for the given URL
func GenQRCode(url string) (barcode.Barcode, error) {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(qrCode, 200, 200)
}
