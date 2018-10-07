package util

import "testing"

func TestParseContinuousColorHexString(t *testing.T) {
	str := "111111222222333333444444555555666666777777888888999999"

	colors := ParseContinuousColorHexString(str)
	if len(colors) != 9 {
		t.Error("expected size 6")
	}
	if colors[0] != "111111" {
		t.Error("expected match")
	}
	if colors[1] != "222222" {
		t.Error("expected match")
	}
	if colors[8] != "999999" {
		t.Error("expected match")
	}
}
