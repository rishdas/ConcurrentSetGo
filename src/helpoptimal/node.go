package helpoptimal

import (
	// "fmt"
	"unsafe"
	"sync/atomic"
)

type node struct {
	key *key
	next *node
	back *node
}

func newNodeNext(key *key, next *node) *node {
	newNode := new(node)
	newNode.key = key
	newNode.next = next
	newNode.back = nil
	
	return newNode
}
func newNodeBack(pre *node, key *key) *node {
	newNode := new(node)
	newNode.key = key
	newNode.back = pre
	
	return newNode
}
func newNodeKey(key *key) *node {
	newNode := new(node)
	newNode.key = key
	
	return newNode
}

func (t *node) casNext(o *node, n *node) bool {
	// fmt.Printf("t.next: %p\n", t.next)
	// fmt.Printf("o: %p\n", o)
	// fmt.Printf("n: %p\n", n)
	return t.next == o &&
		atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&t.next)),
		unsafe.Pointer(o), unsafe.Pointer(n))
}
