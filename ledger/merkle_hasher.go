package ledger

import "crypto/sha256"

func hashLeaf(leaf [leafSize]byte) [nodeSize]byte {
	return sha256.Sum256(leaf[:])
}

func hashChildren(left []byte, right []byte) [nodeSize]byte {
	return sha256.Sum256(append(left[:], right[:]...))
}
