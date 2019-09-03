package trees6006

// An BSTree is a binary search tree of integers.
// A nil *BSTree represents the empty tree.
type BSTree struct {
	key                 int
	parent, left, right *BSTree
}

//Find the tree node of value k
func (t *BSTree) Find(k int) *BSTree {
	for t != nil && k != t.key {
		if k < t.key {
			t = t.left
		} else {
			t = t.right
		}
	}
	return t
}

//FindMin find the node of minimum key in the tree
func (t *BSTree) FindMin() *BSTree {
	for t.left != nil {
		t = t.left
	}
	return t
}

//FindMax find the max key node in the binary search tree
func (t *BSTree) FindMax() *BSTree {
	for t.right != nil {
		t = t.right
	}
	return t
}

//Successor get the node with the next larger key
func (t *BSTree) Successor() *BSTree {
	return t
}

//Predecessor get the node with the next smaller key
func (t *BSTree) Predecessor() *BSTree {
	return t
}

//Insert an node to the binary search tree
func (t *BSTree) Insert(n *BSTree) {
	if n == nil {
		return
	}
	if t == nil {
		t = n
	}
	if n.key < t.key {
		if t.left == nil {
			t.left = n
			n.parent = t
		} else {
			t.left.Insert(n)
		}
	} else {
		if t.right == nil {
			t.right = n
			n.parent = t
		} else {
			t.right.Insert(n)
		}
	}
}

//Delete an node from the bst
func (t *BSTree) Delete() {

}
