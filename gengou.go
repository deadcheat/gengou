package gengou

import (
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

// FindGengo searches gengo from given yyyy/mm/dd string
func FindGengo(date string) ([]Gengo, error) {
	t, err := ParseSlashedYMD(date)
	if err != nil {
		return []Gengo{}, err
	}

	if (t.Equal(nanbokueraFrom) || t.After(nanbokueraFrom)) &&
		(t.Equal(nanbokueraTo) || t.Before(nanbokueraTo)) {
		return searchFromNanbokuchoEra(t), nil
	}

	return searchFromOtherEra(t), nil
}

func searchFromNanbokuchoEra(t time.Time) []Gengo {

	result := make([]Gengo, 0)
	for _, s := range gengoDataNanbokuEra {
		if s.To.After(t) && (s.From.Before(t) || s.From.Equal(t)) {
			result = append(result, s)
		}
	}

	return result
}

func searchFromOtherEra(t time.Time) []Gengo {
	result := make([]Gengo, 0)
	min := 0
	max := len(gengoData)
	for {
		si := (min + max) / 2
		s := gengoData[si]
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

	return result
}
