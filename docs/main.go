package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/deadcheat/gegegengogogo"
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
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		from, err := time.ParseInLocation("2006/1/2", records[i][3], loc)
		to, err := time.ParseInLocation("2006/1/2", records[i][4], loc)

		log.Println(from.String(), to.String())
		g := gegegengogogo.Gengo{
			C:    gegegengogogo.GengoCodeFromString(records[i][0]),
			Name: records[i][1],
			Kana: records[i][2],
			From: from,
			To:   to,
		}
		d, _ := json.Marshal(g)
		log.Println(string(d))
	}
}
