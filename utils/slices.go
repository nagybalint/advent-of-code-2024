package utils

func Filter[T any](in []T, test func(T) bool) (ret []T) {
	for _, s := range in {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func IsNonEmptyString(s string) bool {
	return s != ""
}

func SlidingWindow[T any](sl []T, size int) (ret [][]T) {
	if len(sl) <= size {
		ret = append(ret, sl)
		return
	}
	for i := 0; i <= len(sl)-size; i++ {
		ret = append(ret, sl[i:i+size])
	}
	return
}
