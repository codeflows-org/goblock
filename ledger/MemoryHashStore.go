package ledger

import "crypto/sha256"

const leafSize = 1024
const nodeSize = sha256.Size

type HashStore interface {
	writeLeaf(value [leafSize]byte)
	writeNode(value [nodeSize]byte)
	updateNode(index int, value [nodeSize]byte)
	readLeaf(index int) *[leafSize]byte
	readNode(index int) *[nodeSize]byte
	readLeaves(lIndex int, rIndex int) [][leafSize]byte
	readNodes(lIndex int, rIndex int) [][nodeSize]byte
	getNumberOfNodes() int
	getNumberOfLeaves() int
}

type MemoryHashStore struct {
	nodes  [][nodeSize]byte
	leaves [][leafSize]byte
}

func (m *MemoryHashStore) writeLeaf(value [leafSize]byte) {
	m.leaves = append(m.leaves, value)
}

func (m *MemoryHashStore) writeNode(value [nodeSize]byte) {
	m.nodes = append(m.nodes, value)
}

func (m *MemoryHashStore) updateNode(index int, value [nodeSize]byte) {
	m.nodes[index] = value
}

func (m *MemoryHashStore) readLeaf(index int) *[leafSize]byte {
	if index >= len(m.leaves) {
		return nil
	}
	return &m.leaves[index]
}

func (m *MemoryHashStore) readNode(index int) *[nodeSize]byte {
	if index >= len(m.nodes) {
		return nil
	}
	return &m.nodes[index]
}

func (m *MemoryHashStore) readLeaves(lIndex int, rIndex int) [][leafSize]byte {
	return m.leaves[lIndex:rIndex]
}

func (m *MemoryHashStore) readNodes(lIndex int, rIndex int) [][nodeSize]byte {
	return m.nodes[lIndex:rIndex]
}

func (m *MemoryHashStore) getNumberOfNodes() int {
	return len(m.nodes)
}

func (m *MemoryHashStore) getNumberOfLeaves() int {
	return len(m.leaves)
}
