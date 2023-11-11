package algorithms

import (
	"crypto/sha256"
)

type AlgoSha256 struct {
	size uint
}

func NewSha256() AlgoSha256 {
	return AlgoSha256{
		size: 0,
	}
}

func (a AlgoSha256) Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hash := h.Sum(nil)
	a.size = uint(len(hash))

	return hash
}

func (a AlgoSha256) ConcatAndHash(left []byte, right []byte) []byte {
	concatenated := make([]byte, len(left))
	copy(concatenated, left)

	if len(right) > 0 {
		rightNodeClone := make([]byte, len(right))
		copy(rightNodeClone, right)
		concatenated = append(concatenated, rightNodeClone...)
		return a.Hash(concatenated)
	} else {
		return left
	}
}

func (a AlgoSha256) HashSize() uint {
	return a.size
}
