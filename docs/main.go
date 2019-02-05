package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	gengos := make([]gegegengogogo.Gengo, 0)
	for i := range records {
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		from, err := time.ParseInLocation("2006/1/2", records[i][3], loc)
		to, err := time.ParseInLocation("2006/1/2", records[i][4], loc)
		g := gegegengogogo.Gengo{
			C:    gegegengogogo.GengoCodeFromString(records[i][0]),
			Name: records[i][1],
			Kana: records[i][2],
			From: from,
			To:   to,
		}
		gengos = append(gengos, g)
	}
	d, _ := json.Marshal(gengos)
	fmt.Println(string(d))
}
