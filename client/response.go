package main

import "encoding/xml"

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SOAPBody `xml:"Body"`
}

type SOAPBody struct {
	AddResponse AddResponse `xml:"AddResponse"`
}

type AddResponse struct {
	XMLName   xml.Name `xml:"AddResponse"`
	AddResult int      `xml:"AddResult"`
}
