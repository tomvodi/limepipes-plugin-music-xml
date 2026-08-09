package model

import (
	"encoding/xml"

	"github.com/rs/zerolog/log"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tie"
	mmtuplet "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tuplet"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/fermata"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/tied"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/tuplet"
)

var stemUp = "up"
var stemDown = "down"

type NoteContext struct {
	CurrentTuplet *mmtuplet.Tuplet
}

type Note struct {
	XMLName          xml.Name          `xml:"note"`
	Rest             *Rest             `xml:"rest,omitempty"`
	Grace            *Grace            `xml:"grace,omitempty"`
	Pitch            *Pitch            `xml:"pitch,omitempty"`
	Duration         uint8             `xml:"duration,omitempty"`
	Voice            uint8             `xml:"voice,omitempty"`
	Type             string            `xml:"type"`
	TimeModification *TimeModification `xml:"time-modification,omitempty"`
	Dots             []Dot             `xml:"dot,omitempty"`
	Accidental       *Accidental       `xml:"accidental,omitempty"`
	Stem             *string           `xml:"stem,omitempty"`
	Beams            []Beam            `xml:"beam,omitempty"`
	Notations        *Notations        `xml:"notations,omitempty"`
}

// NotesFromMusicModel converts a music model symbol into the MusicXML notes it
// consists of. An embellishment becomes a leading run of grace notes, so
// expanded carries the pitches the symbol's embellishment expands to. It comes
// in as a parameter because the music model has no field for it.
func NotesFromMusicModel(
	sym *symbols.Symbol,
	expanded []pitch.Pitch,
	noteCtx *NoteContext,
	divisions uint8,
) []Note {
	n := sym.Note

	notes := graceNotesFor(n, expanded)

	xmlNote := Note{
		XMLName: xml.Name{
			Local: "note",
		},
		Pitch:      PitchFromMusicModel(n.Pitch, n.Accidental),
		Duration:   durationFromLength(n.Length, divisions),
		Voice:      1,
		Type:       typeFromLength(n.Length),
		Stem:       stemFromLength(n.Length),
		Accidental: NewAccidentalFromMusicModel(n.Accidental),
		Dots:       dotsFor(n),
		Notations:  notationsFor(n),
	}

	if noteCtx.CurrentTuplet != nil {
		xmlNote.TimeModification = NewTimeModification(noteCtx.CurrentTuplet)
	}

	return append(notes, xmlNote)
}

// graceNotesFor renders an embellishment as the run of grace notes it is played
// as. A run of more than one is beamed together.
func graceNotesFor(n *symbols.Note, expanded []pitch.Pitch) []Note {
	if n.Embellishment == nil || expanded == nil {
		return nil
	}

	graces := make([]Note, 0, len(expanded))

	for i, gracePitch := range expanded {
		grace := Note{
			XMLName: xml.Name{
				Local: "note",
			},
			Grace: NewGrace(),
			Pitch: PitchFromMusicModel(gracePitch, accidental.Accidental_NoAccidental),
			Voice: 1,
			Type:  typeFromLength(length.Length_Thirtysecond),
			Stem:  &stemUp,
		}
		if len(expanded) > 1 {
			grace.Beams = embellishmentBeamsForPosition(i, len(expanded))
		}

		graces = append(graces, grace)
	}

	return graces
}

func dotsFor(n *symbols.Note) []Dot {
	if n.Dots == 0 {
		return nil
	}

	dots := make([]Dot, 0, n.Dots)
	for i := uint32(0); i < n.Dots; i++ {
		dots = append(dots, NewDot())
	}

	return dots
}

// notationsFor collects the marks that hang off a note rather than being part
// of it. Tuplet notations are added later by the caller walking the measure.
func notationsFor(n *symbols.Note) *Notations {
	if !n.Fermata && n.Tie == tie.Tie_NoTie {
		return nil
	}

	notations := NewNotations()

	if n.Fermata {
		notations.Fermata = fermata.NewFermata(fermata.Upright)
	}

	switch n.Tie {
	case tie.Tie_Start:
		notations.Tied = tied.NewTied(tied.Start)
	case tie.Tie_End:
		notations.Tied = tied.NewTied(tied.Stop)
	}

	return notations
}

// SetTuplet attaches a tuplet notation to the note, creating the notations
// element if the note doesn't have one yet.
// In the music model a tuplet is a standalone symbol bracketing a run of notes,
// so the boundary notations are attached by the caller walking the measure
// rather than by the note conversion itself.
func (n *Note) SetTuplet(tpl *tuplet.Tuplet) {
	if n.Notations == nil {
		n.Notations = NewNotations()
	}

	n.Notations.Tuplet = tpl
}

func RestFromMusicModel(rest *symbols.Rest, divisions uint8) Note {
	xmlNote := Note{
		XMLName: xml.Name{
			Local: "note",
		},
		Rest:     NewRest(),
		Duration: durationFromLength(rest.Length, divisions),
		Voice:    1,
		Type:     typeFromLength(rest.Length),
	}

	return xmlNote
}

func embellishmentBeamsForPosition(idx int, len int) []Beam {
	var bType BeamType
	if idx == 0 {
		bType = Begin
	} else if idx == len-1 {
		bType = End
	} else {
		bType = Continue
	}
	return getBeams(3, bType)
}

func getBeams(beamCnt uint8, bType BeamType) []Beam {
	beams := make([]Beam, beamCnt)
	for i := uint8(0); i < beamCnt; i++ {
		beams[i] = NewBeam(i+1, bType)
	}

	return beams
}

func typeFromLength(l length.Length) string {
	switch l {
	case length.Length_Whole:
		return "whole"
	case length.Length_Half:
		return "half"
	case length.Length_Quarter:
		return "quarter"
	case length.Length_Eighth:
		return "eighth"
	case length.Length_Sixteenth:
		return "16th"
	case length.Length_Thirtysecond:
		return "32nd"
	}

	return ""
}

func stemFromLength(l length.Length) *string {
	if l == length.Length_Whole {
		return nil
	}

	return &stemDown
}

func durationFromLength(l length.Length, divisions uint8) uint8 {
	maxDivisions := 255 / 4
	if divisions > uint8(maxDivisions) {
		log.Error().Msgf("divisions can't be greater than %d", maxDivisions)
		return 255
	}

	switch l {
	case length.Length_Whole:
		return 4 * divisions
	case length.Length_Half:
		return 2 * divisions
	case length.Length_Quarter:
		return 1 * divisions
	case length.Length_Eighth:
		return divisions / 2
	case length.Length_Sixteenth:
		return divisions / 4
	case length.Length_Thirtysecond:
		return divisions / 8
	}

	log.Error().Msgf("length %s not supported for calculation of note duration", l.String())

	return divisions
}
