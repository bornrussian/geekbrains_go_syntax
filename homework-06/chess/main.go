//
// Задача:
//
// * Напишите программу, генерирующую png-файл с рисунком шахматной доски.
//

package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

func makeChessBoard(w http.ResponseWriter, r *http.Request) {
	var boardSize int = 600 // pixels
	var cellSize int = boardSize / 8

	m := image.NewRGBA(image.Rect(0, 0, boardSize, boardSize))
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for y := 0; y < 9; y++ {
		for x := 0; x < 4; x++ {
			xShift := 0;
			if y%2 == 0 {
				xShift = cellSize;
			}
			x0 := x*2*cellSize + xShift
			y0 := y * cellSize

			draw.Draw(m, image.Rect(x0, y0, x0+cellSize, y0+cellSize), &image.Uniform{black}, image.ZP, draw.Src)
		}
	}

	var img image.Image = m
	writeImage(w, &img)
}

func writeImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func main() {
	http.HandleFunc("/", makeChessBoard)
	log.Println("Listening on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
