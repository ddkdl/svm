package model

import "gonum.org/v1/gonum/mat"

func CalculateGramMatrix(X *mat.Dense) *mat.Dense {
	rows, _ := X.Dims()
	var res []float64

	for i := 0; i < rows; i++ {
		x1 := X.RowView(i)
		for j := 0; j < rows; j++ {
			x2 := X.RowView(j)

			res = append(res, mat.Dot(x1, x2))
		}
	}

	return mat.NewDense(rows, rows, res)
}

func CalculateYOuterProduct(y *mat.VecDense) *mat.Dense {
	size := y.Len()

	res := mat.NewDense(size, size, nil)
	res.Outer(1, y, y)
	return res
}

func CalculateD(K, yOuter *mat.Dense) *mat.Dense {
	rows, cols := K.Dims()
	res := mat.NewDense(rows, cols, nil)
	res.Product(K, yOuter)
	return res
}
