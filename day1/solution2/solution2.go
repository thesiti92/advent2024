package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
)

//go:embed input.csv
var input []byte

func main() {
	reader := csv.NewReader(bytes.NewReader(input))
	reader.Comma = ' '
	reader.FieldsPerRecord = 2
	reader.TrimLeadingSpace = true
	lhs := make([]int, 0, 1000)
	rhs := make(map[int]int, 1000)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		l, _ := strconv.Atoi(record[0])
		r, _ := strconv.Atoi(record[1])
		lhs = append(lhs, l)
		rhs[r] += 1
	}
	totalScore := 0
	for _, n := range lhs {
		totalScore += n * rhs[n]
	}
	fmt.Println("Answer:")
	fmt.Println(totalScore)
}
