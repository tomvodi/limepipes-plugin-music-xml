package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type taorExpand struct {
}

func (t *taorExpand) ExpandSymbol(
	symbol *symbols.Symbol,
	prevSymPitch pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment

	// regular taorluath
	var expanded []pitch.Pitch
	isHalf := prevSymPitch == pitch.Pitch_LowG
	isB := emb.Pitch == pitch.Pitch_B
	if isHalf {
		expanded = []pitch.Pitch{
			pitch.Pitch_D, pitch.Pitch_LowG, pitch.Pitch_E,
		}
	}
	if isB {
		expanded = []pitch.Pitch{
			pitch.Pitch_LowG, pitch.Pitch_B, pitch.Pitch_LowG, pitch.Pitch_E,
		}
	}
	if expanded == nil {
		expanded = []pitch.Pitch{
			pitch.Pitch_LowG, pitch.Pitch_D, pitch.Pitch_LowG, pitch.Pitch_E,
		}
	}

	return expanded
}

func NewTaorluathsExpander() interfaces.SymbolExpander {
	return &taorExpand{}
}
