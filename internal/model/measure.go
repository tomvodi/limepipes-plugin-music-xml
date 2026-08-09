package model

import (
	"encoding/xml"

	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/barline"
)

type Measure struct {
	XMLName    xml.Name          `xml:"measure"`
	Number     int               `xml:"number,attr"`
	Barlines   []barline.Barline `xml:"barline"`
	Attributes *Attributes       `xml:"attributes,omitempty"`
	Notes      []Note            `xml:"note"`
}
