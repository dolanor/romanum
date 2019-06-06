package romanum_test

import (
	"testing"

	"github.com/dolanor/romanum"
)

// I want to test the public API, but still
// I is shorter than romanum.I.
const (
	I = romanum.I
	V = romanum.V
	X = romanum.X
	L = romanum.L
	C = romanum.C
	D = romanum.D
	M = romanum.M
)

func TestParseSimple(t *testing.T) {
	tc := []struct {
		in   string
		want romanum.Roman
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
		got := romanum.Parse(c.in)

		if got != c.want {
			t.Errorf("want %d, got %d", c.want, got)
		}
	}
}

func TestParseComposed(t *testing.T) {
	tc := []struct {
		in   string
		want romanum.Roman
	}{
		{"II", I + I},
		{"III", I + I + I},
		{"IV", V - I},
		{"VI", V + I},
		{"XVII", X + V + I + I},
		{"XIX", X - I + X},
		{"MCM", M - C + M},
		{"MCMLCXIX", M - C + M + C - L + X - I + X},
	}

	for _, c := range tc {
		got := romanum.Parse(c.in)

		if got != c.want {
			t.Errorf("want %d, got %d", c.want, got)
		}
	}
}
