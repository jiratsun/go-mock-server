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
