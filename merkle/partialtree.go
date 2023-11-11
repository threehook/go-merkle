package merkle

import (
	"errors"
	utils2 "github.com/threehook/go-merkle/merkle/utils"
	"sort"
)

type PartialTree struct {
	Algorithm Hasher
	Layers    [][]NodeTuple
}

func newPartialTree(algo Hasher) *PartialTree {
	layers := make([][]NodeTuple, 0)
	return &PartialTree{Algorithm: algo, Layers: layers}
}

func (pt *PartialTree) GetLayers() [][]NodeTuple {
	return pt.Layers
}

func (pt *PartialTree) build(partialLayers [][]NodeTuple, depth uint) (*PartialTree, error) {
	layers, err := pt.buildTree(partialLayers, depth)
	if err != nil {
		return nil, err
	}
	pt.Layers = layers

	return pt, nil
}

// This is a general algorithm for building a partial tree. It can be used to extract root from merkle proof, or if a
// complete set of leaves provided as a first argument and no helper indices given, will construct the whole tree.
func (pt *PartialTree) buildTree(partialLayers [][]NodeTuple, fullTreeDepth uint) ([][]NodeTuple, error) {

	partialTree := make([][]NodeTuple, 0)
	currentLayer := make([]NodeTuple, 0)

	// Reversing helper nodes, so we can remove one layer starting from 0 each iteration
	var reversedLayers [][]NodeTuple
	reversedLayers = utils2.ReverseSlice(partialLayers)

	// Iterating through fullTreeDepth instead of len(partialLayers) to address the case of applying changes to a tree
	// when tree requires a resize. In that case len(partialLayer) would be lower than the resulting tree depth
	for i := 0; uint(i) < fullTreeDepth; i++ {
		// Appending helper nodes to the current known nodes
		var nodes []NodeTuple
		nodes, reversedLayers = Pop(reversedLayers)
		if len(nodes) > 0 {
			currentLayer = append(currentLayer, nodes...)
		}
		sort.SliceStable(currentLayer, func(i, j int) bool {
			return currentLayer[i].V1 < currentLayer[j].V1
		})
		// Adding partial layer to the tree
		partialTree = append(partialTree, currentLayer)

		// This empties `current` layer and prepares it to be reused for the next iteration
		indices, hashes := Unzip[uint, []byte](currentLayer)
		currentLayer = []NodeTuple{}
		parentLayerIndicesIt := utils2.ParentIndicesIt(indices)

		err := parentLayerIndicesIt.TryForEach(func(idx uint) error {
			j := int(idx)
			if leftNode := hashes[j*2]; leftNode != nil {
				// Populate `current_layer` back for the next iteration
				var node NodeTuple
				if len(hashes) > j*2+1 {
					node = NodeTuple{
						V1: idx,
						V2: pt.Algorithm.ConcatAndHash(leftNode, hashes[j*2+1]),
					}
				} else {
					node = NodeTuple{
						V1: idx,
						V2: pt.Algorithm.ConcatAndHash(leftNode, []byte{}),
					}
				}
				currentLayer = append(currentLayer, node)
			} else {
				return errors.New("not enough hashes to reconstruct the root")
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	partialTree = append(partialTree, currentLayer)

	return partialTree, nil
}

// Consumes other partial tree into itself, replacing any conflicting nodes with nodes from `other` in the process.
// Doesn't rehash the nodes, so the integrity of the result is not verified.
// It gives an advantage in speed, but should be used only if the integrity of the tree can't be broken, for example,
// it is used in the `.commit` method of the `MerkleTree`, since both partial trees are essentially constructed in place
// and there's no need to verify integrity of the result.
func (pt *PartialTree) mergeUnverified(other *PartialTree) {
}
