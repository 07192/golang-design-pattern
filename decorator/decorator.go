package decorator

type Draw interface {
	Draw()
}

type Line struct{}

func (l Line) Draw() string {
	return "draw a line"
}

type ColorLine struct {
	Line
	color string
}

func (cl ColorLine) Draw() string {
	return cl.Line.Draw() + ", paint " + cl.color
}

func NewColorLine(color string) ColorLine {
	return ColorLine{Line{}, color}
}
