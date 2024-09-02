package model

import (
	"encoding/xml"
	"github.com/rs/zerolog/log"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tie"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/fermata"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/tied"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/tuplet"
)

var stemUp = "up"
var stemDown = "down"

type NoteContext struct {
	CurrentTuplet *tuplet.Tuplet
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

func NotesFromMusicModel(
	n *symbols.Note,
	noteCtx *NoteContext,
	divisions uint8,
) []Note {
	var notes []Note

	if n.Embellishment != nil && n.ExpandedEmbellishment != nil {

		for i, pitch := range n.ExpandedEmbellishment {
			grace := Note{
				XMLName: xml.Name{
					Local: "n",
				},
				Grace: NewGrace(),
				Pitch: PitchFromMusicModel(pitch, accidental.Accidental_NoAccidental),
				Voice: 1,
				Type:  typeFromLength(length.Length_Thirtysecond),
				Stem:  &stemUp,
			}
			if len(n.ExpandedEmbellishment) > 1 {
				grace.Beams = embellishmentBeamsForPosition(i, len(n.ExpandedEmbellishment))
			}
			notes = append(notes, grace)
		}
	}

	xmlNote := Note{
		XMLName: xml.Name{
			Local: "n",
		},
		Pitch:      PitchFromMusicModel(n.Pitch, n.Accidental),
		Duration:   durationFromLength(n.Length, divisions),
		Voice:      1,
		Type:       typeFromLength(n.Length),
		Stem:       stemFromLength(n.Length),
		Accidental: NewAccidentalFromMusicModel(n.Accidental),
	}
	if n.Dots > 0 {
		for i := uint32(0); i < n.Dots; i++ {
			xmlNote.Dots = append(xmlNote.Dots, NewDot())
		}
	}
	if n.Tuplet != nil {
		xmlNote.TimeModification = NewTimeModification(n.Tuplet)
	}
	if noteCtx.CurrentTuplet != nil {
		xmlNote.TimeModification = NewTimeModification(noteCtx.CurrentTuplet)
	}
	var notations *Notations
	if n.Fermata || n.Tie != tie.Tie_NoTie || n.Tuplet != nil {
		notations = NewNotations()
	}

	if n.Fermata {
		notations.Fermata = fermata.NewFermata(fermata.Upright)
	}
	if n.Tie != tie.Tie_NoTie {
		switch n.Tie {
		case tie.Tie_Start:
			notations.Tied = tied.NewTied(tied.Start)
		case tie.Tie_End:
			notations.Tied = tied.NewTied(tied.Stop)
		}
	}
	if n.Tuplet != nil {
		notations.Tuplet = tuplet.FromMusicModel(n.Tuplet)
		if n.Tuplet.BoundaryType == tuplet.Start {
			noteCtx.CurrentTuplet = n.Tuplet
		}
		if n.Tuplet.BoundaryType == tuplet.Stop {
			noteCtx.CurrentTuplet = nil
		}
	}
	if notations != nil {
		xmlNote.Notations = notations
	}

	notes = append(notes, xmlNote)

	return notes
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
