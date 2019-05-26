package gengou

import (
	"errors"
	"time"
)

// Gengou is an entity of 元号
type Gengou struct {
	C    GengouCode `json:"-"`
	Name string     `json:"name"`
	Kana []string   `json:"kana"`
	From time.Time  `json:"from"`
	To   time.Time  `json:"to"`
}

// NoSuchGengouFoundError will be returned when Find() given y/m/d didn't hit any gengou
var NoSuchGengouFoundError = errors.New("no such gengou found")

// Find searches gengo from given yyyy/mm/dd string
func Find(date string) (result []Gengou, err error) {
	var t time.Time
	t, err = ParseSlashedYMD(date)
	if err != nil {
		return result, err
	}

	if (t.Equal(nanbokueraFrom) || t.After(nanbokueraFrom)) &&
		(t.Equal(nanbokueraTo) || t.Before(nanbokueraTo)) {
		result = searchFromNanbokuchoEra(t)
	} else {
		result = searchFromOtherEra(t)
	}

	if len(result) == 0 {
		return result, NoSuchGengouFoundError
	}
	return result, nil
}

func searchFromNanbokuchoEra(t time.Time) []Gengou {

	result := make([]Gengou, 0)
	for _, s := range gengoDataNanbokuEra {
		if s.To.After(t) && (s.From.Before(t) || s.From.Equal(t)) {
			result = append(result, s)
		}
	}

	return result
}

func searchFromOtherEra(t time.Time) []Gengou {
	result := make([]Gengou, 0)
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
