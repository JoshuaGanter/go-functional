package functional

func Filter[TValue any](slice []TValue, filterFn func(TValue, int, []TValue) bool) []TValue {
	return Reduce(slice, func(previous []TValue, current TValue, index int, slice []TValue) []TValue {
		if filterFn(current, index, slice) {
			return append(previous, current)
		}
		return previous
	}, []TValue{})
}
