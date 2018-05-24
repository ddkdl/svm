package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	"github.com/ddkdl/svm/model"
	"github.com/ddkdl/svm/preprocessor"
)

func main() {
	tweets := [][]string{
		{"arroz", "feijao", "quiabo"},
		{"feijao", "arroz", "batata", "quiabo"},
		{"matabala", "feijao", "batata"},
		{"quiabo", "couve", "matabala"},
		{"arroz", "tomate", "tomate"},
	}

	X := preprocessor.CreateDTM(tweets)
	fmt.Println("Printing X")
	printMatrix(X)

	K := model.CalculateGramMatrix(X)
	fmt.Println("Printing K")
	printMatrix(K)

	y := mat.NewVecDense(5, []float64{1, 1, 0, 0, 1})

	yOuter := model.CalculateYOuterProduct(y)
	fmt.Println("Printing yOuter")
	printMatrix(yOuter)

	res := model.CalculateD(K, yOuter)
	fmt.Println("Printing D")
	printMatrix(res)

}

func printMatrix(a *mat.Dense) {
	rows, _ := a.Dims()

	for i := 0; i < rows; i++ {
		fmt.Println(a.RowView(i))
	}
}
