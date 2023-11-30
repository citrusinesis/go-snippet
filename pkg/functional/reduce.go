package functional

func Reduce[T, R any](target []T, reducer func(R, T) R, initValue R) R {
	acc := initValue
	for _, v := range target {
		acc = reducer(acc, v)
	}
	return acc
}
