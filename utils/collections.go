package utils

import (
	"golang.org/x/exp/slices"
)

func Difference[T comparable](aValues, bValues []T) []T {
	diffs := make([]T, 0)
	for _, a := range aValues {
		idx := slices.IndexFunc(bValues, func(b T) bool {
			return a == b
		})
		if idx == -1 {
			diffs = append(diffs, a)
		}
	}
	return diffs
}

func Dedup[T comparable](tSlice []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range tSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func ReverseSlice[T any](s []T) []T {
	a := make([]T, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

//func ReverseSlice[T any](s []T) []T {
//	var r []T
//	for i := len(s) - 1; i >= 0; i-- {
//		r = append(r, s[i])
//	}
//	return r
//}
