package ukpostcode

import (
	"testing"
)

func TestLookup(t *testing.T) {

	p := Postcodes{}
	p.Initialise()
	r, _ := p.Search("GL51 3xh")

	if !compare2FloatLiterals(r, []float64{51.882492, -2.105819}) {
		t.Errorf("expected to get 51.882492 -2.105819, but got %f", r)
	}
	//test length, if it doesn't match 2 values then error
	if compare2FloatLiterals(r, []float64{51.882492, -2.105819, 3}) {
		t.Errorf("Length is incorrect, but got %f", r)
	}
}

// test length

func compare2FloatLiterals(f1 []float64, f2 []float64) bool {
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
