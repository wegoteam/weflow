package utils

import "github.com/elliotchance/pie/v2"

func IsContainsSlice[T comparable](slice []T, item T) bool {
	return pie.Contains(slice, item)
}

func IsEmptySlice(slice []interface{}) bool {
	return slice == nil || len(slice) == 0
}
