package functional

func Map[T, R any](target []T, transformer func(T) R) []R {
	value := make([]R, len(target))
	for i, v := range target {
		value[i] = transformer(v)
	}
	return value
}
