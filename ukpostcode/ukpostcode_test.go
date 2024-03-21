package ukpostcode

import (
	"testing"
)

func TestLookup(t *testing.T) {

	pl := PostcodeList{}
	pl.Initialise()
	r, _ := pl.Search("GL51 3xh")

	if !compare2FloatLiterals(r, Postcode{Lat: 51.882492, Long: -2.105819}) {
		t.Errorf("expected to get 51.882492 -2.105819, but got %f", r)
	}
}

// test length

func compare2FloatLiterals(f1 Postcode, f2 Postcode) bool {
	if len(f1) != len(f2) {
		return false
	}
	for i, e := range f1 {
		if e == f2[i] {
		} else {
			return false
		}
	}
	return true
}
