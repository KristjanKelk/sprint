package sprint

var romanNumerals = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRoman(num int) string {
	result := ""

	for _, rn := range romanNumerals {
		for num >= rn.Value {
			result += rn.Symbol
			num -= rn.Value
		}
	}

	return result
}
