package simplefactory

import "fmt"

type Pen interface {
	Draw()
}

func NewPen(color string) Pen {
	switch color {
	case "blue":
		return bluePen{}
	default:
		return blackPen{}
	}
}

type blackPen struct {
}

func (p blackPen) Draw() {
	fmt.Println("draw a black line")
}

type bluePen struct {
}

func (p bluePen) Draw() {
	fmt.Println("draw a blue line")
}
