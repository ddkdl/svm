package preprocessor

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func Parser(filename string) []string {
	var tweetList []string

	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'

	for {
		document, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		tweetList = append(tweetList, document[2])
	}

	return tweetList
}
