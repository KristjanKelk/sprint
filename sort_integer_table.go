package sprint

func SortIntegerTable(table []int) []int {
	max:=tabel[0]
	result:= []int
	for i:=0, i < len(arr); i++{
		if max < tabel[i]{
			max = tabel[i]
		}
	}

	for len(tabel) > 0 {
		min:= 0
		for i:=0; i < len(tabel); i++{
			if tabel[i] < tabel[min] {
				min = i
			}

		}
		result = append(result, tabel[min])
		tabel = append(tabel[:min], tabel[min+1]...)
		
	}

	}
	

		


