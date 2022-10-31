package medium

import (
	"sort"
)

type ordered interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | byte
}

func Outersection[T ordered](collectionA, collectionB []T) []T {
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

func subtraction[T ordered](collectionA, collectionB []T) []T {
	sort.Slice(collectionB, func(i, j int) bool { return collectionB[i] < collectionB[j] })
	var result []T
	for _, a := range collectionA {
		l := len(collectionB)
		index := sort.Search(l, func(i int) bool {
			return collectionB[i] >= a
		})
		if index < l && collectionB[index] == a {
			continue
		}
		result = append(result, a)
	}
	return result
}

func Subtraction[T ordered](collectionA, collectionB []T) []T {
	if len(collectionB) == 0 || len(collectionA) == 0 {
		return collectionA
	}
	return subtraction(collectionA, collectionB)
}
