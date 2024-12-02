package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
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
			if val == cur {
				total += 1
				break
			}
			if status == "increasing" {
				if val-cur > 3 || val-cur <= 0 {
					spew.Dump("invalid")
					spew.Dump(record)

					total += 1
					break
				}
			}
			if status == "decreasing" {
				if cur-val > 3 || cur-val <= 0 {
					spew.Dump("invalid")
					spew.Dump(record)

					total += 1
					break
				}
			}

			cur = val
		}
	}

	fmt.Println("Answer:")
	fmt.Println(1000 - total)
}
