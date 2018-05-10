package model 

import "github.com/ddkdl/svm/kernel"

type Model struct {
	kernel kernel.Kernel
	C int
	w []float64
	b float64
}

func (m *Model) Train(){

}

func (m *Model) Classify(){

}