package SHEIN

import "testing"

func TestGetImageUrl(t *testing.T) {
	input, expected :=
		"background-image:url(//image.webp);",
		"//image.webp"

	if result := extractImage(input); result != expected {
		t.Errorf("Expected %s, received %s", expected, result)
	}
}
