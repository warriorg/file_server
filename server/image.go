package server

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

const JpegMaximumQuality = 100

// EncodeImage encodes the image with the given format
func EncodeImage(targetImage io.Writer, imageData image.Image, imageFormat string) error {
	switch imageFormat {
	case "jpeg":
		jpeg.Encode(targetImage, imageData, &jpeg.Options{JpegMaximumQuality})
	case "png":
		png.Encode(targetImage, imageData)
	case "gif":
		gif.Encode(targetImage, imageData, &gif.Options{256, nil, nil})
	default:
		return fmt.Errorf("invalid imageFormat given")
	}

	return nil
}
