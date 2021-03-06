package gengou

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGengouCodeMarshaller(t *testing.T) {
	d, err := json.Marshal(idMap)
	if err != nil {
		t.Error("error occurred when json.Marshal")
	}

	result := make(map[string]GengouCode)
	if err := json.Unmarshal(d, &result); err != nil {
		t.Error("error occurred when json.Unmarshal")
	}

	if !reflect.DeepEqual(idMap, result) {
		t.Errorf(`Unmarshal returned unexpected value
	expected: %+v
	actual  : %+v
`, idMap, result)
	}
}
