package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type dblExpand struct {
}

func (d *dblExpand) ExpandSymbol(
	symbol *symbols.Symbol,
	_ pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment

	switch emb.Variant {
	case embellishment.Variant_NoVariant:
		return handleRegular(symbol.Note)
	case embellishment.Variant_Thumb:
		return handleThumb(symbol.Note)
	case embellishment.Variant_Half:
		return handleHalf(symbol.Note)
	}

	return nil
}

func handleRegular(note *symbols.Note) []pitch.Pitch {
	if note.Pitch >= pitch.Pitch_HighG {
		return []pitch.Pitch{
			note.Pitch,
			note.Pitch - 1,
		}
	}

	if note.Pitch >= pitch.Pitch_D {
		return []pitch.Pitch{
			pitch.Pitch_HighG,
			note.Pitch,
			note.Pitch + 1,
		}
	}

	return []pitch.Pitch{
		pitch.Pitch_HighG,
		note.Pitch,
		pitch.Pitch_D,
	}
}

func handleThumb(note *symbols.Note) []pitch.Pitch {
	if note.Pitch >= pitch.Pitch_HighG {
		return nil
	}

	if note.Pitch >= pitch.Pitch_D {
		return []pitch.Pitch{
			pitch.Pitch_HighA,
			note.Pitch,
			note.Pitch + 1,
		}
	}

	return []pitch.Pitch{
		pitch.Pitch_HighA,
		note.Pitch,
		pitch.Pitch_D,
	}
}

func handleHalf(note *symbols.Note) []pitch.Pitch {
	if note.Pitch >= pitch.Pitch_HighG {
		return nil
	}

	if note.Pitch >= pitch.Pitch_D {
		return []pitch.Pitch{
			note.Pitch,
			note.Pitch + 1,
		}
	}

	return []pitch.Pitch{
		note.Pitch,
		pitch.Pitch_D,
	}
}

func NewDoublingsExpander() interfaces.SymbolExpander {
	return &dblExpand{}
}
