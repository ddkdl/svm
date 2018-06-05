package preprocessor

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// ParseText is a good funk
func ParseText(filename string) []string {
	var tweetList []string
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
	}

	return tweetList
}

// ParseLabel tem label vermelho
func ParseLabel(filename string) []float64 {
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

		label, _ := strconv.ParseFloat(document[12], 64)
		labelList = append(labelList, label)

	}

	return labelList
}
