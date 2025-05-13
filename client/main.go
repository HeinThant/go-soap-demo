package main

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func main() {
	soapBody := `
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <Add xmlns="http://tempuri.org/">
      <intA>10</intA>
      <intB>15</intB>
    </Add>
  </soap:Body>
</soap:Envelope>`

	req, err := http.NewRequest("POST", "http://localhost:8080/soap", bytes.NewBuffer([]byte(soapBody)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://tempuri.org/Add")

	// Add Basic Auth Header
	auth := base64.StdEncoding.EncodeToString([]byte("admin:Admin@123"))
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Raw XML Response:")
	fmt.Println(string(body))

	// Parse XML response
	var soapResp SOAPEnvelope
	if err := xml.Unmarshal(body, &soapResp); err != nil {
		panic(err)
	}
	fmt.Printf("\nParsed Result: %d\n", soapResp.Body.AddResponse.AddResult)
}
