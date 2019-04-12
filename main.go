package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Question struct {
	Text string
	Pass string
	Fail string
}

func main() {

	questionsFilename := "questions.csv"
	questionsFile, err := os.Open(questionsFilename)
	defer questionsFile.Close()

	if err != nil {
		log.Fatalf("problem opening %s %+v", questionsFilename, err)
	}

	r := csv.NewReader(questionsFile)

	var questions []Question
	for {
		row, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("problem reading csv %+v", err)
		}

		questions = append(questions, Question{
			row[0],
			row[1],
			row[2],
		})
	}

	fmt.Println("Answer Y/N to the following")

	var responses []string

	for _, q := range questions {
		fmt.Println(q.Text)

		var reply string
		fmt.Scanln(&reply)
		reply = strings.ToLower(reply)

		if reply == "y" {
			responses = append(responses, fmt.Sprintf("✔️ %s", q.Pass))
		} else {
			responses = append(responses, fmt.Sprintf("⚠️, %s",q.Fail))
		}
	}

	fmt.Println()

	for _, advice := range responses {
		fmt.Println(advice)
	}
}
