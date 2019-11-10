package photoshop

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func DrawRandomLinesToPNGFile(imgWidth, imgHeight int, linesCount int, pngFileName string) error {
	if imgWidth<1 || imgHeight<1 {
		return errors.New("imgWidth and imgHeight must be >= 1")
	}
	if linesCount<1 {
		return errors.New("linesCount must be >= 1")
	}

	// Создаём холст нужного размера
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Заливаем всё черным цветом
	black := color.RGBA{0,0,0,255}
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	// Рисуем цветные линии в количестве linesCount штук
	for i:=0; i<linesCount; i++ {
		rand.Seed(time.Now().UnixNano())
		randColor :=color.RGBA{128+uint8(rand.Intn(127)),128+uint8(rand.Intn(127)),128+uint8(rand.Intn(127)),255}
		randX :=rand.Intn(imgWidth)
		randY :=rand.Intn(imgHeight)
		randDirection :=rand.Intn(2)
		switch randDirection {
		case 0: // horizontal direction
				rndLength:=rand.Intn(imgWidth - randX)
				for x := randX; x < randX+rndLength; x++ {
					img.Set(x, randY, randColor)
				}
		default:
			rndLength:=rand.Intn(imgHeight - randY)
			for y := randY; y < randY+rndLength; y++ {
				img.Set(randX, y, randColor)
			}
		}
	}

	file, err := os.Create(pngFileName)
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, img)
	return nil
}
