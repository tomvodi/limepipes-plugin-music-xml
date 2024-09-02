package interfaces

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/musicmodel"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
)

type SymbolExpander interface {
	// ExpandSymbol expands all embellishments in the music model symbol
	ExpandSymbol(symbol *symbols.Symbol, prevSymPitch pitch.Pitch)
}

type EmbellishmentExpander interface {
	// ExpandModel expands all embellishments in music model
	ExpandModel(model musicmodel.MusicModel)

	// ExpandTune expands all embellishments in music model tune
	ExpandTune(tune *tune.Tune)
}
