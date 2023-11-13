package sprint

/*type Point struct {
	X    float32
	Y    float32
	Text string
}
func makePoint(x, y float32, text string) Point {
	return Point{X: x, Y: y, Text: text}
}
*/

func pointDiff(p1 Point, p2 Point) Point {
	if p1.X > p2.X || (p1.X == p2.X && p1.Y > p2.Y) {
		return p1
	}
	return p2
}
