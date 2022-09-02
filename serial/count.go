package serial

import "github.com/nadiasvertex/algorithms/common"

func Count[T comparable](collection []T, value T) int {
	count := 0
	for _, item := range collection {
		if item == value {
			count++
		}
	}
	return count
}

func CountIf[T comparable](collection []T, pred common.Predicate[T]) int {
	count := 0
	for _, item := range collection {
		if pred(item) {
			count++
		}
	}
	return count
}
