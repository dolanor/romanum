package romanum

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
	val := 0
	for _, r := range roman {
		val += int(r2a[string(r)])
	}
	return Roman(val)
}
