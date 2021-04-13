package ledger

import "math"

/* balanced bst array impl helping function */

func getDepth(index int) int {
	if index == 0 {
		return 0
	}
	return highestBitSet(index+1) - 1
}

func getParentIndex(index int) int {
	if index == 0 {
		return -1
	}
	leftestSiblingIndex := powInt(2, highestBitSet(index)-1) - 1
	distance := index - leftestSiblingIndex
	parentLeftestSiblingIndex := leftestSiblingIndex / 2
	parentDistance := distance / 2
	return parentLeftestSiblingIndex + parentDistance
}

func getLeftChildrenIndex(index int) int {
	leftestSiblingIndex := powInt(2, highestBitSet(index)-1)
	distance := index - leftestSiblingIndex
	childrenLeftestSiblingIndex := 2*leftestSiblingIndex + 1
	childrenDistance := distance * 2
	return childrenLeftestSiblingIndex + childrenDistance
}

func highestBitSet(n int) int {
	hiBit := 0
	for n != 0 {
		n >>= 1
		hiBit += 1
	}
	return hiBit
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
