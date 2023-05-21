package tests

import (
	"guia12/binarytree"
	"guia12/ejercicios"
	"testing"
)

func TestTamañoYAlturaDeSoloRaiz(t *testing.T) {
	btree := binarytree.NewBinaryTree("-")
	if btree.Size() != 1 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 1, btree.Size())
	}
	if btree.Height() != 0 {
		t.Errorf("Error la altura deberia dar %v, pero dio %v", 0, btree.Height())
	}
}

func TestTamañoYAlturaDeSoloRaizConHijoIzquierdo(t *testing.T) {
	btree1 := binarytree.NewBinaryTree("+")
	btree2 := binarytree.NewBinaryTree("-")
	btree2.InsertLeft(btree1)
	if btree2.Size() != 2 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 2, btree2.Size())
	}
	if btree2.Height() != 1 {
		t.Errorf("Error la altura deberia dar %v, pero dio %v", 1, btree2.Height())
	}
}

func TestTamañoYAlturaDeSoloRaizConHijoDerecho(t *testing.T) {
	btree1 := binarytree.NewBinaryTree("+")
	btree2 := binarytree.NewBinaryTree("-")
	btree2.InsertRight(btree1)
	if btree2.Size() != 2 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 2, btree2.Size())
	}
	if btree2.Height() != 1 {
		t.Errorf("Error la altura deberia dar %v, pero dio %v", 1, btree2.Height())
	}
}

func TestTamañoYAlturaDeRaizYAmbosHijos(t *testing.T) {
	btree1 := binarytree.NewBinaryTree("+")
	btree2 := binarytree.NewBinaryTree("*")
	btree3 := binarytree.NewBinaryTree("-")
	btree3.InsertLeft(btree1)
	btree3.InsertRight(btree2)
	if btree3.Size() != 3 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 3, btree3.Size())
	}
	if btree3.Height() != 1 {
		t.Errorf("Error la altura deberia dar %v, pero dio %v", 1, btree3.Height())
	}
}
func TestSumaNodos(t *testing.T) {
	btree := ejercicios.NewBinaryTree(2)
	btree2 := ejercicios.NewBinaryTree(3)
	btree3 := ejercicios.NewBinaryTree(4)
	btree4 := ejercicios.NewBinaryTree(10)
	btree.InsertLeft(btree2)
	btree.InsertRight(btree3)
	btree4.InsertLeft(btree)
	btree4.InsertRight(btree)
	if btree4.SumaNodos() != 28 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 28, btree4.SumaNodos())
	}
}
func TestSumaHojas(t *testing.T) {
	btree := ejercicios.NewBinaryTree(2)
	btree2 := ejercicios.NewBinaryTree(3)
	btree3 := ejercicios.NewBinaryTree(4)
	btree.InsertLeft(btree2)
	btree.InsertRight(btree3)
	//btree4 := ejercicios.NewBinaryTree(10)
	//btree4.InsertLeft(btree)
	//btree4.InsertRight(btree)
	if btree.SumaHojas() != 7 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 7, btree.SumaHojas())
	}
}
func TestSumarNodosInternos(t *testing.T) {

	btree := ejercicios.NewBinaryTree(2)
	btree2 := ejercicios.NewBinaryTree(3)
	btree3 := ejercicios.NewBinaryTree(4)
	btree4 := ejercicios.NewBinaryTree(10)
	btree.InsertLeft(btree2)
	btree.InsertRight(btree3)
	btree4.InsertLeft(btree)
	btree4.InsertRight(btree)
	if btree4.SumaNodosInternos() != 4 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 4, btree4.SumaNodosInternos())
	}
}
func TestSumarNodosPares(t *testing.T) {
	btree := ejercicios.NewBinaryTree(2)
	btree2 := ejercicios.NewBinaryTree(5)
	btree3 := ejercicios.NewBinaryTree(4)
	//btree4 := ejercicios.NewBinaryTree(10)
	btree.InsertLeft(btree2)
	btree.InsertRight(btree3)
	//btree4.InsertLeft(btree)
	//btree4.InsertRight(btree)
	if btree.SumaNodosPares() != 6 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 6, btree.SumaNodosPares())

	}
}
func TestSumaNodosMayorIgualQueSeis(t *testing.T) {
	btree := ejercicios.NewBinaryTree(2)
	btree2 := ejercicios.NewBinaryTree(6)
	btree3 := ejercicios.NewBinaryTree(8)
	btree.InsertLeft(btree2)
	btree.InsertRight(btree3)
	if btree.SumaNodosMayorIgualQueSeis() != 14 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 14, btree.SumaNodosMayorIgualQueSeis())
	}
}
func TestCalcularAlturaMaxima(t *testing.T) {
	btree := ejercicios.NewBinaryTree(2)
	btree2 := ejercicios.NewBinaryTree(3)
	btree3 := ejercicios.NewBinaryTree(4)
	btree5 := ejercicios.NewBinaryTree(5)
	btree.InsertLeft(btree2)
	btree.InsertRight(btree3)
	btree3.InsertLeft(btree5)

	if btree.CalcularAlturaMaxima() != 2 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 2, btree.CalcularAlturaMaxima())
	}
}
func TestCalcularSumarHojasIzquierdas(t *testing.T) {
	btree := ejercicios.NewBinaryTree(2)
	btree2 := ejercicios.NewBinaryTree(3)
	//btree3 := ejercicios.NewBinaryTree(4)
	btree.InsertLeft(btree2)
	//btree4 := ejercicios.NewBinaryTree(10)
	//btree4.InsertLeft(btree)
	//btree4.InsertRight(btree)
	if btree.SumaHojasIzquierdas() != 3 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 3, btree.SumaHojasIzquierdas())
	}
}
