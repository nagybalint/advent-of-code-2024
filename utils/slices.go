package utils

func Filter(in []string, test func(string) bool) (ret []string) {
	for _, s := range in {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
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
