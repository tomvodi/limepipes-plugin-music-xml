package expander

import (
	"github.com/rs/zerolog/log"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/musicmodel"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type embExpander struct {
	table ExpandTable
}

func (e *embExpander) ExpandModel(model musicmodel.MusicModel) interfaces.Expansions {
	exps := interfaces.Expansions{}
	for _, t := range model {
		e.expandTuneInto(t, exps)
	}

	return exps
}

func (e *embExpander) ExpandTune(t *tune.Tune) interfaces.Expansions {
	exps := interfaces.Expansions{}
	e.expandTuneInto(t, exps)

	return exps
}

func (e *embExpander) expandTuneInto(t *tune.Tune, exps interfaces.Expansions) {
	if t == nil {
		return
	}

	prevSymPitch := pitch.Pitch_NoPitch
	for _, m := range t.Measures {
		for _, symbol := range m.Symbols {
			if expanded := e.expandSymbol(symbol, prevSymPitch); expanded != nil {
				exps[symbol.Note] = expanded
			}

			if symbol.IsValidNote() {
				prevSymPitch = symbol.Note.Pitch
			} else {
				prevSymPitch = pitch.Pitch_NoPitch
			}
		}
	}
}

func (e *embExpander) expandSymbol(
	symbol *symbols.Symbol,
	prevSymPitch pitch.Pitch,
) []pitch.Pitch {
	if !symbol.IsValidNote() {
		return nil
	}

	if symbol.Note.Embellishment == nil {
		return nil
	}

	expander, ok := e.table[keyFromEmbellishment(symbol.Note.Embellishment)]
	if !ok {
		log.Error().Msgf("no embellishment expander for %v", symbol.Note.Embellishment)
		return nil
	}

	return expander.ExpandSymbol(symbol, prevSymPitch)
}

func NewEmbellishmentExpander() interfaces.EmbellishmentExpander {
	return &embExpander{
		table: newSymbolExpanderTable(),
	}
}
