package kernel

import "gonum.org/v1/gonum/mat"

// Kernel is good
type Kernel interface {
	Evaluate(X mat.Vector, Y mat.Vector) float64
}

// LinearKernel is also good
type LinearKernel struct {
	Kernel
}

// PolynomialKernel is decent
type PolynomialKernel struct {
	Kernel
	degree int
}

// RBFKernel is best kernel
type RBFKernel struct {
	Kernel
	sigma float64
}

// NewLinearKernel is a fnc
func NewLinearKernel() Kernel {
	return &LinearKernel{}
}

func NewPolynomialKernel(degree int) Kernel {
	return &PolynomialKernel{degree: degree}
}

// Evaluate is gooder
func (lk LinearKernel) Evaluate(X mat.Vector, Y mat.Vector) float64 {
	return mat.Dot(X, Y)
}

// Evaluate is gooder
func (pk PolynomialKernel) Evaluate(X mat.Vector, Y mat.Vector) float64 {
	result := float64(1)

	for i := 1; i <= pk.degree; i++ {
		result *= 1 + mat.Dot(X, Y)
	}

	return result
}

// Evaluate is gooder
func (rbfk RBFKernel) Evaluate(X float64, Y float64) float64 {
	return 1
}
