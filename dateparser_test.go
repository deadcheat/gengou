package gengou

import (
	"reflect"
	"testing"
	"time"
)

func TestParseSlashedYMD(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Error("error occurred when time.LoadLocation")
	}
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			"success pattern",
			args{
				"645/03/02",
			},
			time.Date(645, 3, 2, 0, 0, 0, 0, loc),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSlashedYMD(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSlashedYMD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSlashedYMD() = %v, want %v", got, tt.want)
			}
		})
	}
}
