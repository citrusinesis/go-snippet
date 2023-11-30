package functional

func Reduce[T, R any](target []T, reducer func(R, T) R, initValue R) R {
	acc := initValue
	for _, cur := range target {
		acc = reducer(acc, cur)
	}
	return acc
}

func ReduceIndexed[T, R any](target []T, reducer func(R, T, int) R, initValue R) R {
	acc := initValue
	for i, cur := range target {
		acc = reducer(acc, cur, i)
	}
	return acc
}
