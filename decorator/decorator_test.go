package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	painter := NewColorLine("red")
	picture := painter.Line.Draw()
	picture2 := painter.Draw()
	fmt.Println(picture)
	fmt.Println(picture2)
}
