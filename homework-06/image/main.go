//
// Задача:
//
// Дополните пример из раздела «Пакет img» изображением горизонтальных и вертикальных
// линий.
//

package main

import (
	"fmt"
	"geekbrains_go_syntax/homework-06/image/photoshop"
	"image/color"
)

func main() {
	err := photoshop.DrawRectangleToPNGFile(100, 50, color.RGBA{0,255,0,255}, "_greenRect100x50.png")
	if err != nil {
		fmt.Println(err)
	}

	err = photoshop.DrawRandomLinesToPNGFile(400,300,100,"_randomLines400x300.png")
	if err != nil {
		fmt.Println(err)
	}
}
