package utils

func UniqueBy[T any, K comparable](slice []T, fn func(T) K) []T {
	seen := make(map[K]struct{})
	result := make([]T, 0, len(slice))

	for _, item := range slice {
		key := fn(item)
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result
}
