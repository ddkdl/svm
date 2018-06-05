package model

import (
	"os"
	"strconv"

	"github.com/ddkdl/svm/kernel"
	"gonum.org/v1/gonum/mat"
)

// TweetClassifier is a good thing to have
type TweetClassifier struct {
	svmModel          *Model
	testData          *mat.Dense
	testLabels        []float64
	numberofDocuments int
	falsePositives    int
	falseNegatives    int
	truePositives     int
	trueNegatives     int
}

// NewTweetClassifier is a thing
func NewTweetClassifier(kernel kernel.Kernel, C, tolerance float64) *TweetClassifier {
	classifier := new(TweetClassifier)

	classifier.svmModel = NewModel(kernel, C, tolerance)
	classifier.testData = mat.NewDense(0, 0, nil)

	classifier.testLabels = nil
	classifier.numberofDocuments = 0

	classifier.falseNegatives = 0
	classifier.falsePositives = 0
	classifier.trueNegatives = 0
	classifier.truePositives = 0

	return classifier
}

// LoadTrainingSet is a thing
func (clf *TweetClassifier) LoadTrainingSet(documentTermMatrix *mat.Dense, trainingLabels mat.Vector) {
	clf.svmModel.loadTrainingSet(documentTermMatrix, trainingLabels)
}

func (clf *TweetClassifier) LoadTestSet(testSet *mat.Dense) {
	clf.testData = testSet
	clf.numberofDocuments, _ = testSet.Dims()
	clf.testLabels = make([]float64, clf.numberofDocuments)
}

func (clf *TweetClassifier) Train(maxPasses int) {
	clf.svmModel.Train(maxPasses)
}

func (clf *TweetClassifier) ClassifyTweets() {
	for i := 0; i < clf.numberofDocuments; i++ {
		clf.testLabels[i] = clf.svmModel.Classify(clf.testData.RowView(i))
	}
}

func (clf *TweetClassifier) StoreResults(labelsFilename string) {
	labelsFile, e := os.Create(labelsFilename)

	if e != nil {
		panic(e)
	}
	defer labelsFile.Close()

	for i := 0; i < clf.numberofDocuments; i++ {
		s := strconv.FormatInt(int64(clf.testLabels[i]), 16)
		labelsFile.WriteString(s)
	}
}
