package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const TOTAL_TIME int = 10

type QA struct {
	Question string
	Answer   string
}

func readCSV(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parseCSV(data []byte) ([]QA, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	var qaData []QA
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error in the processing CSV: ", err)
			break
		}
		if len(record) == 2 {
			qaData = append(qaData, QA{
				Question: record[0],
				Answer:   record[1],
			})
		}
	}
	return qaData, nil
}

func suffleData(data []QA) []QA {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
	return data
}

func quiz(data []QA) {
	data = suffleData(data)
	right, total := 0, len(data)

	timer := time.NewTimer(time.Duration(TOTAL_TIME) * time.Second)
	defer timer.Stop()

	go func() {
		countdown := TOTAL_TIME
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				countdown--
				if countdown <= 0 {
					return
				}
				fmt.Println("Time: ", countdown)
			}
		}
	}()

	for i := 0; i < total; i++ {
		fmt.Printf("%d What is %v, sir?\n", i+1, data[i].Question)

		ansChan := make(chan string)
		go func() {
			var userAns string
			fmt.Scanln(&userAns)
			ansChan <- userAns
		}()
		select {
		case ans := <-ansChan:
			if ans == data[i].Answer {
				right++
			}
		case <-timer.C:
			fmt.Println("Time's UP!")
			fmt.Printf("Score: %d/%d\n", right, total)
			return
		}
	}
	fmt.Printf("Oh man you got it %v/%v\n", right, total)
}

func main() {
	var path string = "./problems.csv"
	data, err := readCSV(path)
	if err != nil {
		fmt.Println("Error in the reading file: ", err)
	}
	qaData, err := parseCSV(data)
	quiz(qaData)
}
