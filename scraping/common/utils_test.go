package common

import (
	"encoding/json"
	"testing"
)

func TestGroup(t *testing.T) {
	numbers := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	expected := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7},
	}

	got := Group[int](numbers, 2)
	strExpected, _ := json.Marshal(expected)
	resExpected, _ := json.Marshal(got)

	if string(strExpected) != string(resExpected) {
		t.Errorf("%v expected but %v got", expected, got)
	}
}
