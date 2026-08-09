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

// SetEnding hangs an ending bracket off the barline at loc, adding a barline
// there if the measure does not already have one.
func (m *Measure) SetEnding(loc barline.Location, ending *barline.Ending) {
	for i := range m.Barlines {
		if m.Barlines[i].Location == loc.String() {
			m.Barlines[i].Ending = ending

			return
		}
	}

	bar := barline.NewEndingBarline(loc, ending)

	// The left edge of the bar is written before the right one, so a measure
	// that already had a right barline gets the new one put in front.
	if loc == barline.Left {
		m.Barlines = append([]barline.Barline{bar}, m.Barlines...)

		return
	}

	m.Barlines = append(m.Barlines, bar)
}
