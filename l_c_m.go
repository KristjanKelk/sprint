package sprint

func LCM(a, b int) int {
	larger := a
	smaller := b
	if larger < b {
		larger = b
		smaller = a
	}

	if smaller == 0 {
		return 0
	}

	for i := larger; ; i += larger {
		if i%a == 0 && i%b == 0 {
			return i
		}
	}
}
