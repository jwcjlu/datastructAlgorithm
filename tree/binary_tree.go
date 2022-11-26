package tree

type Node struct {
	Key   interface{}
	Value interface{}
	Right *Node
	Left  *Node
}

type BinaryTree struct {
	Root    *Node
	Compare func(interface{}, interface{}) int
}

func (bst *BinaryTree) Put(key interface{}, value interface{}) {
	bst.Root = bst.Insert(bst.Root, key, value)
}
func (bst *BinaryTree) Insert(node *Node, key interface{}, value interface{}) *Node {
	if node == nil {
		return &Node{Key: key, Value: value}

	}
	if bst.Compare(node.Key, key) > 0 {
		node.Right = bst.Insert(node.Right, key, value)
	} else if bst.Compare(node.Key, key) < 0 {
		node.Left = bst.Insert(node.Left, key, value)
	} else {
		node.Value = value
	}
	return node
}

func (bst *BinaryTree) Get(node *Node, key interface{}) interface{} {
	if node == nil {
		return nil
	}
	if bst.Compare(node.Key, key) == 0 {
		return node.Value
	}
	if bst.Compare(node.Key, key) > 0 {
		return bst.Get(node.Right, key)
	}
	if bst.Compare(node.Key, key) < 0 {
		return bst.Get(node.Left, key)
	}
	return nil
}
