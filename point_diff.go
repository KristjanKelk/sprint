package sprint

type Point struct {
	X float32
	Y float32
}

func PointDiff(p1, p2 Point) Point {
	if p1.X+p1.Y < p2.X+p2.Y {
		return Point{X: p1.X, Y: p1.Y}
	} else if p1.X+p1.Y > p2.X+p2.Y {
		return Point{X: p2.X, Y: p2.Y}
	}
	return Point{X: p1.X, Y: p1.Y}
}
