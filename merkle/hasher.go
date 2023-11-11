package merkle

type Hasher interface {
	Hash(data []byte) []byte
	ConcatAndHash(left, right []byte) []byte
	HashSize() uint
}
