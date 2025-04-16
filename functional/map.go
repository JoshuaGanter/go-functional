package functional

func Map[TValue comparable, TResult comparable](slice []TValue, mapFn func(TValue, int, []TValue) TResult) []TResult {
	var result []TResult
	for index, element := range slice {
		result = append(result, mapFn(element, index, slice))
	}
	return result
}
