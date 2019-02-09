package gengou

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ErrStringIsNotSlashed pre-defined error that will be returned when given string doesn't satisfy y/m/d format
var ErrStringIsNotSatisfiedFormat = errors.New("given string does not satisfy y/m/d format")

// DefaultTime pre-defined default time.Date value
var DefaultTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)

// ParseSlashedYMD will parse y/m/d(in Go, 2006/1/2) style date string
func ParseSlashedYMD(str string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return DefaultTime, err
	}
	s := strings.Split(str, "/")
	if len(s) < 3 {
		return DefaultTime, ErrStringIsNotSatisfiedFormat
	}

	var year, month, day int
	year, err = strconv.Atoi(s[0])
	if err != nil {
		return DefaultTime, ErrStringIsNotSatisfiedFormat
	}
	month, err = strconv.Atoi(s[1])
	if err != nil {
		return DefaultTime, ErrStringIsNotSatisfiedFormat
	}
	day, err = strconv.Atoi(s[2])
	if err != nil {
		return DefaultTime, ErrStringIsNotSatisfiedFormat
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc), nil
}
