package romanum

import (
	"fmt"
	"math"
)

type Roman int64

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

var roman2arabic = map[string]Roman{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// Parse parses a subtractive notation of roman numerals.
func Parse(roman string) Roman {
	var curr, last, tmpSum, total int64
	last = math.MaxInt64

	for _, r := range roman {
		curr = int64(roman2arabic[string(r)])
		fmt.Println(roman, ":", "[", string(r), "]", "curr:", curr, "tmpSum:", tmpSum, "total:", total)

		switch {
		case curr != last && curr < last && tmpSum != 0:
			total += tmpSum
			tmpSum = curr
		case curr > last:
			total += curr - tmpSum
			tmpSum = 0
		default:
			tmpSum += curr
		}
		fmt.Println(roman, ":", "[", string(r), "]", "curr:", curr, "tmpSum:", tmpSum, "total:", total)
		last = curr
	}
	total += tmpSum
	return Roman(total)
}
