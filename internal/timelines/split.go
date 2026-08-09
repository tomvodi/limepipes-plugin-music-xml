package timelines

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
)

// splitMeasuresAtTimelines cuts every measure so that no time line bracket
// starts or ends in the middle of a bar. MusicXML endings hang off barlines, so
// a bracket has to line up with one.
func splitMeasuresAtTimelines(measures []*measure.Measure) []*measure.Measure {
	var out []*measure.Measure
	for _, m := range measures {
		out = append(out, splitAtBrackets(m)...)
	}

	return out
}

// splitAtBrackets cuts one measure so that each time line bracket inside it
// gets a bar of its own:
//
//	| B ⌐1 D ¬ ⌐2 E ¬ |   becomes   | B | ⌐1 D ¬ | ⌐2 E ¬ |
//
// The bar's outer boundaries are preserved: the left barline and time signature
// stay with the first piece, the right barline with the last. A measure that
// holds no bracket, or that is exactly one bracket already, is returned as is.
func splitAtBrackets(m *measure.Measure) []*measure.Measure {
	if !Has(m) || isSingleBracket(m) {
		return []*measure.Measure{m}
	}

	pieces := cutAtBracketBoundaries(m.Symbols)

	out := make([]*measure.Measure, 0, len(pieces))
	for _, syms := range pieces {
		out = append(out, &measure.Measure{Symbols: syms})
	}

	first, last := out[0], out[len(out)-1]

	first.LeftBarline = m.LeftBarline
	first.Time = m.Time
	first.Comments = m.Comments
	first.InlineTexts = m.InlineTexts

	last.RightBarline = m.RightBarline

	return out
}

// cutAtBracketBoundaries slices a run of symbols so that a cut falls before
// every bracket start and after every bracket end. Empty slices are never
// produced, so two adjacent brackets yield two pieces rather than three.
func cutAtBracketBoundaries(syms []*symbols.Symbol) [][]*symbols.Symbol {
	var pieces [][]*symbols.Symbol
	var current []*symbols.Symbol

	cut := func() {
		if len(current) > 0 {
			pieces = append(pieces, current)
			current = nil
		}
	}

	for _, sym := range syms {
		if IsStart(sym) {
			cut()
		}

		current = append(current, sym)

		if IsEnd(sym) {
			cut()
		}
	}
	cut()

	return pieces
}

// isSingleBracket reports whether the measure is one bracket and nothing else,
// which is the shape MusicXML wants and so needs no cutting.
func isSingleBracket(m *measure.Measure) bool {
	if len(m.Symbols) < 2 {
		return false
	}

	brackets := 0
	for _, sym := range m.Symbols {
		if sym.Timeline != nil {
			brackets++
		}
	}

	if brackets != 2 {
		return false
	}

	return IsStart(m.Symbols[0]) && IsEnd(m.Symbols[len(m.Symbols)-1])
}
