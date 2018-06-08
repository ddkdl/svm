package kernel

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

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
	degree int64
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

func NewPolynomialKernel(degree int64) Kernel {
	return &PolynomialKernel{degree: degree}
}

func NewRBFKernel(sigma float64) Kernel {
	return &RBFKernel{sigma: sigma}
}

// Evaluate is gooder
func (lk LinearKernel) Evaluate(X mat.Vector, Y mat.Vector) float64 {
	return mat.Dot(X, Y)
}

// Evaluate is gooder
func (pk PolynomialKernel) Evaluate(X mat.Vector, Y mat.Vector) float64 {
	result := float64(1)

	for i := 1; i <= int(pk.degree); i++ {
		result *= 1 + mat.Dot(X, Y)
	}

	return result
}

// Evaluate is gooder
func (rbfk RBFKernel) Evaluate(X mat.Vector, Y mat.Vector) float64 {
	result := mat.NewVecDense(X.Len(), nil)

	result.SubVec(X, Y)

	numerator := mat.Norm(result, 2)

	expression := (numerator * numerator) / (2 * rbfk.sigma * rbfk.sigma)
	finalResult := math.Exp(expression)

	return finalResult
}
