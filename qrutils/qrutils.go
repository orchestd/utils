package qrutils

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)

// GetQr will return source and error
// The source can be used in <img> tag without any modifications, like: <img src="${source}" alt="Barcode">
func GetQr(qrValue string, size int) (string, error) {
	q, err := qrcode.New(qrValue, qrcode.Medium)
	if err != nil {
		return "", err
	}

	q.DisableBorder = true

	png, err := q.PNG(size)
	if err != nil {
		return "", err
	}

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(png), nil
}
