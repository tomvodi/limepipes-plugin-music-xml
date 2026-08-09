package model

import (
	"encoding/xml"

	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/fermata"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/tied"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/tuplet"
)

type Notations struct {
	XMLName xml.Name         `xml:"notations"`
	Fermata *fermata.Fermata `xml:"fermata,omitempty"`
	Tied    *tied.Tied       `xml:"tied,omitempty"`
	Tuplet  *tuplet.Tuplet   `xml:"tuplet,omitempty"`
}

func NewNotations() *Notations {
	return &Notations{
		XMLName: xml.Name{
			Local: "notations",
		},
	}
}
