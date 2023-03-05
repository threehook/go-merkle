package merkle

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (a *UtilsTestSuite) SetupTest() {

}

//func (a *UtilsTestSuite) TestZip() {
//	v1 := []uint{99, 100, 101}
//	v2 := []hash.Hash{sha256.New(), sha256.New(), sha256.New()}
//
//	zip := Zip(v1, v2)
//	fmt.Println(zip)
//}

//func (a *UtilsTestSuite) TestUnzip() {
//	v1 := []uint{99, 100, 101}
//	v2 := []hash.Hash{sha256.NewHasher(), sha256.NewHasher(), sha256.NewHasher()}
//	zipped := Zip(v1, v2)
//	indices, hashes := Unzip(zipped)
//
//	fmt.Println(indices)
//	fmt.Println(hashes)
//}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}
