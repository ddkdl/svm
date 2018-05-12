package preprocessor

import (
	"regexp"
	"strings"
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
