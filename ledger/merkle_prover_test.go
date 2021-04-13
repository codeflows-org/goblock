package ledger

import "testing"

func TestAuditProof(t *testing.T) {
	s := NewMemoryStore()
	for i := 1; i <= 10; i++ {
		var arr [1024]byte
		arr[i] = 1
		appendLeaf(s, arr)
	}

	var arr [1024]byte
	arr[1] = 1

	auditProof(s, arr)
}
