package preprocessor

import "strings"

func Tokenize(document string) {
	tokens := strings.Fields(document)

	for _, element := range tokens {
		if stopwords[element] == true {
			continue
		}
	}
}
