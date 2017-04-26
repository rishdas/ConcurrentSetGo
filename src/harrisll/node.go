package harrisll

import (
	// "fmt"
	"unsafe"
	"sync/atomic"
	"utils"
)

type node struct {
	key *utils.Key
	next *node
	isMarker bool
}

func newNodeKey(key *utils.Key) *node {
	newNode := new(node)
	newNode.key = key
	newNode.isMarker = false
	
	return newNode
}

func newMakerNode(n *node) *node {
	k := utils.NewKey()
	newMarker := newNodeKey(utils.NewKeyValue(k.MinValue0))
	newMarker.next = n
	newMarker.isMarker = true

	return newMarker
}

func (t *node) casNext(o *node, n *node) bool {
	return t.next == o &&
		atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&t.next)),
		unsafe.Pointer(o), unsafe.Pointer(n))
}
