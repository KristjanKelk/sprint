package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	const pi = 3.14

	// Calculate diameter, area, and perimeter based on the given radius
	diameter := 2 * r
	area := pi * r * r
	perimeter := 2 * pi * r

	// Create and return a new Circle instance with the calculated values
	return Circle{
		Radius:    r,
		Diameter:  diameter,
		Area:      area,
		Perimeter: perimeter,
	}
}
