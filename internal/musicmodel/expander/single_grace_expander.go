package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type singleGraceExpander struct {
}

func (s *singleGraceExpander) ExpandSymbol(symbol *symbols.Symbol, _ pitch.Pitch) {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return
	}

	emb := symbol.Note.Embellishment
	symbol.Note.ExpandedEmbellishment = []pitch.Pitch{
		emb.Pitch,
	}
}

func NewSingleGraceExpander() interfaces.SymbolExpander {
	return &singleGraceExpander{}
}
