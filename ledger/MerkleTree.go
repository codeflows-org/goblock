package ledger

import "crypto/sha256"

func hashLeaf(leaf [leafSize]byte) [nodeSize]byte {
	return sha256.Sum256(leaf[:])
}

func hashChildren(left []byte, right []byte) [nodeSize]byte {
	return sha256.Sum256(append(left[:], right[:]...))
}

func appendLeaf(store HashStore, leaf [leafSize]byte) {
	hashedLeaf := hashLeaf(leaf)
	store.writeNode(hashedLeaf)
	store.writeLeaf(leaf)
	updateAffectedPath(store, store.getNumberOfNodes()-1)
}

func updateAffectedPath(store HashStore, index int) {
	if index == 0 {
		leftChildren := store.readNode(1)
		rightChildren := store.readNode(2)
		store.updateNode(0, hashChildren(leftChildren[:], rightChildren[:]))
		return
	}
	childrenIndex := getChildrenIndex(index)
	parentIndex := getParentIndex(index)

	if childrenIndex < store.getNumberOfNodes() {
		leftChildren := store.readNode(childrenIndex)
		rightChildren := store.readNode(childrenIndex + 1)
		store.updateNode(0, hashChildren(leftChildren[:], rightChildren[:]))
	}

	updateAffectedPath(store, parentIndex)
}

// Compare two merkle trees

//
