package photoshop

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func DrawRectangleToPNGFile(imgWidth, imgHeight int, imgColor color.RGBA, pngFileName string) error {
	if imgWidth<1 || imgHeight<1 {
		return errors.New("imgWidth and imgHeight must be >= 1")
	}
	rectImg := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{imgColor}, image.ZP, draw.Src)
	file, err := os.Create(pngFileName)
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, rectImg)
	return nil
}
