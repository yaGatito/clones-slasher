package slicex

func Map[S ~[]E, E any, D any](src S, mapFunc func(e E) D) []D {
	res := make([]D, len(src))
	for i, v := range src {
		res[i] = mapFunc(v)
	}
	return res
}

func Filter[S ~[]E, E any](src S, filterFunc func(e E) bool) S {
	res := make(S, 0)
	for _, v := range src {
		if filterFunc(v) {
			res = append(res, v)
		}
	}
	return res
}
