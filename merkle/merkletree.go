package merkle

import (
	"github.com/barweiss/go-tuple"
	"github.com/threehook/go-merkle/merkle/algorithms"
	utils2 "github.com/threehook/go-merkle/merkle/utils"
	"sort"
)

type MerkleTree[T Hasher] struct {
	currentWorkingTree PartialTree
	history            []PartialTree
	uncommittedLeaves  [][]byte
}

type NodeTuple = tuple.T2[uint, []byte]

func NewHasher[T Hasher](hasher T) *MerkleTree[T] {
	layers := make([][]NodeTuple, 0)
	return &MerkleTree[T]{
		currentWorkingTree: PartialTree{
			Algorithm: hasher,
			Layers:    layers,
		},
		history:           make([]PartialTree, 0),
		uncommittedLeaves: make([][]byte, 0),
	}
}

func FromLeaves[T Hasher](leaves [][]byte, hasher T) *MerkleTree[T] {
	tree := NewHasher[T](hasher)
	tree.append(leaves)
	tree.commit()

	return tree
}

func (mt *MerkleTree[T]) append(leaves [][]byte) *MerkleTree[T] {
	mt.uncommittedLeaves = append(mt.uncommittedLeaves, leaves...)
	return mt
}

func (mt *MerkleTree[T]) commit() {
	if diff := mt.uncommittedDiff(); diff != nil {
		mt.history = append(mt.history, *diff)
		mt.currentWorkingTree.mergeUnverified(diff)
		mt.uncommittedLeaves = make([][]byte, 0)
	}
}

func (mt *MerkleTree[T]) leavesLen() uint {
	leaves := mt.leavesTuples()
	if leaves != nil {
		return uint(len(leaves))
	} else {
		return 0
	}
}

func (mt *MerkleTree[T]) leavesTuples() []NodeTuple {
	layerTuples := mt.layerTuples()
	if len(layerTuples) > 0 {
		return layerTuples[0]
	}
	return nil
}

func (mt *MerkleTree[T]) layerTuples() [][]NodeTuple {
	return mt.currentWorkingTree.GetLayers()
}

// Creates a diff from changes that were not committed to the main tree yet.
// Can be used to get uncommitted root or can be merged with the main tree
func (mt *MerkleTree[T]) uncommittedDiff() *PartialTree {
	if len(mt.uncommittedLeaves) == 0 {
		return nil
	}
	committedLeavesCount := mt.leavesLen()
	var shadowIndices []uint
	for idx, _ := range mt.uncommittedLeaves {
		shadowIndices = append(shadowIndices, committedLeavesCount+uint(idx))
	}

	shadowNodeTuples := make([]NodeTuple, 0, len(shadowIndices))
	for i, idx := range shadowIndices {
		nodeTuple := NodeTuple{V1: idx, V2: mt.uncommittedLeaves[i]}
		shadowNodeTuples = append(shadowNodeTuples, nodeTuple)
	}
	partialTreeTuples := mt.helperNodeTuples(shadowIndices)

	// Figure out what the tree height would be if we committed the changes
	leavesInNewTree := mt.leavesLen() + uint(len(mt.uncommittedLeaves))
	uncommittedTreeDepth := utils2.TreeDepth(leavesInNewTree)

	if len(partialTreeTuples) > 0 && len(partialTreeTuples[0]) > 0 {
		firstLayer := partialTreeTuples[0]
		partialTreeTuples[0] = append(firstLayer, shadowNodeTuples...)
		// Sort by size, keeping original order or equal elements.
		sort.SliceStable(firstLayer, func(i, j int) bool {
			return firstLayer[i].V1 < firstLayer[j].V1
		})
	} else {
		partialTreeTuples = append(partialTreeTuples, shadowNodeTuples)
	}

	// Building a partial tree with the changes that would be needed in the working tree
	algo := &algorithms.AlgoSha256{}
	tree, _ := newPartialTree(*algo).build(partialTreeTuples, uncommittedTreeDepth)

	return tree

}

func (mt *MerkleTree[T]) helperNodeTuples(leafIndices []uint) [][]NodeTuple {
	currentLayerIndices := leafIndices
	helperNodes := make([][]NodeTuple, 0)

	for _, treeLayer := range mt.layerTuples() {
		helpersLayer := make([]NodeTuple, 0)
		siblings := utils2.SiblingIndices(currentLayerIndices)
		// Filter all nodes that do not require an additional Hash to be calculated
		helperIndices := utils2.Difference(siblings, currentLayerIndices)

		for _, idx := range helperIndices {
			tuple := treeLayer[idx]
			helpersLayer = append(helpersLayer, tuple)
		}
		helperNodes = append(helperNodes, helpersLayer)
		currentLayerIndices = utils2.ParentIndices(currentLayerIndices)
	}
	return helperNodes
}
