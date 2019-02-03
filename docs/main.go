package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("list.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Comma = '\t'
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := range records {
		log.Println(strings.Join(records[i], ","))
	}
}
