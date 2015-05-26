package util

import "testing"

func TestSnakeCase(t *testing.T) {
	s := "GbRocks"
	want := "Gb_Rocks"
	r := SnakeCase(s)

	if want != r {
		t.Error("want %s got %s", want, r)
	}
}
