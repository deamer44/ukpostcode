package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ukpostcode/ukpostcode"
)

func main() {
	//can HandleFunc pass in reference to postcodelist
	http.HandleFunc("/postcode", handlePostcode)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handlePostcode(w http.ResponseWriter, r *http.Request) {

	p := ukpostcode.PostcodeList{}
	p.Initialise()

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Only GET method allowed")
		return
	}

	postcode := r.URL.Query().Get("p")
	if postcode == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing 'p' parameter")
		return
	}

	validPostcode, err := ukpostcode.CheckPostcode(postcode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid postcode: %v", err)
		return
	}

	latLong, err := p.Search(validPostcode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching latitude and longitude: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(latLong)
}
