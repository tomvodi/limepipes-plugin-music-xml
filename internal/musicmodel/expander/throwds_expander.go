package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type throwdExpand struct {
}

func (t *throwdExpand) ExpandSymbol(
	symbol *symbols.Symbol,
	prevSymPitch pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment

	// A light throw is marked as such; anything else is a heavy throw.
	var expanded []pitch.Pitch
	if emb.Weight == embellishment.Weight_Light {
		expanded = []pitch.Pitch{
			pitch.Pitch_LowG, pitch.Pitch_D, pitch.Pitch_C,
		}
	} else {
		expanded = []pitch.Pitch{
			pitch.Pitch_LowG, pitch.Pitch_D, pitch.Pitch_LowG, pitch.Pitch_C,
		}
	}

	isHalf := prevSymPitch == pitch.Pitch_LowG
	if isHalf && len(expanded) > 0 {
		expanded = expanded[1:]
	}

	return expanded
}

func NewThrowdsExpander() interfaces.SymbolExpander {
	return &throwdExpand{}
}
