package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type peleExp struct {
}

func (p *peleExp) ExpandSymbol(
	symbol *symbols.Symbol,
	_ pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment

	isLight := emb.Weight == embellishment.Weight_Light
	var basicPele []pitch.Pitch
	symbolPitch := symbol.Note.Pitch
	if symbolPitch >= pitch.Pitch_LowA &&
		symbolPitch <= pitch.Pitch_D {
		basicPele = []pitch.Pitch{
			symbolPitch, pitch.Pitch_E, symbolPitch, pitch.Pitch_LowG,
		}
		if isLight {
			basicPele[3] = pitch.Pitch_C
		}
	}
	if symbolPitch == pitch.Pitch_E {
		basicPele = []pitch.Pitch{
			symbolPitch, pitch.Pitch_F, symbolPitch, pitch.Pitch_LowA,
		}
	}
	if symbolPitch == pitch.Pitch_F {
		basicPele = []pitch.Pitch{
			symbolPitch, pitch.Pitch_HighG, symbolPitch, pitch.Pitch_E,
		}
	}
	if symbolPitch == pitch.Pitch_HighG &&
		(emb.Variant == embellishment.Variant_Thumb ||
			emb.Variant == embellishment.Variant_Half) {
		basicPele = []pitch.Pitch{
			symbolPitch, pitch.Pitch_HighA, symbolPitch, pitch.Pitch_F,
		}
	}

	if emb.Variant == embellishment.Variant_Thumb {
		basicPele = append([]pitch.Pitch{pitch.Pitch_HighA}, basicPele...)
	}

	if emb.Variant == embellishment.Variant_NoVariant {
		basicPele = append([]pitch.Pitch{pitch.Pitch_HighG}, basicPele...)
	}

	return basicPele
}

func NewPelesExpander() interfaces.SymbolExpander {
	return &peleExp{}
}
