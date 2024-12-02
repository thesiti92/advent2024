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

	"github.com/davecgh/go-spew/spew"
)

//go:embed input.csv
var input []byte

func main() {
	reader := csv.NewReader(bytes.NewReader(input))
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true
	total := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if checkRecord(record) {
			total += 1
			continue
		}
		for i := range record {

			cur := slices.Delete(slices.Clone(record), i, i+1)
			if checkRecord(cur) {
				total += 1
				break
			}
		}
	}

	fmt.Println("Answer:")
	fmt.Println(total)
}

func checkRecord(record []string) bool {
	cur := -1
	var status string
	for _, e := range record {
		val, _ := strconv.Atoi(e)
		if cur < 0 {
			cur = val
			continue
		}
		if status == "" {
			if val > cur {
				status = "increasing"
			} else {
				status = "decreasing"
			}
		}
		if status == "increasing" {
			if val-cur > 3 || val-cur <= 0 {
				spew.Dump("invalid")
				spew.Dump(record)

				return false
			}
		}
		if status == "decreasing" {
			if cur-val > 3 || cur-val <= 0 {
				spew.Dump("invalid")
				spew.Dump(record)

				return false
			}
		}

		cur = val
	}
	return true
}
