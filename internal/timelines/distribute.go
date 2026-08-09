package timelines

import (
	"github.com/rs/zerolog/log"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/boundary"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	tl "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/timeline"
	"google.golang.org/protobuf/proto"
)

// relocatable is a "second of N" bracket: where it is written and which parts
// it is actually played in.
//
// The measure and symbol positions are relative to the part it was found in.
type relocatable struct {
	partIndex int // part the bracket is written in, counting from 0

	openMeasure int // measure holding the opening marker
	openIndex   int // position of the opening marker within that measure
	endMeasure  int // measure holding the closing marker
	endIndex    int // position of the closing marker within that measure

	playedIn []int // part numbers it belongs to, counting from 1

	// marker identifies the bracket across edits. Positions shift as measures
	// are rewritten, but the symbol pointer stays the same.
	marker *symbols.Symbol
}

// distributeTimelines moves every "second of N" bracket to the part it is
// played in, so that each bracket ends up inside the repeat it belongs to.
//
// A bracket that cannot be placed — the part does not exist, or has no
// first-time bracket to sit next to — is left untouched and reported.
func distributeTimelines(parts [][]*measure.Measure) [][]*measure.Measure {
	// Every move rewrites measures and shifts everything after it, so rather
	// than keeping stale positions alive the search starts over after each
	// move. Tunes have a handful of time lines, so the repeated scan is free.
	unplaceable := map[*symbols.Symbol]bool{}

	for {
		bracket, found := findRelocatable(parts, unplaceable)
		if !found {
			return parts
		}

		moved := false
		parts, moved = moveToPlayedParts(parts, bracket)

		if !moved {
			unplaceable[bracket.marker] = true
		}
	}
}

// findRelocatable returns the first "second of N" bracket in the tune that has
// not already been given up on.
func findRelocatable(
	parts [][]*measure.Measure,
	skip map[*symbols.Symbol]bool,
) (relocatable, bool) {
	for partIndex, part := range parts {
		if found, ok := findRelocatableInPart(part, partIndex, skip); ok {
			return found, true
		}
	}

	return relocatable{}, false
}

// findRelocatableInPart walks one part's symbols, pairing each bracket's
// opening marker with its closing one, and returns the first pair that names
// another part to be played in.
func findRelocatableInPart(
	part []*measure.Measure,
	partIndex int,
	skip map[*symbols.Symbol]bool,
) (relocatable, bool) {
	var open *relocatable

	for _, at := range symbolsOf(part) {
		if IsStart(at.symbol) {
			open = openedAt(at, partIndex)

			continue
		}

		if !IsEnd(at.symbol) || open == nil {
			continue
		}

		open.endMeasure = at.measureIndex
		open.endIndex = at.symbolIndex

		if len(open.playedIn) > 0 && !skip[open.marker] {
			return *open, true
		}

		open = nil
	}

	return relocatable{}, false
}

func openedAt(at positioned, partIndex int) *relocatable {
	return &relocatable{
		partIndex:   partIndex,
		openMeasure: at.measureIndex,
		openIndex:   at.symbolIndex,
		playedIn:    playedInParts(at.symbol.Timeline.Type),
		marker:      at.symbol,
	}
}

// positioned is a symbol together with where it sits in its part.
type positioned struct {
	symbol       *symbols.Symbol
	measureIndex int
	symbolIndex  int
}

// symbolsOf lists a part's symbols in playing order, each with its position, so
// that a scan over a part reads as one loop instead of two nested ones.
func symbolsOf(part []*measure.Measure) []positioned {
	var all []positioned

	for measureIndex, m := range part {
		for symbolIndex, sym := range m.Symbols {
			all = append(all, positioned{
				symbol:       sym,
				measureIndex: measureIndex,
				symbolIndex:  symbolIndex,
			})
		}
	}

	return all
}

// moveToPlayedParts copies the bracket's music into each part it is played in
// and, if that succeeded at least once, strips the bracket from where it was
// written. It reports whether anything was moved.
func moveToPlayedParts(
	parts [][]*measure.Measure,
	bracket relocatable,
) ([][]*measure.Measure, bool) {
	source := parts[bracket.partIndex]
	music := enclosedMusic(source, bracket)

	movedAny := false

	// A bracket written in the very part it is played in stays where it is; it
	// only has to stop calling itself a "second of N". Rewriting it in place
	// also keeps the positions in bracket valid, which copying into the part
	// being read from would not.
	staysInSource := false

	for _, partNumber := range bracket.playedIn {
		targetIndex := partNumber - 1

		if targetIndex == bracket.partIndex {
			bracket.marker.Timeline.Type = tl.Type_Second
			staysInSource = true
			movedAny = true

			continue
		}

		if targetIndex < 0 || targetIndex >= len(parts) {
			log.Error().Msgf(
				"time line is played in part %d but the tune has only %d parts",
				partNumber, len(parts),
			)

			continue
		}

		target, placed := addSecondTime(parts[targetIndex], music)
		if !placed {
			log.Error().Msgf(
				"part %d has no first time bracket to add a second time to",
				partNumber,
			)

			continue
		}

		parts[targetIndex] = target
		movedAny = true
	}

	if movedAny && !staysInSource {
		removeMarkers(source, bracket)
	}

	return parts, movedAny
}

// enclosedMusic returns the symbols the bracket encloses, one run per measure
// it spans. A bracket that opens and closes in the same bar gives one run.
//
// The symbols are cloned, because the music stays where it was written as well
// as being added to the part it is played in.
func enclosedMusic(part []*measure.Measure, bracket relocatable) [][]*symbols.Symbol {
	if bracket.openMeasure == bracket.endMeasure {
		inSameBar := part[bracket.openMeasure].Symbols[bracket.openIndex+1 : bracket.endIndex]

		return [][]*symbols.Symbol{cloneSymbols(inSameBar)}
	}

	var runs [][]*symbols.Symbol

	// what follows the opening marker in its own bar
	first := part[bracket.openMeasure].Symbols[bracket.openIndex+1:]
	runs = append(runs, cloneSymbols(first))

	// whole bars in between
	for i := bracket.openMeasure + 1; i < bracket.endMeasure; i++ {
		runs = append(runs, cloneSymbols(part[i].Symbols))
	}

	// what precedes the closing marker in its own bar
	last := part[bracket.endMeasure].Symbols[:bracket.endIndex]
	runs = append(runs, cloneSymbols(last))

	return runs
}

// removeMarkers deletes the bracket's two time line symbols, leaving the notes
// they enclosed in place as ordinary music.
func removeMarkers(part []*measure.Measure, bracket relocatable) {
	// The closing marker goes first: removing it cannot shift the opening one,
	// which sits earlier.
	endBar := part[bracket.endMeasure]
	endBar.Symbols = removeAt(endBar.Symbols, bracket.endIndex)

	openBar := part[bracket.openMeasure]
	openBar.Symbols = removeAt(openBar.Symbols, bracket.openIndex)
}

// addSecondTime puts a second-time bracket holding music into the part,
// directly after the part's first-time bracket. It reports whether the part had
// a first time to attach to.
func addSecondTime(
	part []*measure.Measure,
	music [][]*symbols.Symbol,
) ([]*measure.Measure, bool) {
	hostMeasure, afterSymbol, found := findFirstTimeEnd(part)
	if !found {
		return part, false
	}

	host := part[hostMeasure]

	openBracket := []*symbols.Symbol{newMarker(tl.Type_Second, boundary.Boundary_Start)}
	closeBracket := []*symbols.Symbol{newMarker(tl.Type_NoType, boundary.Boundary_End)}

	head := host.Symbols[:afterSymbol+1]
	tail := append([]*symbols.Symbol(nil), host.Symbols[afterSymbol+1:]...)

	// The common case: the music fits in the bar that already holds the first
	// time, so the whole second time is spliced in beside it.
	if len(music) == 1 {
		host.Symbols = joinSymbols(head, openBracket, music[0], closeBracket, tail)

		return part, true
	}

	// The music spans several bars, so the host bar ends after the first run
	// and the remaining runs become bars of their own. The host's right barline
	// travels to the last of them, because that is where the bar now ends.
	host.Symbols = joinSymbols(head, openBracket, music[0])

	endOfHost := host.RightBarline
	host.RightBarline = nil

	var added []*measure.Measure
	for _, run := range music[1 : len(music)-1] {
		added = append(added, &measure.Measure{Symbols: run})
	}

	added = append(added, &measure.Measure{
		Symbols:      joinSymbols(music[len(music)-1], closeBracket, tail),
		RightBarline: endOfHost,
	})

	return insertMeasuresAfter(part, hostMeasure, added), true
}

// findFirstTimeEnd locates the closing marker of the part's first-time bracket,
// which is where a second time is attached.
func findFirstTimeEnd(part []*measure.Measure) (measureIndex, symbolIndex int, found bool) {
	inFirstTime := false

	for _, at := range symbolsOf(part) {
		if IsStart(at.symbol) {
			inFirstTime = at.symbol.Timeline.Type == tl.Type_First
		}

		if IsEnd(at.symbol) && inFirstTime {
			return at.measureIndex, at.symbolIndex, true
		}
	}

	return 0, 0, false
}

func newMarker(t tl.Type, b boundary.Boundary) *symbols.Symbol {
	return &symbols.Symbol{
		Timeline: &tl.TimeLine{
			BoundaryType: b,
			Type:         t,
		},
	}
}

func cloneSymbols(syms []*symbols.Symbol) []*symbols.Symbol {
	out := make([]*symbols.Symbol, 0, len(syms))
	for _, sym := range syms {
		out = append(out, proto.Clone(sym).(*symbols.Symbol))
	}

	return out
}

func joinSymbols(runs ...[]*symbols.Symbol) []*symbols.Symbol {
	var out []*symbols.Symbol
	for _, run := range runs {
		out = append(out, run...)
	}

	return out
}

func removeAt(syms []*symbols.Symbol, i int) []*symbols.Symbol {
	out := make([]*symbols.Symbol, 0, len(syms)-1)
	out = append(out, syms[:i]...)

	return append(out, syms[i+1:]...)
}

func insertMeasuresAfter(
	part []*measure.Measure,
	i int,
	added []*measure.Measure,
) []*measure.Measure {
	out := make([]*measure.Measure, 0, len(part)+len(added))
	out = append(out, part[:i+1]...)
	out = append(out, added...)

	return append(out, part[i+1:]...)
}
