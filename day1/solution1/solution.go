package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"slices"
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
	rhs := make([]int, 0, 1000)
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
		rhs = append(rhs, r)
	}
	slices.Sort(lhs)
	slices.Sort(rhs)
	totalDiff := 0
	for i := range len(lhs) {
		diff := lhs[i] - rhs[i]
		if diff < 0 {
			diff = diff * -1
		}
		totalDiff += diff
	}
	fmt.Println("Answer:")
	fmt.Println(totalDiff)
}
