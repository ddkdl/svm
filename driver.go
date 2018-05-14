package main

import (
	"fmt"

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

	fmt.Println(preprocessor.CreateDTM(tweets))
}
