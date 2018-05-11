package preprocessor

import "strings"

func Tokenize(document string) []string {
	var tokens []string
	words := strings.Fields(document)

	for _, element := range words {

		element = strings.ToLower(element)
		if stopwords[element] == true {
			continue
		}
		tokens = append(tokens, element)
	}

	return tokens
}

// TO DO: Remove punctuation, remove hashtags and mentions,
// remove hyperlinks, and remove extra spaces
func TweetTokenizer(listOfTweets []string) [][]string {
	var tokens [][]string

	for _, element := range listOfTweets {
		token := Tokenize(element)
		tokens = append(tokens, token)
	}

	return tokens
}
