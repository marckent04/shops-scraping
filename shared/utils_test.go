package shared

import (
	"encoding/json"
	"testing"
)

type Product struct {
	Name  string
	Price float32
}

func TestSlicesFilterWithStrings(t *testing.T) {
	arr := []string{
		"ab",
		"abc",
		"dbcd",
		"dfgh",
	}

	expected := []string{
		"dbcd",
		"dfgh",
	}

	got := SlicesFilter(arr, func(s string) bool {
		return len(s) > 3
	})

	expStr, _ := json.Marshal(expected)
	gotStr, _ := json.Marshal(got)

	if string(expStr) != string(gotStr) {
		t.Errorf("%v expected %v found", expected, got)
	}

}

func TestSlicesFilterWithStruts(t *testing.T) {

	arr := []Product{
		{Name: "Cerelac", Price: 2000},
		{Name: "Doudou", Price: 2400},
		{Name: "Pampers", Price: 3000},
	}

	expected := []Product{
		{Name: "Cerelac", Price: 2000},
		{Name: "Doudou", Price: 2400},
	}

	got := SlicesFilter(arr, func(p Product) bool {
		return p.Price < 3000
	})

	expStr, _ := json.Marshal(expected)
	gotStr, _ := json.Marshal(got)

	if string(expStr) != string(gotStr) {
		t.Errorf("%v expected %v found", expected, got)
	}

}
