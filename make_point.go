package sprint

// Point represents text at a specific point with X and Y coordinates.
type Point1 struct {
	X    float32
	Y    float32
	Text string
}

// MakePoint creates and returns a new Point struct with the specified values.
func MakePoint(x, y float32, text string) Point1 {
	return Point1{X: x, Y: y, Text: text}
}
