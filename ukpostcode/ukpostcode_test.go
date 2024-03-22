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

func compare2FloatLiterals(p1 Postcode, p2 Postcode) bool {
	if p1.Lat == p2.Lat && p1.Long == p2.Long {
		return true
	}
	return false
}
