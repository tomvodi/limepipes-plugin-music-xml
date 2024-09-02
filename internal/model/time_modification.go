package model

import (
	"encoding/xml"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tuplet"
)

type TimeModification struct {
	XMLName     xml.Name `xml:"time-modification"`
	ActualNotes uint32   `xml:"actual-notes"`
	NormalNotes uint32   `xml:"normal-notes"`
}

func NewTimeModification(tpl *tuplet.Tuplet) *TimeModification {
	return &TimeModification{
		XMLName: xml.Name{
			Local: "time-modification",
		},
		ActualNotes: tpl.VisibleNotes,
		NormalNotes: tpl.PlayedNotes,
	}
}
