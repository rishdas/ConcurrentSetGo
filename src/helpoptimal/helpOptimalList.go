package helpoptimal

import (
	// "fmt"
)

type HelpOptimalLFList struct {
	head *node
	headNext *node
	tail *node
	tailNext *node
}

func NewHelpOptimalLFList() *HelpOptimalLFList {
	hoLFList := new(HelpOptimalLFList)
	key := newKey()
	hoLFList.tailNext = newNodeKey(NewKeyValue(key.maxValue0));
        hoLFList.tail = newNodeNext(NewKeyValue(key.maxValue1), hoLFList.tailNext);
        hoLFList.headNext = newNodeNext(NewKeyValue(key.minValue1), hoLFList.tail);
        hoLFList.head = newNodeNext(NewKeyValue(key.minValue0), hoLFList.headNext);
	return hoLFList
}

func (hoLFList *HelpOptimalLFList) getRef(n *node) *node{
	if n.back == nil {
		return n
	} else {
		return n.next
	}
}
func (hoLFList *HelpOptimalLFList) getNext(n *node) *node{
	return hoLFList.getRef(n.next)
}

func (hoLFList *HelpOptimalLFList) Contains(k *key) bool{
	cur := hoLFList.getNext(hoLFList.headNext);
	for cur.key.compareTo(k) == true {
		cur = cur.next
	}
	return k.equals(cur.key) && cur.next.back == nil;
}

func (hoLFList *HelpOptimalLFList) Add(k *key) bool{
	pre := hoLFList.head
	suc := hoLFList.headNext
	cur := hoLFList.headNext
	nex := cur.next
	for true {
		//Search
		for cur.key.compareTo(k) == true {
			if nex.back == nil {
				pre = cur
				suc = nex
				cur = suc
			} else {
				cur = nex.next
			}
			nex = cur.next
		}
		//isSplice Node
		if (nex.back != nil) {
			//Traverses Splice nodes and Helps removal
			for nex.back != nil {
				cur = nex.next
				nex = cur.next
			}
		} else if cur.key.equals(k) == true {
			return false
		}
		//CAS p(k).nxt from s(k) to k  

		if pre.casNext(suc, newNodeNext(k, cur)) == true {
			return true
		}
		//Local Back track
		suc = pre.next
		for suc.back != nil {
			pre = suc.back
			suc = pre.next
		}
		cur = pre
		nex = suc
	}
	//Dead Code
	return false
}
func (hoLFList *HelpOptimalLFList) Remove(k *key) bool {
	pre := hoLFList.head
	suc := hoLFList.headNext
	cur := hoLFList.headNext
	nex := cur.next
	var marker *node
	mode := true
	nk := newKey()

	for true {
		for cur.key.compareTo(k) == true {
			if nex.back == nil {
				pre = cur
				suc = nex
				cur = suc
			} else {
				cur = nex.next
			}
			nex = cur.next
		}
		if mode == true {
			//key not found or already logically removed
			if k.equals(cur.key) == false || nex.back != nil {
				return false
			}
			marker = newNodeBack(pre, NewKeyValue(nk.minValue0))
			for true {
				marker.next = nex
				if cur.casNext(nex, marker) == true { //Logically Removing
					if pre.casNext(suc, nex) {
						return true
					}
					mode = false
					break
				}
				nex = cur.next
				if nex.back != nil {
					return false
				}
			}
		} else if nex != marker || pre.casNext(suc, nex.next) {
			return true
		}
		suc = pre.next
		for suc.back != nil {
			pre = suc.back
			suc = pre.next
		}
		cur = pre
		nex = suc
	}
	//Dead Code
	return false
}
func (hoLFList *HelpOptimalLFList) TraversalTest() bool {
	cur := hoLFList.head
	nex := hoLFList.getRef(cur.next)

	for cur != hoLFList.tail {
		cur = nex
		nex = hoLFList.getRef(cur.next)
		if cur.key.compareTo(nex.key) == false {
			return false
		}
	}
	return true
}
