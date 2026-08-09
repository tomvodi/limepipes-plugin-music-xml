package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type grpExpander struct {
}

func (g *grpExpander) ExpandSymbol(
	symbol *symbols.Symbol,
	prevSymPitch pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment
	if emb.Variant == embellishment.Variant_NoVariant {
		// regular grip
		var expanded []pitch.Pitch
		isHalf := prevSymPitch == pitch.Pitch_LowG
		isB := emb.Pitch == pitch.Pitch_B
		if isHalf {
			expanded = []pitch.Pitch{pitch.Pitch_D, pitch.Pitch_LowG}
		}
		if isB {
			expanded = []pitch.Pitch{
				pitch.Pitch_LowG, pitch.Pitch_B, pitch.Pitch_LowG,
			}
		}
		if expanded == nil {
			expanded = []pitch.Pitch{
				pitch.Pitch_LowG, pitch.Pitch_D, pitch.Pitch_LowG,
			}
		}

		return expanded
	}

	basicGrp := []pitch.Pitch{
		symbol.Note.Pitch, pitch.Pitch_LowG, pitch.Pitch_D, pitch.Pitch_LowG,
	}
	if emb.Pitch == pitch.Pitch_B {
		basicGrp[2] = pitch.Pitch_B
	}
	if symbol.Note.Pitch == pitch.Pitch_F {
		basicGrp[2] = pitch.Pitch_F
	}
	if symbol.Note.Pitch == pitch.Pitch_HighG &&
		emb.Variant == embellishment.Variant_Thumb {
		basicGrp[2] = pitch.Pitch_F
	}

	if emb.Variant == embellishment.Variant_G {
		basicGrp = append([]pitch.Pitch{pitch.Pitch_HighG}, basicGrp...)
	}
	if emb.Variant == embellishment.Variant_Thumb {
		basicGrp = append([]pitch.Pitch{pitch.Pitch_HighA}, basicGrp...)
	}

	return basicGrp
}

func NewGripsExpander() interfaces.SymbolExpander {
	return &grpExpander{}
}
