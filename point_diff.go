package sprint

type Point struct {
	X int
	Y int
}

func PointDiff(p1, p2 Point) []Point {
	if p1.X+p1.Y < p2.X+p2.Y {
		return []Point{p2}
	} else if p1.X+p1.Y > p2.X+p2.Y {
		return []Point{p1}
	}
	return []Point{p1, p2}
}
