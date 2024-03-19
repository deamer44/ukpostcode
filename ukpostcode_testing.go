package ukpostcode

import "testing"

func TestLookup(t *testing.T) {

	p := Postcodes{}
	p.Initialise()
	r, _ := p.Search("GL51 3xh")

	if r != []float64{51.88249, -2.10582} {
		t.Errorf("expected to get 51.88249, -2.10582, but got %d", r)
	}

}
