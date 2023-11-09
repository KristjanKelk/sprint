package sprint

// Point represents text at a specific point with X and Y coordinates.
type Point struct {
	X    float32
	Y    float32
	Text string
}

// MakePoint creates and returns a new Point struct with the specified values.
func MakePoint(x, y float32, text string) Point {
	return Point{X: x, Y: y, Text: text}
}
