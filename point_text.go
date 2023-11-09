package sprint

import "fmt"

type Point struct {
	X    float32
	Y    float32
	Text string
}

func makePoint(x, y float32, text string) Point {
	return Point{X: x, Y: y, Text: text}
}

func PointText(p Point) Point {
	formattedText := fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return Point{X: p.X, Y: p.Y, Text: formattedText}
}
