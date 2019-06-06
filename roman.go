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

// String displays the roman numeral in additive notation:
// XVIIII instead of XIX.
func (r Roman) String() string {
	var val = int64(r)
	var output string
	for val > 0 {
		fmt.Println("output:", output, "val:", val)
		switch {
		case val-int64(M) >= 0:
			val -= int64(M)
			output += "M"
		case val-int64(D) >= 0:
			val -= int64(D)
			output += "D"
		case val-int64(C) >= 0:
			val -= int64(C)
			output += "C"
		case val-int64(L) >= 0:
			val -= int64(L)
			output += "L"
		case val-int64(X) >= 0:
			val -= int64(X)
			output += "X"
		case val-int64(V) >= 0:
			val -= int64(V)
			output += "V"
		case val-int64(I) >= 0:
			val -= int64(I)
			output += "I"
		}
	}

	return output
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
