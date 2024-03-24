package ukpostcode

import (
	"testing"
)

func TestLookup(t *testing.T) {

	pl := PostcodeList{}
	pl.Initialise()
	r, _ := pl.Search("GL51 3xh")

	t.Run("Testing postcode lookup", func(t *testing.T) {
		if !compare2FloatLiterals(r, Postcode{Lat: 51.882492, Long: -2.105819}) {
			t.Errorf("expected to get 51.882492 -2.105819, but got %f", r)
		}
	})

	t.Run("Testing syntatically correct postcode, but one without a long / lat", func(t *testing.T) {
		_, err := pl.Search("gl52 3xt")
		if err != nil {
			return
		} else {
			t.Errorf("long / lat should not be found and an error should have been created, %v", err)
		}
	})
}

// test length

func compare2FloatLiterals(p1 Postcode, p2 Postcode) bool {
	if p1.Lat == p2.Lat && p1.Long == p2.Long {
		return true
	}
	return false
}
