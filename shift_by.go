package sprint

func ShiftBy(r rune, step int) rune {
	// if rune is a, set index to 0
	index := int(r - 'a')

	// shifting index
	shiftedIndex := (index + step) % 26

	//shifting rune

	shiftedRune := rune('a' + shiftedIndex)

	return shiftedRune
}
