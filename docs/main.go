package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/deadcheat/gengou"
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
	gengos := make([]gengou.Gengo, 0)
	for i := range records {
		from := parseDate(records[i][3])
		to := parseDate(records[i][4])
		g := gengou.Gengo{
			C:    gengou.GengoCodeFromString(records[i][0]),
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

func parseDate(str string) time.Time {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	s := strings.SplitN(str, "/", 3)
	year, _ := strconv.Atoi(s[0])
	month, _ := strconv.Atoi(s[1])
	day, _ := strconv.Atoi(s[2])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
}
