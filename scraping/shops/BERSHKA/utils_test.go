package BERSHKA

import "testing"

func TestGetPProductPrice(t *testing.T) {
	expected := float32(22.99)
	got := getProductPrice("22,99&nbsp;€")
	if got != expected {
		t.Errorf("Product price expected %f,  got %f", expected, got)
	}
}
