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

func SlidingWindow[T any](sl []T, size int) <-chan []T {
	out := make(chan []T)
	if len(sl) <= size {
		go func () {
			out <- sl
			close(out)
		}()
		return out
	}
	go func () {
		for i := 0; i <= len(sl)-size; i++ {
			out <- sl[i:i+size]
		}
		close(out)
	}()
	return out
}

func Pairs[T any](sl []T) <-chan []T {
	out := make(chan []T)
	if len(sl) < 2 {
		go func() {
			close(out)
		}()
		return out
	}
	go func() {
		for i := 0; i < len(sl)-1; i++ {
			for j := i + 1; j < len(sl); j++ {
				out <- []T{sl[i], sl[j]}
			}
		}
		close(out)
	}()
	return out
}

func GroupBy[V any, K comparable](sl []V, f func(V) K) map[K][]V {
	grouped := make(map[K][]V)
	for _, elem := range sl {
		k := f(elem)
		if _, ok := grouped[k]; !ok {
			grouped[k] = []V{}
		}
		grouped[k] = append(grouped[k], elem)
	}
	return grouped
}
