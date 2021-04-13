package ledger

import (
	"crypto/sha256"
	"reflect"
)

const leafSize = 1024
const nodeSize = sha256.Size

type Store interface {
	writeLeaf(value [leafSize]byte)
	writeNode(value [nodeSize]byte)
	updateNode(index int, value [nodeSize]byte)
	readLeaf(index int) *[leafSize]byte
	readNode(index int) *[nodeSize]byte
	findLeafIndex(value [leafSize]byte) int
	readLeaves(lIndex int, rIndex int) [][leafSize]byte
	readNodes(lIndex int, rIndex int) [][nodeSize]byte
	getNumberOfNodes() int
	getNumberOfLeaves() int
}

type MemoryStore struct {
	nodes  [][nodeSize]byte
	leaves [][leafSize]byte
}

func NewMemoryStore() Store {
	return &MemoryStore{
		nodes:  make([][nodeSize]byte, 0),
		leaves: make([][leafSize]byte, 0),
	}
}

func (m *MemoryStore) writeLeaf(value [leafSize]byte) {
	m.leaves = append(m.leaves, value)
}

func (m *MemoryStore) writeNode(value [nodeSize]byte) {
	m.nodes = append(m.nodes, value)
}

func (m *MemoryStore) findLeafIndex(value [leafSize]byte) int {
	for index, elem := range m.leaves {
		if reflect.DeepEqual(elem, value) {
			return index
		}
	}
	return -1
}

func (m *MemoryStore) updateNode(index int, value [nodeSize]byte) {
	m.nodes[index] = value
}

func (m *MemoryStore) readLeaf(index int) *[leafSize]byte {
	if index >= len(m.leaves) {
		var arr [leafSize]byte
		return &arr
	}
	return &m.leaves[index]
}

func (m *MemoryStore) readNode(index int) *[nodeSize]byte {
	if index >= len(m.nodes) {
		var arr [nodeSize]byte
		return &arr
	}
	return &m.nodes[index]
}

func (m *MemoryStore) readLeaves(lIndex int, rIndex int) [][leafSize]byte {
	return m.leaves[lIndex:rIndex]
}

func (m *MemoryStore) readNodes(lIndex int, rIndex int) [][nodeSize]byte {
	return m.nodes[lIndex:rIndex]
}

func (m *MemoryStore) getNumberOfNodes() int {
	return len(m.nodes)
}

func (m *MemoryStore) getNumberOfLeaves() int {
	return len(m.leaves)
}
