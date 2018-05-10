package kernel

type Kernel interface {
	Evaluate(X float32, Y float32) float32
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
	sigma float32
}

func (lk *LinearKernel) Evaluate(X float32, Y float32) float32 {
	return X * Y
}

func (pk *PolynomialKernel) Evaluate(X float32, Y float32) float32 {
	result := float32(1)

	for i := 1; i <= pk.degree; i++ {
		result *= 1 + X * Y
	}

	return result
}

func (rbfk *RBFKernel) Evaluate(X float32, Y float32) float32 {
	return 1
}