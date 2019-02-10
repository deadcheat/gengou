package gengou

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/deadcheat/gengou/mocks"
	"github.com/golang/mock/gomock"
)

func TestParseSlashedYMDReturnsDateSuccessfully(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Error("error occurred when time.LoadLocation")
	}

	testee := "645/03/02"
	expect := time.Date(645, 3, 2, 0, 0, 0, 0, loc)

	actual, err := ParseSlashedYMD(testee)
	if err != nil {
		t.Errorf("ParseSlashedYMD should not return any error but %+v", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(`ParseSlashedYMD returned unexpected value
	expected: %+v
	actual  : %+v
`, expect, actual)
	}
}

func TestParseSlashedYMDShouldReturnErrorWhenFailedLoadLocation(t *testing.T) {

	c := gomock.NewController(t)
	m := mocks.NewMockTimeAdapter(c)
	m.EXPECT().LoadLocation("Asia/Tokyo").Return(nil, errors.New("test error"))

	tmpTa := ta
	ta = m

	testee := "645/03/02"
	expect := DefaultTime

	actual, err := ParseSlashedYMD(testee)
	if err == nil {
		t.Errorf("ParseSlashedYMD should return error")
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(`ParseSlashedYMD returned unexpected value
	expected: %+v
	actual  : %+v
`, expect, actual)
	}

	ta = tmpTa
}

func TestParseSlashedYMDShouldReturnErrorWhenIllegalYMDFormat(t *testing.T) {

	expect := DefaultTime

	// testee contains less than 2 slashes
	testee := "645/03"
	actual, err := ParseSlashedYMD(testee)
	if err == nil {
		t.Errorf("ParseSlashedYMD should return error")
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(`ParseSlashedYMD returned unexpected value
	expected: %+v
	actual  : %+v
`, expect, actual)
	}

	// year is not number
	testee = "abcd/03/02"
	actual, err = ParseSlashedYMD(testee)
	if err == nil {
		t.Errorf("ParseSlashedYMD should return error")
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(`ParseSlashedYMD returned unexpected value
	expected: %+v
	actual  : %+v
`, expect, actual)
	}

	// month is not number
	testee = "645/abcd/02"
	actual, err = ParseSlashedYMD(testee)
	if err == nil {
		t.Errorf("ParseSlashedYMD should return error")
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(`ParseSlashedYMD returned unexpected value
	expected: %+v
	actual  : %+v
`, expect, actual)
	}

	// day is not number
	testee = "645/03/abcd"
	actual, err = ParseSlashedYMD(testee)
	if err == nil {
		t.Errorf("ParseSlashedYMD should return error")
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(`ParseSlashedYMD returned unexpected value
	expected: %+v
	actual  : %+v
`, expect, actual)
	}
}
