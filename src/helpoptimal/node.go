package helpoptimal

import (
	// "fmt"
	"unsafe"
	"sync/atomic"
	"utils"
)

type node struct {
	key *utils.Key
	next *node
	back *node
}

func newNodeNext(key *utils.Key, next *node) *node {
	newNode := new(node)
	newNode.key = key
	newNode.next = next
	newNode.back = nil
	
	return newNode
}
func newNodeBack(pre *node, key *utils.Key) *node {
	newNode := new(node)
	newNode.key = key
	newNode.back = pre
	
	return newNode
}
func newNodeKey(key *utils.Key) *node {
	newNode := new(node)
	newNode.key = key
	
	return newNode
}

func (t *node) casNext(o *node, n *node) bool {
	return t.next == o &&
		atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&t.next)),
		unsafe.Pointer(o), unsafe.Pointer(n))
}
