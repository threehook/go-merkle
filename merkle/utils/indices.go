package utils

import (
	"math"
	"mtoohey.com/iter"
)

func isLeftIndex(index uint) bool {
	return index%2 == 0
}

func getSiblingIndex(index uint) uint {
	if isLeftIndex(index) {
		// Right sibling index
		return index + 1
	} else {
		// Left sibling index
		return index - 1
	}
}

func SiblingIndices(indices []uint) []uint {
	siblings := make([]uint, 0)
	for _, idx := range indices {
		siblings = append(siblings, getSiblingIndex(idx))
	}
	return siblings
}

func parentIndex(idx uint) uint {
	if isLeftIndex(idx) {
		return idx / 2
	}
	return getSiblingIndex(idx) / 2
}

func ParentIndices(indices []uint) []uint {
	parentIndices := make([]uint, 0)
	for _, idx := range indices {
		parentIndices = append(parentIndices, parentIndex(idx))
	}
	parentIndices = Dedup(parentIndices)
	return parentIndices
}

func ParentIndicesIt(indices []uint) iter.Iter[uint] {
	parentIndices := make([]uint, 0)
	for _, idx := range indices {
		parentIndices = append(parentIndices, parentIndex(idx))
	}
	parentIndices = Dedup(parentIndices)
	return iter.Elems(parentIndices)
}

func TreeDepth(leavesCount uint) uint {
	if leavesCount == 1 {
		return 1
	} else {
		val := float64(leavesCount)
		return uint(math.Ceil(math.Log2(val)))
	}
}

//pub fn tree_depth(leaves_count: usize) -> usize {
//    if leaves_count == 1 {
//        1
//    } else {
//        let val = micromath::F32(leaves_count as f32);
//        val.log2().ceil().0 as usize
//    }
//}
