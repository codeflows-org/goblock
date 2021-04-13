package ledger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepth(t *testing.T) {
	assert.Equal(t, getDepth(0), 0)
	assert.Equal(t, getDepth(2), 1)
	assert.Equal(t, getDepth(6), 2)
	assert.Equal(t, getDepth(14), 3)
	assert.Equal(t, getDepth(30), 4)
}

func TestParentIndex(t *testing.T) {
	assert.Equal(t, getParentIndex(0), -1)
	assert.Equal(t, getParentIndex(1), 0)
	assert.Equal(t, getParentIndex(2), 0)
	assert.Equal(t, getParentIndex(3), 1)
	assert.Equal(t, getParentIndex(4), 1)
	assert.Equal(t, getParentIndex(5), 2)
	assert.Equal(t, getParentIndex(6), 2)
	assert.Equal(t, getParentIndex(7), 3)
	assert.Equal(t, getParentIndex(8), 3)
	assert.Equal(t, getParentIndex(9), 4)
	assert.Equal(t, getParentIndex(10), 4)
	assert.Equal(t, getParentIndex(11), 5)
	assert.Equal(t, getParentIndex(12), 5)
	assert.Equal(t, getParentIndex(13), 6)
	assert.Equal(t, getParentIndex(14), 6)
}

func TestLeftChildrenIndex(t *testing.T) {
	assert.Equal(t, getLeftChildrenIndex(0), 1)
	assert.Equal(t, getLeftChildrenIndex(1), 3)
	assert.Equal(t, getLeftChildrenIndex(2), 5)
	assert.Equal(t, getLeftChildrenIndex(3), 7)
	assert.Equal(t, getLeftChildrenIndex(6), 13)
}
