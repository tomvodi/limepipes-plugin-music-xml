package model

import (
	"encoding/xml"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
	"strings"
)

type Accidental struct {
	XMLName xml.Name `xml:"accidental"`
	Value   string   `xml:",chardata"`
}

func NewAccidentalFromMusicModel(acc accidental.Accidental) *Accidental {
	if acc == accidental.Accidental_NoAccidental {
		return nil
	}

	return &Accidental{
		XMLName: xml.Name{
			Local: "accidental",
		},
		Value: strings.ToLower(acc.String()),
	}
}
