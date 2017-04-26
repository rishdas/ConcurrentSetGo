package harrisll

import (
	// "fmt"
	"utils"
)

type HarrisLL struct {
	head *node
	tail *node
}

func NewHarrisLL() *HarrisLL {
	harrisLL := new(HarrisLL)
	key := utils.NewKey()
	harrisLL.head = newNodeKey(utils.NewKeyValue(key.MinValue0))
        harrisLL.tail = newNodeKey(utils.NewKeyValue(key.MaxValue0))
	harrisLL.head.next = harrisLL.tail
	return harrisLL
}

func (harrisLL *HarrisLL) getRef(n *node) *node{
	if n != nil && n.isMarker == true {
		return n
	} else {
		return n.next
	}
}

func (harrisLL *HarrisLL) Contains(k *utils.Key) bool{
	cur := harrisLL.head
	for cur.key.CompareTo(k) == true {
		cur = cur.next
	}
	return k.Equals(cur.key) && !(cur.next != nil && cur.next.isMarker == true);
}

func (harrisLL *HarrisLL) Add(k *utils.Key) bool{
	var pred *node
	var curr *node
	var succ *node
	retry:
	for true {
		pred = harrisLL.head
		curr = pred.next
		for true {
			succ = curr.next
			for succ != nil && succ.isMarker == true {
				succ = succ.next
				if pred.casNext(curr, succ) == false {
					continue retry
				}
				curr = succ
				succ = succ.next
			}
			if curr.key.CompareTo(k) == true {
				pred = curr
				curr = succ
			} else if curr.key.Equals(k) {
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
func (harrisLL *HarrisLL) Remove(k *utils.Key) bool {
	var pred *node
	var curr *node
	var succ *node
	retry:

	for true {
		pred = harrisLL.head
		curr = pred.next
		for true {
			succ = curr.next
			for succ != nil && succ.isMarker == true {
				succ = succ.next
				if pred.casNext(curr, succ) == false {
					continue retry
				}
				curr = succ
				succ = succ.next
			}
			if curr.key.CompareTo(k) == true {
				pred = curr
				curr = succ
			} else if curr.key.Equals(k) == true {
				if curr.casNext(succ, newMakerNode(succ)) == false {
					continue retry
				}
				pred.casNext(curr, succ)
				return true
			} else if k.CompareTo(curr.key) {
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
		if cur.key.CompareTo(nex.key) == false {
			return false
		}
	}
	return true
}
