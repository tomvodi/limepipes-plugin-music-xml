package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type trplStrikeExp struct {
}

func (t *trplStrikeExp) ExpandSymbol(
	symbol *symbols.Symbol,
	_ pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment

	isLight := emb.Weight == embellishment.Weight_Light
	var basicTrplStrike []pitch.Pitch
	symbolPitch := symbol.Note.Pitch
	if symbolPitch >= pitch.Pitch_LowA &&
		symbolPitch <= pitch.Pitch_D {
		basicTrplStrike = []pitch.Pitch{
			pitch.Pitch_LowG, symbolPitch, pitch.Pitch_LowG,
			symbolPitch, pitch.Pitch_LowG,
		}
		if isLight {
			basicTrplStrike[0] = pitch.Pitch_C
			basicTrplStrike[2] = pitch.Pitch_C
			basicTrplStrike[4] = pitch.Pitch_C
		}
	}
	if symbolPitch == pitch.Pitch_E {
		basicTrplStrike = []pitch.Pitch{
			pitch.Pitch_LowA, symbolPitch, pitch.Pitch_LowA,
			symbolPitch, pitch.Pitch_LowA,
		}
	}
	if symbolPitch == pitch.Pitch_F {
		basicTrplStrike = []pitch.Pitch{
			pitch.Pitch_E, symbolPitch, pitch.Pitch_E,
			symbolPitch, pitch.Pitch_E,
		}
	}
	if symbolPitch == pitch.Pitch_HighG {
		basicTrplStrike = []pitch.Pitch{
			pitch.Pitch_F, symbolPitch, pitch.Pitch_F,
			symbolPitch, pitch.Pitch_F,
		}
	}
	if symbolPitch == pitch.Pitch_HighA &&
		(emb.Variant == embellishment.Variant_Half ||
			emb.Variant == embellishment.Variant_NoVariant) {
		basicTrplStrike = []pitch.Pitch{
			pitch.Pitch_HighG, symbolPitch, pitch.Pitch_HighG,
			symbolPitch, pitch.Pitch_HighG,
		}
	}
	if emb.Variant != embellishment.Variant_NoVariant {
		basicTrplStrike = append([]pitch.Pitch{symbolPitch}, basicTrplStrike...)
	}

	if emb.Variant == embellishment.Variant_Thumb {
		basicTrplStrike = append([]pitch.Pitch{pitch.Pitch_HighA}, basicTrplStrike...)
	}
	if emb.Variant == embellishment.Variant_G {
		basicTrplStrike = append([]pitch.Pitch{pitch.Pitch_HighG}, basicTrplStrike...)
	}

	return basicTrplStrike
}

func NewTripleStrikesExpander() interfaces.SymbolExpander {
	return &trplStrikeExp{}
}
