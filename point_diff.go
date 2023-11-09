package sprint

type Point struct {
	X    float32
	Y    float32
	Text string
}

func MakePoint(x, y float32, text string) Point {
	return Point{X: x, Y: y, Text: text}
}

func PointDiff(p1, p2 Point, text string) Point {
	if p1.X > p2.X || (p1.X == p2.X && p1.Y > p2.Y) {
		return p1
	}
	return p2
}
