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
