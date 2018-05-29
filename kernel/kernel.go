package kernel

type Kernel interface {
	Evaluate(X float64, Y float64) float64
}

type LinearKernel struct {
	Kernel
}

type PolynomialKernel struct {
	Kernel
	degree int
}

type RBFKernel struct {
	Kernel
	sigma float64
}

func (lk *LinearKernel) Evaluate(X float64, Y float64) float64 {
	return X * Y
}

func (pk *PolynomialKernel) Evaluate(X float64, Y float64) float64 {
	result := float64(1)

	for i := 1; i <= pk.degree; i++ {
		result *= 1 + X*Y
	}

	return result
}

func (rbfk *RBFKernel) Evaluate(X float64, Y float64) float64 {
	return 1
}
