package ukpostcode

import (
	"bytes"
	_ "embed"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// figure out how to create a module
// figure out best practice for initialisation and have functions others can call

func init() {

	//go: embed content
	content := desrializePostcode(readFile("content"))
	p := "GL51 xh"
	r, err := FindLatLong(p, content)
	if err != nil {
		fmt.Printf("error")
	} else {
		fmt.Println(r)
	}
}

func FindLatLong(postcode string, list map[string][]float64) ([]float64, error) {
	postcode, err := checkPostcode(postcode)
	if err != nil {
		fmt.Printf("string %s is incorrect\n", postcode)
	}
	return list[postcode], err
}

func checkPostcode(postcode string) (string, error) {
	postcode = strings.ToUpper(postcode)
	p := regexp.MustCompile("^([Gg][Ii][Rr] 0[Aa]{2})|((([A-Za-z][0-9]{1,2})|(([A-Za-z][A-Ha-hJ-Yj-y][0-9]{1,2})|(([AZa-z][0-9][A-Za-z])|([A-Za-z][A-Ha-hJ-Yj-y][0-9]?[A-Za-z])))) [0-9][A-Za-z]{2})$").MatchString(postcode)
	if p {
		return postcode, nil
	}
	return postcode, errors.New("incorrect postcode")
}

func readFile(file string) []byte {
	data, err := os.ReadFile(file)
	check(err)
	return data
}

func desrializePostcode(file []byte) map[string][]float64 {

	//return the file of bytes here!!! //
	b := bytes.NewBuffer(file)
	var postcodes map[string][]float64
	dec := gob.NewDecoder(b)
	if err := dec.Decode(&postcodes); err != nil {
		fmt.Println("Error decoding struct:", err)
		return make(map[string][]float64)
	}
	return postcodes
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
