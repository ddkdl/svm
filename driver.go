package main

import (
	"github.com/ddkdl/svm/kernel"
	"github.com/ddkdl/svm/model"
	"github.com/ddkdl/svm/preprocessor"
)

func main() {
	tokenizer := preprocessor.NewTokenizer()
	tweets := preprocessor.ParseText("Cancer_Sample.csv")
	labels := preprocessor.ParseLabel("Cancer_Sample.csv")
	tokenizedTweets := tokenizer.TokenizeTweets(tweets)
	dtm := tokenizer.CreateDocumentTermMatrix(tokenizedTweets)
	realLabels := preprocessor.CreateLabelVector(labels)

	testSet := preprocessor.ParseText("Cancer_test.csv")
	tokenizedTestSet := tokenizer.TokenizeTweets(testSet)
	testDTM := tokenizer.CreateDocumentTermMatrix(tokenizedTestSet)

	classifier := model.NewTweetClassifier(kernel.NewRBFKernel(0.5), 10, 0.001)

	classifier.LoadTrainingSet(dtm, realLabels)
	classifier.LoadTestSet(testDTM)
	classifier.LoadValidationLabels("Cancer_Test_Results.csv")
	classifier.Train(5)
	classifier.ClassifyTweets()
	classifier.Validate()
	classifier.StoreResults("labels_after_classification.txt", "stats_after_classification.txt")

}
