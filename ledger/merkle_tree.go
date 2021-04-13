package ledger

func appendLeaf(store Store, leaf [leafSize]byte) {
	hashedLeaf := hashLeaf(leaf)
	store.writeNode(hashedLeaf)
	store.writeLeaf(leaf)
	update(store, store.getNumberOfNodes()-1)
}

/* add new nodes */
func update(store Store, index int) {
	if index == 0 {
		leftChildren := store.readNode(1)
		rightChildren := store.readNode(2)
		store.updateNode(0, hashChildren(leftChildren[:], rightChildren[:]))
		return
	}
	childrenIndex := getLeftChildrenIndex(index)
	parentIndex := getParentIndex(index)

	if childrenIndex < store.getNumberOfNodes() {
		leftChildren := store.readNode(childrenIndex)
		rightChildren := store.readNode(childrenIndex + 1)
		store.updateNode(0, hashChildren(leftChildren[:], rightChildren[:]))
	}

	update(store, parentIndex)
}
