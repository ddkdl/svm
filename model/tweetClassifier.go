package model

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
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
	validationLabels  []float64
	numberofDocuments int
	falsePositives    int
	falseNegatives    int
	truePositives     int
	trueNegatives     int
	accuracy          float64
	hitRate           float64
	specificity       float64
	precision         float64
	f1Score           float64
}

// NewTweetClassifier is a thing
func NewTweetClassifier(kernel kernel.Kernel, C, tolerance float64) *TweetClassifier {
	classifier := new(TweetClassifier)

	classifier.svmModel = NewModel(kernel, C, tolerance)
	classifier.testData = mat.NewDense(0, 0, nil)

	classifier.testLabels = nil
	classifier.validationLabels = nil
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

func (clf *TweetClassifier) LoadValidationLabels(validatonLabelsFilename string) {
	file, e := os.Open(validatonLabelsFilename)
	if e != nil {
		panic(e)
	}

	clf.validationLabels = make([]float64, 0)

	first := true

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ','

	for {
		document, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if first {
			first = false
			continue
		}
		num, _ := strconv.Atoi(document[1])

		clf.validationLabels = append(clf.validationLabels, float64(num))
	}
}

func (clf *TweetClassifier) Train(maxPasses int) {
	clf.svmModel.Train(maxPasses)
}

func (clf *TweetClassifier) ClassifyTweets() {
	for i := 0; i < clf.numberofDocuments; i++ {
		clf.testLabels[i] = clf.svmModel.Classify(clf.testData.RowView(i))
	}
}

func (clf *TweetClassifier) StoreResults(labelsFilename, statsFilename string) {
	labelsFile, e := os.Create(labelsFilename)

	if e != nil {
		panic(e)
	}
	defer labelsFile.Close()

	for i := 0; i < clf.numberofDocuments; i++ {
		s := strconv.FormatInt(int64(clf.testLabels[i]), 16)
		labelsFile.WriteString(s)
		labelsFile.WriteString("\n")
	}

	statsFile, e := os.Create(statsFilename)

	if e != nil {
		panic(e)
	}
	defer statsFile.Close()

	statsFile.WriteString("TP: ")
	statsFile.WriteString(fmt.Sprintln(clf.truePositives))
	statsFile.WriteString("\n")

	statsFile.WriteString("TN: ")
	statsFile.WriteString(fmt.Sprintln(clf.trueNegatives))
	statsFile.WriteString("\n")

	statsFile.WriteString("FP: ")
	statsFile.WriteString(fmt.Sprintln(clf.falsePositives))
	statsFile.WriteString("\n")

	statsFile.WriteString("FN: ")
	statsFile.WriteString(fmt.Sprintln(clf.falseNegatives))
	statsFile.WriteString("\n")

	statsFile.WriteString("Accuracy: ")
	statsFile.WriteString(fmt.Sprintln(clf.accuracy))
	statsFile.WriteString("\n")

	statsFile.WriteString("Hit Rate: ")
	statsFile.WriteString(fmt.Sprintln(clf.hitRate))
	statsFile.WriteString("\n")

	statsFile.WriteString("Specificity: ")
	statsFile.WriteString(fmt.Sprintln(clf.specificity))
	statsFile.WriteString("\n")

	statsFile.WriteString("Precision: ")
	statsFile.WriteString(fmt.Sprintln(clf.precision))
	statsFile.WriteString("\n")

	statsFile.WriteString("F1-Score: ")
	statsFile.WriteString(fmt.Sprintln(clf.f1Score))
	statsFile.WriteString("\n")
}

func (clf *TweetClassifier) Validate() {
	for i := range clf.validationLabels {
		if clf.validationLabels[i] == -1 && clf.testLabels[i] == -1 {
			clf.trueNegatives++
		}
		if clf.validationLabels[i] == 1 && clf.testLabels[i] == -1 {
			clf.falseNegatives++
		}
		if clf.validationLabels[i] == -1 && clf.testLabels[i] == 1 {
			clf.falsePositives++
		}
		if clf.validationLabels[i] == 1 && clf.testLabels[i] == 1 {
			clf.truePositives++
		}
	}

	clf.calculateAccuracy()
	clf.calculateHitRate()
	clf.calculateSpecificity()
	clf.calculatePrecision()
	clf.calculateF1Score()

}

func (clf *TweetClassifier) calculateAccuracy() {
	numerator := float64(clf.truePositives + clf.trueNegatives)
	denominator := float64(clf.trueNegatives + clf.truePositives + clf.falseNegatives + clf.falsePositives)

	clf.accuracy = numerator / denominator
}

func (clf *TweetClassifier) calculateHitRate() {
	numerator := float64(clf.truePositives)
	denominator := float64(clf.truePositives + clf.falseNegatives)

	clf.hitRate = numerator / denominator
}

func (clf *TweetClassifier) calculateSpecificity() {
	numerator := float64(clf.trueNegatives)
	denominator := float64(clf.trueNegatives + clf.falsePositives)

	clf.specificity = numerator / denominator
}

func (clf *TweetClassifier) calculatePrecision() {
	numerator := float64(clf.truePositives)
	denominator := float64(clf.truePositives + clf.falsePositives)

	clf.precision = numerator / denominator
}

func (clf *TweetClassifier) calculateF1Score() {
	numerator := float64(2 * clf.truePositives)
	denominator := float64(2*clf.truePositives + clf.falseNegatives + clf.falsePositives)

	clf.f1Score = numerator / denominator
}
