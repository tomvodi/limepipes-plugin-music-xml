package barline

import (
	"encoding/xml"
	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/barline"
)

type Barline struct {
	XMLName  xml.Name `xml:"barline"`
	Location string   `xml:"location,attr"`
	Style    BarStyle `xml:"bar-style"`
	Repeat   *Repeat  `xml:"repeat,omitempty"`
}

func FromMusicModel(muMoBar *barline.Barline, loc Location) Barline {
	style := convertBarlineType(muMoBar.Type)

	barL := Barline{
		XMLName: xml.Name{
			Local: "barline",
		},
		Location: loc.String(),
	}

	if muMoBar.Time == barline.Time_Repeat {
		// A repeat is drawn as a thick-thin pair with the thick line on the
		// outside of the repeated section.
		dir := Forward
		style = HeavyLight
		if loc == Right {
			dir = Backward
			style = LightHeavy
		}
		barL.Repeat = NewRepeat(dir)
	}

	barL.Style = NewBarStyle(style)

	return barL
}

func convertBarlineType(barlineType barline.Type) Style {
	kebapC := strcase.KebabCase(barlineType.String())
	style, err := StyleString(kebapC)
	if err != nil {
		log.Error().Err(err).Msg("failed converting barline type (MusicModel) " +
			"to barline tyle (musicxml)")
		return None
	}
	return style
}
