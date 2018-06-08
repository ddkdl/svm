package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ddkdl/svm/kernel"
	"github.com/ddkdl/svm/model"
	"github.com/ddkdl/svm/preprocessor"
)

func main() {
	var kernelType string
	var C float64
	var maxTries int64
	var tolerance float64
	var polynomialDegree int64
	var gamma float64
	var trainFilename string
	var testFilename string
	var validationLabelsFilename string
	var outputFilename string

	var classifier *model.TweetClassifier

	argument := os.Args[1:]
	numberOfArguments := len(argument)

	if numberOfArguments > 8 || numberOfArguments < 7 {
		fmt.Println("Insufficient arguments")
		usage()
		return
	}

	kernelType = argument[0]

	switch kernelType {
	case "-l":
		C, _ = strconv.ParseFloat(argument[1], 64)
		maxTries, _ = strconv.ParseInt(argument[2], 10, 64)
		trainFilename = argument[3]
		testFilename = argument[4]
		validationLabelsFilename = argument[5]
		outputFilename = argument[6]
	case "-p":
		C, _ = strconv.ParseFloat(argument[1], 64)
		polynomialDegree, _ = strconv.ParseInt(argument[2], 10, 64)
		maxTries, _ = strconv.ParseInt(argument[3], 10, 16)
		trainFilename = argument[4]
		testFilename = argument[5]
		validationLabelsFilename = argument[6]
		outputFilename = argument[7]
	case "-r":
		C, _ = strconv.ParseFloat(argument[1], 64)
		gamma, _ = strconv.ParseFloat(argument[2], 64)
		maxTries, _ = strconv.ParseInt(argument[3], 10, 64)
		trainFilename = argument[4]
		testFilename = argument[5]
		validationLabelsFilename = argument[6]
		outputFilename = argument[7]
	}

	tokenizer := preprocessor.NewTokenizer()
	tweets := preprocessor.ParseText(trainFilename)
	unformattedLabels := preprocessor.ParseLabel(trainFilename)
	tokenizedTweets := tokenizer.TokenizeTweets(tweets)
	dtm := tokenizer.CreateDocumentTermMatrix(tokenizedTweets)
	labels := preprocessor.CreateLabelVector(unformattedLabels)

	testSet := preprocessor.ParseText(testFilename)
	tokenizedTestSet := tokenizer.TokenizeTweets(testSet)
	testDTM := tokenizer.CreateDocumentTermMatrix(tokenizedTestSet)

	switch kernelType {
	case "-l":
		classifier = model.NewTweetClassifier(kernel.NewLinearKernel(), C, tolerance)
	case "-p":
		classifier = model.NewTweetClassifier(kernel.NewPolynomialKernel(polynomialDegree), C, tolerance)
	case "-r":
		classifier = model.NewTweetClassifier(kernel.NewRBFKernel(gamma), C, tolerance)
	}

	classifier.LoadTrainingSet(dtm, labels)
	classifier.LoadTestSet(testDTM)
	classifier.LoadValidationLabels(validationLabelsFilename)
	classifier.Train(int(maxTries))
	classifier.ClassifyTweets()
	classifier.Validate()
	classifier.StoreResults("labels_"+outputFilename, "stats_"+outputFilename)

}

func usage() {
	fmt.Println("svm usage:")
	fmt.Println("\tLinear Kernel")
	fmt.Println("\tsvm -l [C value] [maxTries] [trainFile] [testFile] [validationLabels] [outputFile]")
	fmt.Println("\tPolynomial Kernel")
	fmt.Println("\tsvm -p [C value] [polynomialDegree] [maxTries] [trainFile] [testFile] [validationLabels] [outputFile]")
	fmt.Println("\tRBF Kernel")
	fmt.Println("\tsvm -r [C value] [gamma] [maxTries] [trainFile] [testFile] [validationLabels] [outputFile]")
}
