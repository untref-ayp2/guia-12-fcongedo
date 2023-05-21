package ejercicios

import "fmt"

type BinaryTree struct {
	root *BinaryNode
}

func NewBinaryTree(data int) *BinaryTree {
	node := NewBinaryNode(data, nil, nil)
	return &BinaryTree{root: node}

}

func (t *BinaryTree) InsertLeft(bt *BinaryTree) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.left = bt.root
	}
}
func (t *BinaryTree) InsertRight(bt *BinaryTree) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.right = bt.root
	}
}

func (t *BinaryTree) PrintPreOrder() {
	if t.root != nil {
		t.root.PrintPreOrder()
	}
}

func (t *BinaryTree) PrintInOrder() {
	if t.root != nil {
		t.root.PrintInOrder()
	}
}

func (t *BinaryTree) PrintPostOrder() {
	if t.root != nil {
		t.root.PrintPostOrder()
	}
}

func (t *BinaryTree) Empty() {
	t.root = nil
}

func (t *BinaryTree) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree) Size() int {
	if t.root != nil {
		return t.root.Size()
	} else {
		return 0
	}
}

func (t *BinaryTree) Height() int {
	if t.root != nil {
		return t.root.Height()
	} else {
		return -1
	}
}

func (t *BinaryTree) SumaNodos() int {
	if t.root != nil {
		return t.root.SumaNodos()
	} else {
		return 0
	}
}

func (t *BinaryTree) SumaHojas() int {
	if t.root != nil {
		return t.root.SumaHojas()
	} else {
		return 0
	}
}

func (t *BinaryTree) SumaNodosInternos() int {
	if t.root != nil {
		return t.root.SumaNodosInternos() - t.root.data
	} else {
		return 0
	}
}
func (t *BinaryTree) SumaNodosPares() int {
	if t.root != nil {
		return t.root.SumaNodosPares()
	} else {
		return 0
	}
}
func (t *BinaryTree) SumaNodosMayorIgualQueSeis() int {
	if t.root != nil {
		return t.root.SumaNodosMayorIgualQueSeis()
	} else {
		return 0
	}
}
func (t *BinaryTree) CalcularAlturaMaxima() int {
	if t.root != nil {
		return t.root.CalcularAlturaMaxima()
	} else {
		return -1
	}
}

func (t *BinaryTree) SumaHojasIzquierdas() int {
	if t.root != nil {
		return t.root.SumaHojasIzquierdas()
	} else {
		return 0
	}
}

func (t *BinaryTree) SumaDeNodosDerechosImpares() int {
	if t.root != nil {
		return t.root.SumaDeNodosDerechosImpares()
	} else {
		return 0
	}
}
func (t *BinaryTree) PrintLevelOrder() {
	if t.root == nil {
		return
	}

	queue := []*BinaryNode{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.data)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}
