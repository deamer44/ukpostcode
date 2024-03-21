package ukpostcode

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
)

// figure out how to create a module
// figure out best practice for initialisation and have functions others can call

func main_create() {
	postcodes := readFromCsv("../ukpostcodes.csv")
	writeFile(serializePostcode(postcodes))
}

func writeFile(data []byte) {
	f, err := os.Create("content")
	check(err)
	numberofbytes, err := f.Write(data)
	check(err)
	fmt.Println(numberofbytes, "bytes written")
	err = f.Close()
	check(err)
}

func serializePostcode(postcodes map[string][]float64) []byte {
	var b bytes.Buffer
	var myByte byte = 0
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(postcodes); err != nil {
		fmt.Println("Error encoding postcodes", err)
		//i don't really understand this slice literal bits with the {}
		return []byte{myByte}
	}
	serializedData := b.Bytes()
	return serializedData
}

func readFromCsv(postcode string) map[string][]float64 {
	file, err := os.Open(postcode)
	check(err)
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()

	//create map of all postcodes to long/lat
	postcodes := make(map[string][]float64)
	// row[1] = postcode, row[2] = long, row[3] = lat
	for _, row := range data {
		lat, _ := strconv.ParseFloat(row[2], 64)
		long, _ := strconv.ParseFloat(row[3], 64)
		postcodes[row[1]] = append(postcodes[row[1]], lat)
		postcodes[row[1]] = append(postcodes[row[1]], long)
	}

	return postcodes
}
