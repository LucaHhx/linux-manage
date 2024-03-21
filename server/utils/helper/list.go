package helper

func ListToMap[T any, V comparable](list []T, f func(t T) (V, bool)) map[V]T {
	m := make(map[V]T)
	for _, v := range list {
		key, ok := f(v)
		if ok {
			m[key] = v
		}
	}
	return m
}
