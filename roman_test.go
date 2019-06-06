package romanum

import "testing"

func TestFromStringSimple(t *testing.T) {
	tc := []struct {
		in   string
		want Roman
	}{
		{"I", I},
		{"V", V},
		{"X", X},
		{"L", L},
		{"C", C},
		{"D", D},
		{"M", M},
	}

	for _, c := range tc {
		got := FromString(c.in)

		if got != c.want {
			t.Errorf("want %d, got %d", c.want, got)
		}
	}
}
