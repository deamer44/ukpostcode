package main

import (
	"fmt"
	"ukpostcode/ukpostcode"
)

func main() {
	fmt.Println("IT works!")
	p := ukpostcode.Postcodes{}
	p.Initialise()
	r, _ := p.Search("GL51 3xh")
	print(r)
}
