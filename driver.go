package main

import (
	"fmt"

	"github.com/ddkdl/svm/preprocessor"
)

func main() {
	tweets := []string{"Hello", "http://sldfjdlfjs askf", "@holler this is the best", "#yo lo"}

	fmt.Println(preprocessor.TweetTokenizer(tweets))
}
