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




