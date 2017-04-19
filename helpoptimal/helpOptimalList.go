package helpoptimal


type helpOptimalLFList struct {
	head *node
	headNext *node
	tail *node
	tailNext *node
}

func newHelpOptimalLFList() *helpOptimalLFList {
	hoLFList := new(helpOptimalLFList)
	key := newKey()
	hoLFList.tailNext = newNodeKey(newKeyValue(key.maxValue0));
        hoLFList.tail = newNodeNext(newKeyValue(key.maxValue1), hoLFList.tailNext);
        hoLFList.headNext = newNodeNext(newKeyValue(key.minValue1), hoLFList.tail);
        hoLFList.head = newNodeNext(newKeyValue(key.minValue0), hoLFList.headNext);
	return hoLFList
}

func (hoLFList *helpOptimalLFList) getRef(n *node) *node{
	if n.back == nil {
		return n
	} else {
		return n.next
	}
}
func (hoLFList *helpOptimalLFList) getNext(n *node) *node{
	return hoLFList.getRef(n.next)
}

func (hoLFList *helpOptimalLFList) contains(k *key) bool{
	cur := hoLFList.getNext(hoLFList.headNext);
	for cur.key.compareTo(k) == true {
		cur = cur.next
	}
	return k.equals(cur.key) && cur.next.back == nil;
}

func (hoLFList *helpOptimalLFList) add(k *key) bool{
	pre := hoLFList.head
	suc := hoLFList.headNext
	cur := hoLFList.headNext
	nex := cur.next
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
		if (nex.back != nil) {
			for nex.back != nil {
				cur = nex.next
				nex = cur.next
			}
		} else if cur.key.equals(k) == true {
			return false
		}
		if pre.casNext(suc, newNodeNext(k, cur)) == true {
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
func (hoLFList *helpOptimalLFList) remove(k *key) bool {
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
			if k.equals(cur.key) == false || nex.back != nil {
				return false
			}
			marker = newNodeBack(pre, newKeyValue(nk.minValue0))
			for true {
				marker.next = nex
				if cur.casNext(nex, marker) == true {
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
