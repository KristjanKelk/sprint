package sprint

func IntVsFloat(i int, f float32) string {
	if float32(i) > f {
		return "integer"
	} else if float32(i) == f {
		return "same"
	} else {
		return "float"
	}
}
