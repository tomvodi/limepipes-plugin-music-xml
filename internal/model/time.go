package model

import (
	"encoding/xml"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
)

type Time struct {
	XMLName  xml.Name `xml:"time"`
	Beats    uint32   `xml:"beats"`
	BeatType uint32   `xml:"beat-type"`
}

func NewTime(time *measure.TimeSignature) *Time {
	return &Time{
		XMLName: xml.Name{
			Local: "time",
		},
		Beats:    time.Beats,
		BeatType: time.BeatType,
	}
}
