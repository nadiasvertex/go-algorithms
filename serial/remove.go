package serial

import "github.com/nadiasvertex/algorithms/common"

func Remove[T comparable](collection []T, value T) []T {
	pivot := Partition(collection, func(item T) bool {
		return item != value
	})

	if pivot == len(collection) {
		return nil
	}

	return collection[0:pivot]
}

func RemoveIf[T comparable](collection []T, pred common.Predicate[T]) []T {
	pivot := Partition(collection, func(item T) bool {
		return !pred(item)
	})

	if pivot == len(collection) {
		return nil
	}

	return collection[0:pivot]
}
