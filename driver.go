package main

import (
	"fmt"

	"github.com/ddkdl/svm/preprocessor"
)

func main() {

	tweetList := preprocessor.Parser("Asthma_Sample.csv")
	tokens := preprocessor.TweetTokenizer(tweetList)

	for _, element := range tokens {
		fmt.Println(element)
	}
}
