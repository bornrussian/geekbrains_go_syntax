//
// Задача:
//
// *Написать функцию,
//  которая будет получать позицию коня на шахматной доске,
//  а возвращать массив из точек, в которые конь сможет сделать ход.
//
//	a. Точку следует обозначить как структуру, содержащую ​ x ​ и ​ y ​ типа ​ int
//	b. Получение значений и их запись в точку должны происходить только с помощью
//		отдельных методов. В них надо проводить проверку на то, что такая точка может
//		существовать на шахматной доске.
//

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Point struct {
	X,Y int
}

// Получение значения из точки с проверкой возможности ее существования на карте
func (p Point) Get() (curX, curY int) {
	if p.X<0 || p.X>7 || p.Y<0 || p.Y>7 {
		panic ("Точка находится за пределами шахматной доски.")
	}
	return p.X,p.Y
}

// Установка точки на карту с проверкой возможности
func (p *Point) Put(newX, newY int) (bool) {
	if newX>=0 && newX<=7 && newY>=0 && newY<=7 {
		p.X, p.Y = newX, newY
		return true
	} else {
		return false
	}
}

type ChessKnight struct {
}

// Та самая функция,
////  которая будет получать позицию коня на шахматной доске,
////  а возвращать массив из точек, в которые конь сможет сделать ход.
func (knight ChessKnight) ListPossibleTurnsFrom (position Point) (to []Point) {
	nowX,nowY:=position.Get()
	var turn Point
	if turn.Put(nowX-1,nowY-2) {
		to=append(to,turn)
	}
	if turn.Put(nowX-2,nowY-1) {
		to=append(to,turn)
	}
	if turn.Put(nowX+1,nowY+2) {
		to=append(to,turn)
	}
	if turn.Put(nowX+2,nowY+1) {
		to=append(to,turn)
	}
	if turn.Put(nowX-1,nowY+2) {
		to=append(to,turn)
	}
	if turn.Put(nowX-2,nowY+1) {
		to=append(to,turn)
	}
	if turn.Put(nowX+1,nowY-2) {
		to=append(to,turn)
	}
	if turn.Put(nowX+2,nowY-1) {
		to=append(to,turn)
	}
	return
}

func PrintlnJSON (v interface{}) {
	js,err:=json.Marshal(v)
	if err==nil {
		fmt.Println(string(js))
	} else {
		fmt.Println("Ошибка преобразования в JSON:", err)
	}
}

func PrintChessBoard (startPos Point, turns []Point) {
	var isPrinted bool
	for y:=0;y<8;y++ {
		for x:=0;x<8;x++ {
			isPrinted = false
			if x==startPos.X && y==startPos.Y {
				fmt.Printf("  Г")
				isPrinted = true
			}
			for _,turn:=range turns {
				if x==turn.X && y==turn.Y {
					fmt.Printf("  x")
					isPrinted = true
				}
			}
			if !isPrinted {
				fmt.Printf("  .")
			}
		}
		fmt.Printf("\n")
	}

}

func main () {
    // Обработчик паники
    defer func() {
        if r:= recover(); r!=nil {
            fmt.Println ("Паника:",r)
        }
    } ()

    // Генерируем коня
    var knight ChessKnight

    // Придумываем коню случайное место на доске:
	fmt.Printf("Если позиция коня на шахматной доске будет такой:\n")
	rand.Seed(time.Now().Unix())
	startPos:=Point{rand.Intn(8),rand.Intn(8)}
    PrintlnJSON(startPos)

	// Получаем возможные ходы:
	fmt.Printf("\nТо он имеет возможность ходить в следующие клетки:\n")
    turns:=knight.ListPossibleTurnsFrom(startPos)
	PrintlnJSON(turns)

    // Отображаем эти точки в виде псевдографики
	fmt.Printf("\nВизуально такая шахматная доска будет выглядеть так:\n")
	PrintChessBoard(startPos,turns)
}
