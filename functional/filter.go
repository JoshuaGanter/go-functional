package functional

func Filter(slice []int, filterFn func(int, int, []int) bool) []int {
	result := []int{}
	for index, element := range slice {
		if filterFn(element, index, slice) {
			result = append(result, element)
		}
	}
	return result
}
