package inverted

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/tokenize_tests.json")
	if err != nil {
		panic(err)
	}
	tests := make(map[string][]string)
	json.Unmarshal(data, &tests)

	InitTokenize()
	for input := range tests {
		expected := tests[input]
		got := Tokenize(input)
		if !reflect.DeepEqual(expected, got) {
			t.Error("expected", expected, "got")
		}
	}
}
