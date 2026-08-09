package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

type strikesExp struct {
}

func (s *strikesExp) ExpandSymbol(
	symbol *symbols.Symbol,
	_ pitch.Pitch,
) []pitch.Pitch {
	if symbol == nil || symbol.Note == nil || symbol.Note.Embellishment == nil {
		return nil
	}

	emb := symbol.Note.Embellishment
	if emb.Variant == embellishment.Variant_NoVariant {
		// Single strike. The struck pitch is always the note below the melody
		// note, so the model does not carry it explicitly.
		if symbol.Note.Pitch <= pitch.Pitch_LowG {
			return nil
		}

		return []pitch.Pitch{symbol.Note.Pitch - 1}
	}
	if emb.Variant == embellishment.Variant_Half {
		return halfStrikePitchesForWeight(symbol.Note.Pitch, emb.Weight)
	}
	if emb.Variant == embellishment.Variant_Thumb {
		pitchesHalf := halfStrikePitchesForWeight(symbol.Note.Pitch, emb.Weight)

		return append([]pitch.Pitch{pitch.Pitch_HighA}, pitchesHalf...)
	}
	if emb.Variant == embellishment.Variant_G {
		pitchesHalf := halfStrikePitchesForWeight(symbol.Note.Pitch, emb.Weight)

		return append([]pitch.Pitch{pitch.Pitch_HighG}, pitchesHalf...)
	}

	return nil
}

func halfStrikePitchesForWeight(
	p pitch.Pitch,
	weight embellishment.Weight,
) []pitch.Pitch {
	pitches := []pitch.Pitch{p}
	if p >= pitch.Pitch_LowA && p <= pitch.Pitch_C {
		pitches = append(pitches, pitch.Pitch_LowG)
	}

	if p == pitch.Pitch_D {
		if weight == embellishment.Weight_Light {
			pitches = append(pitches, pitch.Pitch_C)
		} else {
			pitches = append(pitches, pitch.Pitch_LowG)
		}
	}
	if p == pitch.Pitch_E {
		pitches = append(pitches, pitch.Pitch_LowA)
	}
	if p >= pitch.Pitch_F && p <= pitch.Pitch_HighG {
		pitches = append(pitches, p-1)
	}

	return pitches
}

func NewStrikesExpander() interfaces.SymbolExpander {
	return &strikesExp{}
}
