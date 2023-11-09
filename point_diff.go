package sprint

type Point struct {
	P1 float32
	P2 float32
}

func PointDiff(p1, p2 Point) Point {
	if p1.P1+p1.P2 < p2.P1+p2.P2 {
		return Point{P1: p1.P1, P2: p1.P2}
	} else if p1.P1+p1.P2 > p2.P1+p2.P2 {
		return Point{P1: p2.P1, P2: p2.P2}
	}
	return Point{P1: p1.P1, P2: p1.P2}
}
