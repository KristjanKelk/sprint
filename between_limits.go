package sprint

func BetweenLimits(from rune, to rune) string {
	if from < to {
		runes := ""
		for i := from + 1; i < to; i++ {
			runes += string(i)
		}
		return runes
	} else if from > to {
		runes := ""
		for i := to + 1; i < from; i++ {
			runes += string(i)
		}
		return runes
	}
	return ""
}
