package sprint

func Overlap(arr1, arr2 []int) []int {
	if len(arr1) < len(arr2) {
		arr1, arr2 = arr2, arr1
	}
	overlapped := make(map[int]bool)
	for _, val := range arr1 {
		overlapped[val] = false
	}
	result := []int{}
	bubbleSort(arr2)

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr1); j++ {
			if arr1[j] == arr2[i] && !overlapped[arr1[j]] {
				result = append(result, arr1[j])
				overlapped[arr1[j]] = true
				break
			}
		}
	}
	return result
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
