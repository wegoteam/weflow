package utils

import "github.com/elliotchance/pie/v2"

func IsContainsSlice[T comparable](slice []T, item T) bool {
	return pie.Contains(slice, item)
}

func IsNotContainsSlice[T comparable](slice []T, item T) bool {
	return !pie.Contains(slice, item)
}

func IsEmptySlice[T comparable](slice []T) bool {
	return slice == nil || len(slice) == 0
}

func IsNotEmptySlice[T comparable](slice []T) bool {
	return !IsEmptySlice(slice)
}
