package trees6006

import "fmt"

//Node cant be nil
type Node struct {
	Key                 int
	parent, left, right *Node
}

func (n *Node) min() *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *Node) max() *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}
func (n *Node) find(k int) *Node {
	for n != nil && k != n.Key {
		if k < n.Key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}

func (n *Node) insert(m *Node) error {
	if m == nil {
		return fmt.Errorf("Insert an empty node")
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

func (n *Node) delete() {
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

//successor get the node with the next larger key
func (n *Node) successor() *Node {
	if n.right != nil {
		return n.right.min()
	}
	for n.parent != nil && n == n.parent.right {
		n = n.parent
	}
	return n
}

//predecessor get the node with the next smaller key
func (n *Node) predecessor() *Node {
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
	root *Node
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

//Find the tree node of value k
func (t *BSTree) Find(k int) (*Node, error) {
	if t.root == nil {
		return nil, fmt.Errorf("Find() called on empty tree")
	}
	c := t.root
	n := c.find(k)
	return n, nil
}

//Insert an node to the binary search tree
func (t *BSTree) Insert(n *Node) error {
	if n == nil {
		return fmt.Errorf("Insert an empty node")
	}

	if t.root == nil {
		t.root = n
		return nil
	}

	c := t.root
	return c.insert(n)
}

//Delete an node from the bst
func (t *BSTree) Delete(k int) error {
	if t.root == nil {
		return fmt.Errorf("Delete() called on empty tree")
	}
	n := t.root.find(k)
	if n == nil {
		return fmt.Errorf("cant Delete an nil node")
	}
	if n == t.root {
		t.root = nil
	} else {
		n.delete()
	}
	return nil

}
