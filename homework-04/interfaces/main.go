//
// Задача: Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.
//

package main

import (
	"fmt"
)

type Dude interface {
	Greeting() string
}

type Russian struct {
}

func (d Russian) Greeting() string {
	return "Привет"
}

type Englishman struct {
}

func (c Englishman) Greeting() string {
	return "Hello"
}

type American struct {
}

func (l American) Greeting() string {
	return "Hi"
}

type German struct {
}

func (j German) Greeting() string {
	return "Hallo"
}

type Chinese struct {
}

func (j Chinese) Greeting() string {
	return "你好"
}

func main() {
	dudes := []Dude{Russian{},Englishman{},American{},German{},Chinese{}}

	for _,dude:=range dudes {
		fmt.Println(dude.Greeting())
	}
}