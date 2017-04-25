package harrisll

import (
	// "fmt"
)

type HarrisLL struct {
	head *node
	tail *node
}

func NewHarrisLL() *HarrisLL {
	harrisLL := new(HarrisLL)
	key := newKey()
	harrisLL.head = newNodeKey(NewKeyValue(key.minValue0))
        harrisLL.tail = newNodeKey(NewKeyValue(key.maxValue0))
	harrisLL.head.next = harrisLL.tail
	return harrisLL
}

func (harrisLL *HarrisLL) getRef(n *node) *node{
	if n.isMarker == true {
		return n
	} else {
		return n.next
	}
}

func (harrisLL *HarrisLL) Contains(k *key) bool{
	cur := harrisLL.head
	for cur.key.compareTo(k) == true {
		cur = cur.next
	}
	return k.equals(cur.key) && !(cur.next.isMarker == true);
}

func (harrisLL *HarrisLL) Add(k *key) bool{
	var pred *node
	var curr *node
	var succ *node
	retry:
	for true {
		pred = harrisLL.head
		curr = pred.next
		for true {
			succ = curr.next
			for succ.isMarker == true {
				succ = succ.next
				if ! pred.casNext(curr, succ) {
					continue retry
				}
				curr = succ
				succ = succ.next
			}
			if curr.key.compareTo(k) == true {
				pred = curr
				curr = succ
			} else if curr.key.equals(k) {
				return false
			} else {
				node := newNodeKey(k)
				node.next = curr
				if pred.casNext(curr, node) {
					return true
				} else {
					continue retry
				}
			}
		}
	}
	//Dead Code
	return false
}
func (harrisLL *HarrisLL) Remove(k *key) bool {
	var pred *node
	var curr *node
	var succ *node
	retry:

	for true {
		pred = harrisLL.head
		curr = pred.next
		for true {
			succ = curr.next
			for succ.isMarker == true {
				succ = succ.next
				if !pred.casNext(curr, succ) {
					continue retry
				}
				curr = succ
				succ = succ.next
			}
			if curr.key.compareTo(k) {
				pred = curr
				curr = succ
			} else if curr.key.compareTo(k) {
				if !curr.casNext(succ, newMakerNode(succ)) {
					continue retry
				}
				pred.casNext(curr, succ)
				return true
			} else if k.compareTo(curr.key) {
				return false
			}
		}
	}
	//Dead Code
	return false
}
func (harrisLL *HarrisLL) TraversalTest() bool {
	cur := harrisLL.head
	nex := harrisLL.getRef(cur.next)

	for nex != harrisLL.tail {
		cur = nex
		nex = harrisLL.getRef(cur.next)
		if cur.key.compareTo(nex.key) == false {
			return false
		}
	}
	return true
}
