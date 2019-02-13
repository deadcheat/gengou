package gengou

import (
	"fmt"
	"time"
)

// Gengo is an entity of 元号
type Gengo struct {
	C    GengoCode `json:"code"`
	Name string    `json:"name"`
	Kana string    `json:"kana"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

// SearchGengo searches gengo from given yyyy/mm/dd string
func SearchGengo(date string) ([]Gengo, error) {
	result := make([]Gengo, 0)
	t, err := ParseSlashedYMD(date)
	if err != nil {
		return result, err
	}

	min := 0
	max := len(gengoData)
	for {
		si := (min + max) / 2
		s := gengoData[si]
		fmt.Println(min, max, si, s.Name, t, s.From, s.To)
		if s.To.After(t) && (s.From.Before(t) || s.From.Equal(t)) {
			result = append(result, s)
		}

		if (max - min) <= 1 {
			break
		}

		if s.To.After(t) {
			max = si
			continue
		}
		min = si
	}

	return result, nil
}
