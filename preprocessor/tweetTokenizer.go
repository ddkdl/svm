// Package preprocessor is a good package. Very good indeed.
package preprocessor

import (
	"regexp"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type TweetTokenizer struct {
	featureList    map[string]int
	hasFeatureList bool
}

func NewTokenizer() *TweetTokenizer {
	tokenizer := new(TweetTokenizer)
	tokenizer.featureList = make(map[string]int)
	tokenizer.hasFeatureList = false

	return tokenizer
}

// TweetTokenizer is a function that operates on things.
func (tokenizer *TweetTokenizer) TokenizeTweets(listOfTweets []string) [][]string {
	var listOfTokens [][]string

	for _, tweet := range listOfTweets {
		tokens := tokenizer.tokenize(tweet)
		listOfTokens = append(listOfTokens, tokens)
	}

	return listOfTokens
}

func (tokenizer *TweetTokenizer) tokenize(sentence string) []string {
	var tokens []string
	listOfWords := strings.Fields(sentence)

	for _, word := range listOfWords {

		word = strings.ToLower(word)

		if isStopword(word) {
			continue
		}

		if hasHyperlink(word) {
			continue
		}

		if hasHashtag(word) {
			continue
		}

		if hasMention(word) {
			continue
		}

		word = removeSpecialCharacters(word)

		tokens = append(tokens, word)
	}

	return tokens
}

func isStopword(word string) bool {
	return stopwords[word]
}

func hasHyperlink(word string) bool {
	var hastagPattern = regexp.MustCompile(`https*://\w`)

	return hastagPattern.MatchString(word)
}

func hasHashtag(word string) bool {
	var hastagPattern = regexp.MustCompile(`#\w`)

	return hastagPattern.MatchString(word)
}

func hasMention(word string) bool {
	var hastagPattern = regexp.MustCompile(`@\w`)

	return hastagPattern.MatchString(word)
}

func removeSpecialCharacters(word string) string {
	var punctuationPattern = regexp.MustCompile(`\W*`)

	return punctuationPattern.ReplaceAllLiteralString(word, ``)
}

func (tokenizer *TweetTokenizer) createFeatureList(tokenizedTweets [][]string) {
	featureIndex := 0

	for _, tweet := range tokenizedTweets {
		for _, word := range tweet {
			if _, isInList := tokenizer.featureList[word]; !isInList {
				tokenizer.featureList[word] = featureIndex
				featureIndex++
			}
		}
	}
}

// CreateDocumentTermMatrix be good like Vera
func (tokenizer *TweetTokenizer) CreateDocumentTermMatrix(tweetList [][]string) *mat.Dense {
	if !tokenizer.hasFeatureList {
		tokenizer.createFeatureList(tweetList)
		tokenizer.hasFeatureList = true
	}

	tempDTM := make([][]int, len(tweetList))
	var DTM []float64

	for i := range tempDTM {
		tempDTM[i] = make([]int, len(tokenizer.featureList))
	}

	for i, tweet := range tweetList {
		for _, word := range tweet {
			tempDTM[i][tokenizer.featureList[word]]++
		}
	}

	for i := range tempDTM {
		for j := range tempDTM[i] {
			DTM = append(DTM, float64(tempDTM[i][j]))
		}
	}

	return mat.NewDense(len(tweetList), len(tokenizer.featureList), DTM)
}

// CreateLabelVector is a function I am being forced to comment cuz reasons
func CreateLabelVector(labelList []float64) mat.Vector {
	return mat.NewVecDense(len(labelList), labelList)
}
