package model

import "encoding/xml"

type Part struct {
	XMLName  xml.Name  `xml:"part"`
	ID       string    `xml:"id,attr"`
	Measures []Measure `xml:"measure"`
}
