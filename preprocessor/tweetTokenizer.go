// Package preprocessor is a good package. Very good indeed.
package preprocessor

import (
	"regexp"
	"strings"

	"gonum.org/v1/gonum/mat"
)

// TweetTokenizer is a function that operates on things.
func TweetTokenizer(listOfTweets []string) [][]string {
	var listOfTokens [][]string

	for _, tweet := range listOfTweets {
		tokens := tokenize(tweet)
		listOfTokens = append(listOfTokens, tokens)
	}

	return listOfTokens
}

func tokenize(sentence string) []string {
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

func createFeatureList(tokenizedTweets [][]string) map[string]int {
	featureList := make(map[string]int)
	featureIndex := 0

	for _, tweet := range tokenizedTweets {
		for _, word := range tweet {
			if _, isInList := featureList[word]; !isInList {
				featureList[word] = featureIndex
				featureIndex++
			}
		}
	}

	return featureList
}

func createDocumentTermMatrix(tweetList [][]string) *mat.Dense {
	featureList := createFeatureList(tweetList)

	tempDTM := make([][]int, len(tweetList))
	var DTM []float64

	for i := range tempDTM {
		tempDTM[i] = make([]int, len(featureList))
	}

	for i, tweet := range tweetList {
		for _, word := range tweet {
			tempDTM[i][featureList[word]]++
		}
	}

	for i := range tempDTM {
		for j := range tempDTM[i] {
			DTM = append(DTM, float64(tempDTM[i][j]))
		}
	}

	return mat.NewDense(len(tweetList), len(featureList), DTM)
}
