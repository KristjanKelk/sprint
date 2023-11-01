package sprint

func BetweenLimits(from rune, to rune) string {
	if from < to {
		runes := ""
		for i := from; i < to; i++ {
			runes += string(i)
		}
		return runes
	} else if from > to {
		runes := ""
		for i := to; i < from; i++ {
			runes += string(i)
		}
		return runes
	}
	return ""
}
