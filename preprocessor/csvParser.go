package preprocessor

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"fmt"
)

func Parser(filename string) {
	// var features []string

	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'

	for {
		document, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Println(document[2])
	}
}