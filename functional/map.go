package functional

func Map[TValue any, TResult any](slice []TValue, mapFn func(TValue, int, []TValue) TResult) []TResult {
	return Reduce(slice, func(previous []TResult, current TValue, index int, innerSlice []TValue) []TResult {
		previous[index] = mapFn(current, index, innerSlice)
		return previous
	}, make([]TResult, len(slice)))
}
