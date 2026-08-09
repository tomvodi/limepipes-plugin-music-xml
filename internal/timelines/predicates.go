package timelines

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/barline"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/boundary"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	tl "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/timeline"
)

// IsStart reports whether the symbol opens a time line bracket.
func IsStart(sym *symbols.Symbol) bool {
	return hasBoundary(sym, boundary.Boundary_Start)
}

// IsEnd reports whether the symbol closes a time line bracket.
func IsEnd(sym *symbols.Symbol) bool {
	return hasBoundary(sym, boundary.Boundary_End)
}

// Has reports whether the measure contains any time line symbol.
func Has(m *measure.Measure) bool {
	for _, sym := range m.Symbols {
		if sym.Timeline != nil {
			return true
		}
	}

	return false
}

func hasBoundary(sym *symbols.Symbol, b boundary.Boundary) bool {
	return sym != nil &&
		sym.Timeline != nil &&
		sym.Timeline.BoundaryType == b
}

// playedInParts returns the parts a time line belongs to, numbered from 1.
//
// Only the "second of N" family names a part other than the one it is written
// in. Every other time line — a plain first or second ending, a singling, a
// bis — is played where it stands, so it returns nil for those.
func playedInParts(t tl.Type) []int {
	switch t {
	case tl.Type_SecondOf2:
		return []int{2}
	case tl.Type_SecondOf3:
		return []int{3}
	case tl.Type_SecondOf4:
		return []int{4}
	case tl.Type_SecondOf5:
		return []int{5}
	case tl.Type_SecondOf6:
		return []int{6}
	case tl.Type_SecondOf7:
		return []int{7}
	case tl.Type_SecondOf8:
		return []int{8}
	case tl.Type_SecondOf2And4:
		return []int{2, 4}
	default:
		return nil
	}
}

// endsPart reports whether a heavy barline closes the part it sits at the end
// of. Parts are delimited by heavy barlines; the light ones are ordinary bar
// divisions inside a part.
func endsPart(b *barline.Barline) bool {
	if b == nil {
		return false
	}

	switch b.Type {
	case barline.Type_Heavy,
		barline.Type_HeavyHeavy,
		barline.Type_HeavyLight,
		barline.Type_LightHeavy:
		return true
	default:
		return false
	}
}
