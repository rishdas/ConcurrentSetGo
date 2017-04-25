package harrisll

import (
	// "fmt"
	"unsafe"
	"sync/atomic"
)

type node struct {
	key *key
	next *node
	isMarker bool
}

func newNodeKey(key *key) *node {
	newNode := new(node)
	newNode.key = key
	newNode.isMarker = false
	
	return newNode
}

func newMakerNode(n *node) *node {
	k := newKey()
	newMarker := newNodeKey(NewKeyValue(k.minValue0))
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
