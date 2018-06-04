package main

import (
	"fmt"

	"github.com/ddkdl/svm/kernel"

	"github.com/ddkdl/svm/model"
	"github.com/ddkdl/svm/preprocessor"
	"gonum.org/v1/gonum/mat"
)

func main() {
	tweets, labels := preprocessor.Parser("Asthma_Sample.csv")
	tokenizedTweets := preprocessor.TweetTokenizer(tweets)
	dtm := preprocessor.CreateDocumentTermMatrix(tokenizedTweets)
	realLabels := preprocessor.CreateLabelVector(labels)

	svmModel := model.NewModel(kernel.NewLinearKernel(), 1.0, 0.001)

	svmModel.LoadTrainingSet(dtm, realLabels)
	// fmt.Println(svmModel.Y)
	// fmt.Println(svmModel.X)
	svmModel.Train(5)

	fmt.Println(svmModel)
}

func printMatrix(a *mat.Dense) {
	rows, _ := a.Dims()

	for i := 0; i < rows; i++ {
		fmt.Println(a.RowView(i))
	}
}
