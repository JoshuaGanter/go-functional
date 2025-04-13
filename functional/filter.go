package functional

func Filter(elements []int, filterFn func(int) bool) []int {
	result := []int{}
	for _, element := range elements {
		if filterFn(element) {
			result = append(result, element)
		}
	}
	return result
}
