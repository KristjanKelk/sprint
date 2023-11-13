package sprint

func CombN2(n int) []string {
	var combs []string
	genComb("", 0, n, &combs)
	return combs
}

func genComb(curComb string, start, n int, combs *[]string) {
	if len(curComb) == n {
		*combs = append(*combs, curComb)
		return
	}

	for i := start; i <= 9; i++ {
		newComb := curComb + string('0'+i)
		genComb(newComb, i+1, n, combs)
	}
}
