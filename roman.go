package romanum

import (
	"fmt"
	"math"
)

type Roman int

const (
	I Roman = 1
	V       = 5
	X       = 10
	L       = 50
	C       = 100
	D       = 500
	M       = 1000
)

func (r Roman) String() string {
	return ""
}

var r2a = map[string]Roman{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func FromString(roman string) Roman {
	total := 0
	last := math.MaxInt64
	curr := 0
	val := 0

	for _, r := range roman {
		curr = int(r2a[string(r)])
		fmt.Println(roman, ":", "[", string(r), "]", "curr:", curr, "val:", val, "total:", total)

		switch {
		case curr != last && curr < last && val != 0:
			total += val
			val = curr
		case curr > last:
			total += curr - val
			val = 0
		default:
			val += curr
		}
		fmt.Println(roman, ":", "[", string(r), "]", "curr:", curr, "val:", val, "total:", total)
		last = curr
	}
	total += val
	return Roman(total)
}
