package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type bubblysExpand struct {
}

func (b *bubblysExpand) ExpandSymbol(
	symbol *symbols.Symbol,
	prevSymPitch pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	isHalf := prevSymPitch == pitch.Pitch_LowG
	if isHalf {
		return []pitch.Pitch{
			pitch.Pitch_D, pitch.Pitch_LowG, pitch.Pitch_C, pitch.Pitch_LowG,
		}
	}

	return []pitch.Pitch{
		pitch.Pitch_LowG, pitch.Pitch_D, pitch.Pitch_LowG,
		pitch.Pitch_C, pitch.Pitch_LowG,
	}
}

func NewBubblysExpander() interfaces.SymbolExpander {
	return &bubblysExpand{}
}
