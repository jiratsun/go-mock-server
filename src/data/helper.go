package data

func Map[T, U any](input []T, f func(T) U) []U {
	result := make([]U, len(input))

	for i, v := range input {
		result[i] = f(v)
	}

	return result
}

func ForEach[T any](input []T, f func(T)) {
	for _, v := range input {
		f(v)
	}
}

func Default[T any](v *T, def T) T {
	if v == nil {
		return def
	}

	return *v
}

func Left(s string, num int) string {
	if len(s) > num {
		return s[:num]
	}
	return s
}
