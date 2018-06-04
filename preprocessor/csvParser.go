package preprocessor

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// Parser is a good funk
func Parser(filename string) ([]string, []float64) {
	var tweetList []string
	var labelList []float64
	first := true

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
		if first {
			first = false
			continue
		}

		tweetList = append(tweetList, document[2])
		label, _ := strconv.ParseFloat(document[12], 64)
		labelList = append(labelList, label)
	}

	return tweetList, labelList
}
