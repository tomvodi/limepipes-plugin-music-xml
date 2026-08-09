package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type dblStrikeExp struct {
}

func (d *dblStrikeExp) ExpandSymbol(
	symbol *symbols.Symbol,
	_ pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment

	isLight := emb.Weight == embellishment.Weight_Light
	var basicDblStrike []pitch.Pitch
	symbolPitch := symbol.Note.Pitch
	if symbolPitch >= pitch.Pitch_LowA &&
		symbolPitch <= pitch.Pitch_D {
		basicDblStrike = []pitch.Pitch{
			pitch.Pitch_LowG, symbolPitch, pitch.Pitch_LowG,
		}
		if isLight {
			basicDblStrike[0] = pitch.Pitch_C
			basicDblStrike[2] = pitch.Pitch_C
		}
	}
	if symbolPitch == pitch.Pitch_E {
		basicDblStrike = []pitch.Pitch{
			pitch.Pitch_LowA, symbolPitch, pitch.Pitch_LowA,
		}
	}
	if symbolPitch == pitch.Pitch_F {
		basicDblStrike = []pitch.Pitch{
			pitch.Pitch_E, symbolPitch, pitch.Pitch_E,
		}
	}
	if symbolPitch == pitch.Pitch_HighG {
		basicDblStrike = []pitch.Pitch{
			pitch.Pitch_F, symbolPitch, pitch.Pitch_F,
		}
	}
	if symbolPitch == pitch.Pitch_HighA &&
		(emb.Variant == embellishment.Variant_Half ||
			emb.Variant == embellishment.Variant_NoVariant) {
		basicDblStrike = []pitch.Pitch{
			pitch.Pitch_HighG, symbolPitch, pitch.Pitch_HighG,
		}
	}
	if emb.Variant != embellishment.Variant_NoVariant {
		basicDblStrike = append([]pitch.Pitch{symbolPitch}, basicDblStrike...)
	}

	if emb.Variant == embellishment.Variant_Thumb {
		basicDblStrike = append([]pitch.Pitch{pitch.Pitch_HighA}, basicDblStrike...)
	}
	if emb.Variant == embellishment.Variant_G {
		basicDblStrike = append([]pitch.Pitch{pitch.Pitch_HighG}, basicDblStrike...)
	}

	return basicDblStrike
}

func NewDoubleStrikesExpander() interfaces.SymbolExpander {
	return &dblStrikeExp{}
}
