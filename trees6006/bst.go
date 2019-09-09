package trees6006

import "fmt"

//BSTNode cant be nil
type BSTNode struct {
	Key                 int
	parent, left, right *BSTNode
}

func (n *BSTNode) min() *BSTNode {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *BSTNode) max() *BSTNode {
	for n.right != nil {
		n = n.right
	}
	return n
}
func (n *BSTNode) find(k int) *BSTNode {
	for n != nil && k != n.Key {
		if k < n.Key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}

func (n *BSTNode) insert(m *BSTNode) error {
	if m == nil {
		return fmt.Errorf("Insert an empty BSTnode")
	}

	if m.Key < n.Key {
		if n.left == nil {
			n.left = m
			m.parent = n
		} else {
			n.left.insert(m)
		}
	} else {
		if n.right == nil {
			n.right = m
			m.parent = n
		} else {
			n.right.insert(m)
		}
	}
	return nil
}

func (n *BSTNode) delete() {
	if n.parent == nil {
		return
	}
	if n.left == nil && n.right == nil {
		if n == n.parent.right {
			n.parent.right = nil
		} else {
			n.parent.left = nil
		}
	} else if n.left == nil && n.right != nil {
		n.right.parent = n.parent
		if n == n.parent.right {
			n.parent.right = n.right
		} else {
			n.parent.left = n.right
		}
	} else if n.left != nil && n.right == nil {
		n.left.parent = n.parent
		if n == n.parent.right {
			n.parent.right = n.left
		} else {
			n.parent.left = n.left
		}
	} else {
		next := n.successor()
		n.Key, next.Key = next.Key, n.Key
		next.delete()

		// if n == n.parent.right {
		// 	n.parent.right = next
		// } else {
		// 	n.parent.left = next
		// }
		// next.parent = n.parent
		// next.left = n.left
		// next.left.parent = next
		// next.right = n.right
		// next.right.parent = next
	}
}

//successor get the BSTnode with the next larger key
func (n *BSTNode) successor() *BSTNode {
	if n.right != nil {
		return n.right.min()
	}
	for n.parent != nil && n == n.parent.right {
		n = n.parent
	}
	return n
}

func (n *BSTNode) checkri() bool {
	l, r := true, true
	if n.left != nil {
		if n.left.Key > n.Key {
			fmt.Printf("BST RI violated by a left node key. left: %v node: %v\n", n.left.Key, n.Key)
			return false
		}
		if n.left.parent != n {
			fmt.Printf("BST RI violated by a left node parent pointer. node %v.\n", n.Key)
			return false
		}
		l = n.left.checkri()
	}
	if n.right != nil {
		if n.right.Key < n.Key {
			fmt.Printf("BST RI violated by a right node key.right: %v, node： %v\n", n.right.Key, n.Key)
			return false
		}
		if n.right.parent != n {
			fmt.Printf("BST RI violated by a right node parent pointer.node: %v\n", n.Key)
			return false
		}
		r = n.right.checkri()
	}
	return l && r
}

//predecessor get the BSTnode with the next smaller key
func (n *BSTNode) predecessor() *BSTNode {
	if n.left != nil {
		return n.left.max()
	}
	for n.parent != nil && n == n.parent.left {
		n = n.parent
	}
	return n
}

// An BSTree is a binary search tree of integers.
// A nil *BSTree represents the empty tree.
type BSTree struct {
	root *BSTNode
}

func (t *BSTree) checkri() bool {
	if t.root != nil {
		if t.root.parent != nil {
			fmt.Printf("BST RI violated by the root node's parent pointer.\n")
		}
		return t.root.checkri()
	}
	return true
}

//Min find the minimum key in the binary search tree
func (t *BSTree) Min() (int, error) {
	if t.root == nil {
		return 0, fmt.Errorf("Min() called on empty tree")
	}
	key := t.root.min().Key
	return key, nil
}

//Max find the max key in the binary search tree
func (t *BSTree) Max() (int, error) {
	if t.root == nil {
		return 0, fmt.Errorf("Min() called on empty tree")
	}
	k := t.root.max().Key
	return k, nil
}

//Find the tree BSTnode of value k
func (t *BSTree) Find(k int) (*BSTNode, error) {
	if t.root == nil {
		return nil, fmt.Errorf("Find() called on empty tree")
	}
	c := t.root
	n := c.find(k)
	return n, nil
}

//Insert an BSTnode to the binary search tree
func (t *BSTree) Insert(n *BSTNode) error {
	if n == nil {
		return fmt.Errorf("Insert an empty BSTnode")
	}

	if t.root == nil {
		t.root = n
		return nil
	}

	c := t.root
	return c.insert(n)
}

//Delete an BSTnode from the bst
func (t *BSTree) Delete(k int) error {
	if t.root == nil {
		return fmt.Errorf("Delete() called on empty tree")
	}
	n := t.root.find(k)
	if n == nil {
		return fmt.Errorf("cant Delete an nil BSTnode")
	}
	if n == t.root {
		t.root = nil
	} else {
		n.delete()
	}
	return nil

}
