package ejercicios

import (
	"fmt"
	"math"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

func NewBinaryNode(data int, left *BinaryNode, right *BinaryNode) *BinaryNode {
	return &BinaryNode{left: left, right: right, data: data}
}

func (n *BinaryNode) PrintPreOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

func (n *BinaryNode) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

func (n *BinaryNode) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.data)
}

func (n *BinaryNode) Size() int {
	size := 1
	if n.left != nil {
		size += n.left.Size()
	}
	if n.right != nil {
		size += n.right.Size()
	}
	return size
}

func (n *BinaryNode) Height() int {
	leftHeight := -1
	rightHeight := -1
	if n.left != nil {
		leftHeight = n.left.Height()
	}
	if n.right != nil {
		rightHeight = n.right.Height()
	}
	return int(1 + math.Max(float64(leftHeight), float64(rightHeight)))
}

func (n *BinaryNode) SumaNodos() int {
	suma := n.data
	if n.left != nil {
		suma += n.left.SumaNodos()
	}
	if n.right != nil {
		suma += n.right.SumaNodos()
	}
	return suma
}

func (n *BinaryNode) SumaHojas() int {
	if n.left == nil && n.right == nil {
		return n.data
	}
	suma := 0
	if n.left != nil {
		suma += n.left.SumaHojas()
	}
	if n.right != nil {
		suma += n.right.SumaHojas()
	}
	return suma
}

func (n *BinaryNode) SumaNodosInternos() int {
	if n.right == nil && n.left == nil {
		return 0
	}
	suma := n.data
	if n.right != nil {
		suma += n.right.SumaNodosInternos()
	}
	if n.left != nil {
		suma += n.left.SumaNodosInternos()
	}
	return suma

}
func (n *BinaryNode) SumaNodosPares() int {
	suma := 0
	if n != nil {
		if n.data%2 == 0 {
			suma += n.data
		}
		if n.left != nil {
			suma += n.left.SumaNodosPares()
		}
		if n.right != nil {
			suma += n.right.SumaNodosPares()
		}
	}
	return suma
}
func (n *BinaryNode) SumaNodosMayorIgualQueSeis() int {
	suma := 0
	if n != nil {
		if n.data >= 6 {
			suma += n.data
		}
		if n.left != nil {
			suma += n.left.SumaNodosMayorIgualQueSeis()
		}
		if n.right != nil {
			suma += n.right.SumaNodosMayorIgualQueSeis()
		}
	}
	return suma
}
func (n *BinaryNode) CalcularAlturaMaxima() int {
	leftHeight := -1
	rightHeight := -1
	if n.left != nil {
		leftHeight = n.left.CalcularAlturaMaxima()
	}
	if n.right != nil {
		rightHeight = n.right.CalcularAlturaMaxima()
	}
	return int(1 + math.Max(float64(leftHeight), float64(rightHeight)))
}

func (n *BinaryNode) SumaHojasIzquierdas() int {
	sum := 0
	if n.left != nil {
		if n.left.left == nil && n.left.right == nil {
			sum += n.left.data
		}
		sum += n.left.SumaHojasIzquierdas()
	}
	sum += n.right.SumaHojasIzquierdas()
	return sum
}

func (n *BinaryNode) SumaDeNodosDerechosImpares() int {
	suma := 0
	if n != nil {
		if n.right != nil && n.right.data%2 != 0 {
			suma += n.right.data
		}
	}
	if n.left != nil {
		suma += n.left.SumaDeNodosDerechosImpares()
	}
	if n.right != nil {
		suma += n.right.SumaDeNodosDerechosImpares()
	}
	return suma

}
