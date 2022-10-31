package large

func OutersectionMap[T comparable](collectionA, collectionB []T) []T {
	if len(collectionB) == 0 {
		return collectionA
	}
	if len(collectionA) == 0 {
		return collectionB
	}
	mapA := make(map[T]struct{})
	for _, a := range collectionA {
		mapA[a] = struct{}{}
	}
	mapB := make(map[T]struct{})
	for _, b := range collectionB {
		mapB[b] = struct{}{}
	}
	for b := range mapB {
		if _, ok := mapA[b]; ok {
			delete(mapA, b)
		} else {
			mapA[b] = struct{}{}
		}
	}
	result := make([]T, len(mapA))
	i := 0
	for k := range mapA {
		result[i] = k
		i++
	}
	return result
}
