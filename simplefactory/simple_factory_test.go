package simplefactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	blackPen := NewPen("black")
	bluePen := NewPen("blue")
	blackPen.Draw()
	bluePen.Draw()
}
