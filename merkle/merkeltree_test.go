package merkle

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/threehook/go-merkle/merkle/algorithms"
	"testing"
)

type MerkleTreeTestSuite struct {
	suite.Suite
}

func (a *MerkleTreeTestSuite) SetupTest() {}

type TestData struct {
	LeafValues      []string
	ExpectedRootHex string
	LeafHashes      [][]byte
}

func (a *MerkleTreeTestSuite) TestGiveCorrectRootAfterCommit() {
	data := setup()
	mtree := FromLeaves[algorithms.AlgoSha256]([][]byte{}, algorithms.NewSha256())
	mtree2 := FromLeaves[algorithms.AlgoSha256](data.LeafHashes, algorithms.NewSha256())

	fmt.Println(mtree)
	fmt.Println(mtree2)
}

func TestMerkleTreeTestSuite(t *testing.T) {
	suite.Run(t, new(MerkleTreeTestSuite))
}

func setup() TestData {
	leaveValues := []string{"a", "b", "c", "d", "e", "f"}
	leafHashes := make([][]byte, 0, 6)

	sha256Hasher := algorithms.NewSha256()
	for _, s := range leaveValues {
		hash := sha256Hasher.Hash([]byte(s))
		leafHashes = append(leafHashes, hash)
	}

	return TestData{
		LeafValues:      leaveValues,
		ExpectedRootHex: "1f7379539707bcaea00564168d1d4d626b09b73f8a2a365234c62d763f854da2",
		LeafHashes:      leafHashes,
	}

}
