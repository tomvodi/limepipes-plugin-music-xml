package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type birlsExp struct {
}

func (b *birlsExp) ExpandSymbol(
	symbol *symbols.Symbol,
	prevSymPitch pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment

	var expanded []pitch.Pitch
	isHalf := prevSymPitch == pitch.Pitch_LowA
	if isHalf {
		expanded = []pitch.Pitch{
			pitch.Pitch_LowG, pitch.Pitch_LowA, pitch.Pitch_LowG,
		}
	}
	// A grace birl is a birl led by a single grace note. The model has one
	// type for both the g grace and the thumb grace form, so it is always
	// rendered as the g grace one.
	if emb.Type == embellishment.Type_GraceBirl ||
		emb.Variant == embellishment.Variant_G {
		expanded = []pitch.Pitch{
			pitch.Pitch_HighG, pitch.Pitch_LowA, pitch.Pitch_LowG,
			pitch.Pitch_LowA, pitch.Pitch_LowG,
		}
	}
	if emb.Variant == embellishment.Variant_Thumb {
		expanded = []pitch.Pitch{
			pitch.Pitch_HighA, pitch.Pitch_LowA, pitch.Pitch_LowG,
			pitch.Pitch_LowA, pitch.Pitch_LowG,
		}
	}
	if expanded == nil {
		expanded = []pitch.Pitch{
			pitch.Pitch_LowA, pitch.Pitch_LowG,
			pitch.Pitch_LowA, pitch.Pitch_LowG,
		}
	}

	return expanded
}

func NewBirlsExpander() interfaces.SymbolExpander {
	return &birlsExp{}
}
