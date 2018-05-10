package main

import (
	"fmt"
	"github.com/ddkdl/svm/kernel"
	"github.com/ddkdl/svm/preprocessor"
)

func main() {
	var kernel kernel.LinearKernel

	fmt.Println(kernel.Evaluate(1,2))
	fmt.Println("I am work!")
	preprocessor.Parser("Asthma_Sample.csv")
}