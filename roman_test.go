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

func TestFromStringComposed(t *testing.T) {
	tc := []struct {
		in   string
		want Roman
	}{
		{"II", I + I},
		{"III", 3 * I},
		{"IV", V - I},
		{"IIV", V - I - I},
		{"XIX", X - I + X},
		{"XVII", X + V + I + I},
		{"MDM", M - D + M},
		{"MDMLCXIX", M - D + M + C - L + X + X - I},
	}

	for _, c := range tc {
		got := FromString(c.in)

		if got != c.want {
			t.Errorf("want %d, got %d", c.want, got)
		}
	}
}
