package functional

func Filter[TValue comparable](slice []TValue, filterFn func(TValue, int, []TValue) bool) []TValue {
	result := []TValue{}
	for index, element := range slice {
		if filterFn(element, index, slice) {
			result = append(result, element)
		}
	}
	return result
}
