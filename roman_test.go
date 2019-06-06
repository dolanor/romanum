package romanum

import "testing"

func TestParseSimple(t *testing.T) {
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
		got := Parse(c.in)

		if got != c.want {
			t.Errorf("want %d, got %d", c.want, got)
		}
	}
}

func TestParseComposed(t *testing.T) {
	tc := []struct {
		in   string
		want Roman
	}{
		{"II", I + I},
		{"III", I + I + I},
		{"IV", V - I},
		{"IIV", V - I - I},
		{"XIX", X - I + X},
		{"XVII", X + V + I + I},
		{"MDM", M - D + M},
		{"MCMLCXIX", M - C + M + C - L + X - I + X},
	}

	for _, c := range tc {
		got := Parse(c.in)

		if got != c.want {
			t.Errorf("want %d, got %d", c.want, got)
		}
	}
}
