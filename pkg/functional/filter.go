package functional

func Filter[T any](target []T, checker func(T) bool) []T {
	var value []T
	for _, v := range target {
		if checker(v) {
			value = append(value, v)
		}
	}
	return value
}

func FilterIndexed[T any](target []T, checker func(T, int) bool) []T {
	var value []T
	for i, v := range target {
		if checker(v, i) {
			value = append(value, v)
		}
	}
	return value
}
