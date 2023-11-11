package merkle

import (
	"github.com/barweiss/go-tuple"
	"mtoohey.com/iter"
)

func Zip[T uint, U []byte](v1 []T, v2 []U) []tuple.T2[T, U] {
	return iter.Zip(iter.Elems(v1), iter.Elems(v2)).Collect()
}

// []tuple.T2[T, U]
//func Unzip[T uint, U []byte](nt []NodeTuple) ([]T, []U) {
//	nodeTupleIt := iter.Elems[NodeTuple](nt)
//	v1, v2 := iter.Unzip[T, U](nodeTupleIt)
//	indices := make([]T, 0, len(nt))
//	v1.CollectInto(indices)
//	hashes := make([]U, 0, len(nt))
//	v2.CollectInto(hashes)
//
//	return indices, hashes
//}

func Unzip[T uint, U []byte](nt []tuple.T2[T, U]) ([]T, []U) {
	v1, v2 := iter.Unzip[T, U](iter.Elems[tuple.T2[T, U]](nt))
	indices := make([]T, len(nt))
	v1.CollectInto(indices)
	hashes := make([]U, len(nt))
	v2.CollectInto(hashes)

	return indices, hashes
}

// Pop removes the last item of a NodeTuple slice (if exists) and returns it
func Pop[T []NodeTuple](s []T) (T, []T) {
	last := []NodeTuple{}
	if len(s) > 0 {
		last = s[len(s)-1]
		s = s[:len(s)-1]
	}

	return last, s
}
