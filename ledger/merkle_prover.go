package ledger

func auditProof(store Store, leaf [leafSize]byte) *[][nodeSize]byte {
	index := store.findLeafIndex(leaf)
	parentIndex := getParentIndex(index)
	var auditPath *[][nodeSize]byte
	buildAuditTrail(store, parentIndex, index, auditPath)
	return auditPath
}

func buildAuditTrail(store Store, parentIndex int, nodeIndex int, auditPath *[][nodeSize]byte) {
	leftChildrenIndex := getLeftChildrenIndex(parentIndex)
	rightChildrenIndex := leftChildrenIndex + 1
	var nextChildren int
	if leftChildrenIndex == nodeIndex {
		nextChildren = rightChildrenIndex
	} else {
		nextChildren = leftChildrenIndex
	}
	*auditPath = append(*auditPath, *store.readNode(nextChildren))
	buildAuditTrail(store, getLeftChildrenIndex(parentIndex), parentIndex, auditPath)
}
