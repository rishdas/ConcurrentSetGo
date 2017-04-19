package helpoptimal

type node struct {
	key *key
	next *node
	back *node
}

func newNodeNext(key *key, next *node) *node {
	newNode := new(node)
	newNode.key = key
	newNode.next = next
	
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
