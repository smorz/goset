package small

func IsMember[T comparable](elem T, collection []T) bool {
	for _, e := range collection {
		if elem == e {
			return true
		}
	}
	return false
}

func Outersection[T comparable](collectionA, collectionB []T) []T {
	if len(collectionB) == 0 {
		return collectionA
	}
	if len(collectionA) == 0 {
		return collectionB
	}
	return append(
		subtraction(collectionA, collectionB),
		subtraction(collectionB, collectionA)...,
	)
}

func subtraction[T comparable](collectionA, collectionB []T) []T {
	result := make([]T, 0, len(collectionA))
	for _, a := range collectionA {
		func() {
			for _, b := range collectionB {
				if a == b {
					return
				}
			}
			result = append(result, a)
		}()
	}
	return result
}

func Subtraction[T comparable](collectionA, collectionB []T) []T {
	if len(collectionB) == 0 || len(collectionA) == 0 {
		return collectionA
	}
	return subtraction(collectionA, collectionB)
}

func Union[T any](collectionA, collectionB []T) []T {
	if len(collectionB) == 0 {
		return collectionA
	}
	if len(collectionA) == 0 {
		return collectionB
	}
	return append(collectionA, collectionB...)
}

func IsSubset[T comparable](subset, collection []T) bool {
	if len(subset) == 0 {
		return true
	}
	if len(collection) == 0 {
		return false
	}
	var isMember bool
	for _, s := range subset {
		isMember = false
		for _, c := range collection {
			if s == c {
				isMember = true
				break
			}
		}
		if !isMember {
			return false
		}
	}
	return true
}
