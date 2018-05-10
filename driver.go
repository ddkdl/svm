package main

import (
	"fmt"
	"github.com/ddkdl/svm/kernel"
)

func main() {
	var kernel kernel.LinearKernel

	fmt.Println(kernel.Evaluate(1,2))
}