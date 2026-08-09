package interfaces

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/musicmodel"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
)

// Expansions holds the grace note pitches an embellishment expands to, keyed by
// the note the embellishment belongs to.
// The music model has no field for this because an expansion is a rendering
// detail and not part of the interchange format, so it is kept beside the model.
type Expansions map[*symbols.Note][]pitch.Pitch

// Get returns the expanded pitches for note. It is nil safe so callers can use
// it on notes that have no embellishment.
func (e Expansions) Get(note *symbols.Note) []pitch.Pitch {
	if e == nil || note == nil {
		return nil
	}

	return e[note]
}

type SymbolExpander interface {
	// ExpandSymbol returns the pitches the symbol's embellishment expands to.
	// It returns nil if the symbol carries no expandable embellishment.
	ExpandSymbol(symbol *symbols.Symbol, prevSymPitch pitch.Pitch) []pitch.Pitch
}

type EmbellishmentExpander interface {
	// ExpandModel expands all embellishments in music model
	ExpandModel(model musicmodel.MusicModel) Expansions

	// ExpandTune expands all embellishments in music model tune
	ExpandTune(tune *tune.Tune) Expansions
}
