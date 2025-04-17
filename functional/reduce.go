package functional

func Reduce[TValue comparable, TResult comparable](slice []TValue, reduceFn func(previous TResult, current TValue, index int, slice []TValue) TResult, initial TResult) TResult {
	result := initial

	for index, element := range slice {
		result = reduceFn(result, element, index, slice)
	}

	return result
}
