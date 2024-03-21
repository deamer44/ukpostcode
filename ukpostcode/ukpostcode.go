package ukpostcode

import (
	"bytes"
	"embed"
	_ "embed"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type PostcodeList struct {
	data map[string][]float64
}

type Postcode struct {
	Lat  float64
	Long float64
}

func (p *Postcode) len() int {
	return reflect.TypeOf(p).NumField()
}

func (p *PostcodeList) Initialise() {
	p.data = desrializePostcode(readData("content"))
}

func (p *PostcodeList) Search(postcode string) (Postcode, error) {
	postcode, err := checkPostcode(postcode)
	if err != nil {
		fmt.Printf("string %s is incorrect\n", postcode)
	}
	return returnJson(p.data[postcode])
}

func checkPostcode(postcode string) (string, error) {
	postcode = strings.ToUpper(postcode)
	p := regexp.MustCompile("^([Gg][Ii][Rr] 0[Aa]{2})|((([A-Za-z][0-9]{1,2})|(([A-Za-z][A-Ha-hJ-Yj-y][0-9]{1,2})|(([AZa-z][0-9][A-Za-z])|([A-Za-z][A-Ha-hJ-Yj-y][0-9]?[A-Za-z])))) [0-9][A-Za-z]{2})$").MatchString(postcode)
	if p {
		return postcode, nil
	}
	//returns lat long
	return postcode, errors.New("incorrect postcode")
}

//go:embed content
var content embed.FS

func readData(file string) []byte {
	data, err := content.ReadFile(file)
	check(err)
	return data
}

func desrializePostcode(file []byte) map[string][]float64 {
	b := bytes.NewBuffer(file)
	PostcodeList := make(map[string][]float64)
	dec := gob.NewDecoder(b)
	if err := dec.Decode(&PostcodeList); err != nil {
		fmt.Println("Error decoding struct:", err)
		return make(map[string][]float64)
	}
	return PostcodeList
}

func returnJson(c []float64) []byte {
	p := Postcode{Lat: c[0], Long: c[1]}
	json, err := json.Marshal(p)
	check(err)
	return json
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
