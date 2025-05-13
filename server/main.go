package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	Add Add `xml:"Add"`
}

type Add struct {
	IntA int `xml:"intA"`
	IntB int `xml:"intB"`
}

type AddResponseEnvelope struct {
	XMLName xml.Name        `xml:"soap:Envelope"`
	SoapNS  string          `xml:"xmlns:soap,attr"`
	Body    AddResponseBody `xml:"soap:Body"`
}

type AddResponseBody struct {
	AddResponse AddResult `xml:"AddResponse"`
}

type AddResult struct {
	Xmlns     string `xml:"xmlns,attr"`
	AddResult int    `xml:"AddResult"`
}

func soapHandler(w http.ResponseWriter, r *http.Request) {
	// Basic auth check
	username, password, ok := r.BasicAuth()
	if !ok || username != "admin" || password != "Admin@123" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	var envelope Envelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		http.Error(w, "Invalid SOAP", http.StatusBadRequest)
		return
	}

	sum := envelope.Body.Add.IntA + envelope.Body.Add.IntB
	response := AddResponseEnvelope{
		SoapNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: AddResponseBody{
			AddResponse: AddResult{
				Xmlns:     "http://tempuri.org/",
				AddResult: sum,
			},
		},
	}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/soap", soapHandler)
	fmt.Println("SOAP Server running at http://localhost:8080/soap")
	http.ListenAndServe(":8080", nil)
}
