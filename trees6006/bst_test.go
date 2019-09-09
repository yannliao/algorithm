package trees6006

import "testing"

func TestBSTInsert(t *testing.T) {
	var tests = []struct {
		input []int
	}{
		{
			input: []int{8, 2, 4, 9, 3, 6},
		},
	}
	for _, test := range tests {
		t.Logf("Test BST Insert %v", test.input)
		var tree BSTree
		for _, v := range test.input {
			node := &BSTNode{Key: v}
			tree.Insert(node)
			ri := tree.checkri()
			if ri != true {
				t.Errorf("BSTree Insert Fail result = %v", v)
			}
			n, _ := tree.Find(v)
			if n.Key != v {
				t.Errorf("BSTree Insert Fail result = %v", v)
			}
		}
	}
}

func TestDeleteNodeWithoutChildren(t *testing.T) {
	var tree BSTree
	tree.Insert(&BSTNode{Key: 23})
	tree.Insert(&BSTNode{Key: 8})
	tree.Insert(&BSTNode{Key: 4})
	tree.Insert(&BSTNode{Key: 16})
	tree.Insert(&BSTNode{Key: 15})
	tree.Insert(&BSTNode{Key: 42})
	e := tree.Delete(15)
	ri := tree.checkri()
	if e != nil {
		t.Errorf("BSTree Delete Fail")
	}
	if ri != true {
		t.Errorf("BSTree Delete Fail")
	}
	n, _ := tree.Find(15)
	if n != nil {
		t.Errorf("BSTree Delete Fail")
	}
}
