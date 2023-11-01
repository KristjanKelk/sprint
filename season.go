package sprint

func Season(month string) string {
	switch month {
	case "jan", "feb", "dec":
		return "Winter"
	case "mar", "apr", "may":
		return "Spring"
	case "jun", "jul", "aug":
		return "Summer"
	case "sep", "oct", "nov":
		return "Autumn"
	default:
		return "unkown"
	}
}
