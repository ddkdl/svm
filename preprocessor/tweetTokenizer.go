package preprocessor

import (
	"regexp"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func HasHashtag(word string) bool {
	var hastagPattern = regexp.MustCompile(`#\w`)

	return hastagPattern.MatchString(word)
}

func HasMention(word string) bool {
	var hastagPattern = regexp.MustCompile(`@\w`)

	return hastagPattern.MatchString(word)
}

func HasHyperlink(word string) bool {
	var hastagPattern = regexp.MustCompile(`https*://\w`)

	return hastagPattern.MatchString(word)
}

func RemoveSpecialCharacters(word string) string {
	var punctuationPattern = regexp.MustCompile(`\W*`)

	return punctuationPattern.ReplaceAllLiteralString(word, ``)
}

func Tokenize(document string) []string {
	var tokens []string
	words := strings.Fields(document)

	for _, element := range words {

		element = strings.ToLower(element)
		if stopwords[element] {
			continue
		}
		if HasHyperlink(element) {
			continue
		}
		if HasHashtag(element) {
			continue
		}
		if HasMention(element) {
			continue
		}
		element = RemoveSpecialCharacters(element)
		tokens = append(tokens, element)
	}

	return tokens
}

func TweetTokenizer(listOfTweets []string) [][]string {
	var tokens [][]string

	for _, element := range listOfTweets {
		token := Tokenize(element)
		tokens = append(tokens, token)
	}

	return tokens
}

func FeatureList(tweets [][]string) map[string]int {
	featureList := make(map[string]int)
	counter := 0

	for _, tweet := range tweets {
		for _, word := range tweet {
			if _, isInList := featureList[word]; !isInList {
				featureList[word] = counter
				counter++
			}
		}
	}

	return featureList
}

func CreateDTM(tweets [][]string) *mat.Dense {
	featureList := FeatureList(tweets)

	temp_dtm := make([][]int, len(tweets))
	var dtm []float64

	for i, _ := range temp_dtm {
		temp_dtm[i] = make([]int, len(featureList))
	}

	for i, tweet := range tweets {
		for _, word := range tweet {
			temp_dtm[i][featureList[word]] += 1
		}
	}

	for i, _ := range temp_dtm {
		for j, _ := range temp_dtm[i] {
			dtm = append(dtm, float64(temp_dtm[i][j]))
		}
	}

	return mat.NewDense(len(tweets), len(featureList), dtm)
}
