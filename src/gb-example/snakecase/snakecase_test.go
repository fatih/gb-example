package snakecase

import "testing"

func TestConvert(t *testing.T) {
	s := "GbRocks"
	want := "gb_rocks"
	r := Convert(s)

	if want != r {
		t.Error("want %s got %s", want, r)
	}
}
