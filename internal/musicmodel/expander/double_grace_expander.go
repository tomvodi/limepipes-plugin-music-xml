package expander

import (
	"github.com/rs/zerolog/log"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type dblGraceExpand struct {
}

func (d *dblGraceExpand) ExpandSymbol(
	symbol *symbols.Symbol,
	_ pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	if symbol.Note.Pitch == pitch.Pitch_LowG {
		log.Error().Msg("can't play double grace on LowG")
		return nil
	}

	emb := symbol.Note.Embellishment

	if symbol.Note.Pitch > emb.Pitch {
		log.Error().Msgf("can't play double grace %s on a melody note with pitch %s",
			emb.Pitch.String(), symbol.Note.Pitch.String())
		return nil
	}

	return []pitch.Pitch{
		emb.Pitch,
		symbol.Note.Pitch - 1,
	}
}

func NewDoubleGraceExpander() interfaces.SymbolExpander {
	return &dblGraceExpand{}
}
