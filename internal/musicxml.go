package musicxml

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/boundary"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
	mmtuplet "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tuplet"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/barline"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/tuplet"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/timelines"
)

func WriteScore(score *model.Score, writer io.Writer) error {
	data, err := xml.MarshalIndent(score, " ", "  ")
	if err != nil {
		return err
	}

	data = append([]byte(musicXMLHeader), data...)
	data = bytes.ReplaceAll(data, []byte("></grace>"), []byte("/>"))
	data = bytes.ReplaceAll(data, []byte("></repeat>"), []byte("/>"))
	data = bytes.ReplaceAll(data, []byte("></rest>"), []byte("/>"))
	data = bytes.ReplaceAll(data, []byte("></dot>"), []byte("/>"))
	data = bytes.ReplaceAll(data, []byte("></fermata>"), []byte("/>"))
	data = bytes.ReplaceAll(data, []byte("></tied>"), []byte("/>"))
	data = bytes.ReplaceAll(data, []byte("></tuplet>"), []byte("/>"))
	data = bytes.ReplaceAll(data, []byte("></ending>"), []byte("/>"))
	if _, err := writer.Write(data); err != nil {
		return err
	}

	return nil
}

func ReadScore(reader io.Reader) (*model.Score, error) {
	fileData, _ := io.ReadAll(reader)

	score := &model.Score{}
	err := xml.Unmarshal(fileData, score)
	if err != nil {
		return nil, err
	}

	return score, nil
}

// ScoreFromMusicModelTune converts a music model tune into a MusicXML score.
//
// The expander is taken rather than a ready-made set of expansions because the
// tune is rewritten before conversion: Normalize works on a copy, and the
// expansions have to be keyed on the notes of that copy.
func ScoreFromMusicModelTune(
	t *tune.Tune,
	expander interfaces.EmbellishmentExpander,
) (*model.Score, error) {
	normalized := timelines.Normalize(t)
	exps := expander.ExpandTune(normalized)

	var measures []model.Measure

	// A bracket can run over several bars, so the ending that is open carries
	// from one measure to the next.
	var openEnding []int

	for _, measure := range normalized.Measures {
		// A staff can start without carrying anything of its own. Such a
		// measure has no MusicXML equivalent and would only shift the
		// numbering of the ones that follow.
		if isEmptyMeasure(measure) {
			continue
		}

		xmlMeasure := xmlMeasureFromMusicModelMeasure(measure, exps, len(measures), 32)
		openEnding = addEndings(&xmlMeasure, measure, openEnding)

		measures = append(measures, xmlMeasure)
	}

	score := &model.Score{
		XMLName: xml.Name{
			Local: "score-partwise",
		},
		Version: "3.1",
		PartList: model.ScorePartList{
			XMLName: xml.Name{
				Local: "part-list",
			},
			Parts: []model.ScorePart{
				{
					XMLName: xml.Name{
						Local: "score-part",
					},
					Id:   "P1",
					Name: "Bagpipe",
					Instrument: model.ScoreInstrument{
						XMLName: xml.Name{
							Local: "score-instrument",
						},
						Id:   "P1-I1",
						Name: "Bagpipe",
					},
					MidiDevice: model.MidiDevice{
						XMLName: xml.Name{
							Local: "midi-device",
						},
						Id:   "P1-I1",
						Port: 1,
					},
					MidiInstrument: model.MidiInstrument{
						XMLName: xml.Name{
							Local: "midi-instrument",
						},
						Id:      "P1-I1",
						Channel: 1,
						Program: 110,
						Volume:  78.7402,
						Pan:     0,
					},
				},
			},
		},
		Part: model.Part{
			XMLName: xml.Name{
				Local: "part",
			},
			Id:       "P1",
			Measures: measures,
		},
	}

	return score, nil
}

// addEndings hangs the first and second time brackets of a measure off its
// barlines: the opening one on the left, the closing one on the right.
//
// openEnding is the bracket still open when the measure begins, and the
// returned value is the one still open when it ends — a bracket may span
// several bars, in which case only the first and last carry a marker.
func addEndings(
	xmlMeasure *model.Measure,
	m *measure.Measure,
	openEnding []int,
) []int {
	for _, sym := range m.Symbols {
		switch {
		case timelines.IsStart(sym):
			times := timelines.EndingNumbers(sym.Timeline.Type)
			if times == nil {
				// not a first or second time, so nothing MusicXML can show
				continue
			}

			xmlMeasure.SetEnding(barline.Left, barline.NewEnding(times, barline.Start))
			openEnding = times

		case timelines.IsEnd(sym) && openEnding != nil:
			xmlMeasure.SetEnding(barline.Right, barline.NewEnding(openEnding, barline.Stop))
			openEnding = nil
		}
	}

	return openEnding
}

// isEmptyMeasure reports whether a measure carries nothing at all: no barlines,
// no time signature and no symbols.
func isEmptyMeasure(m *measure.Measure) bool {
	return m.LeftBarline == nil &&
		m.RightBarline == nil &&
		m.Time == nil &&
		len(m.Symbols) == 0
}

func xmlMeasureFromMusicModelMeasure(
	measure *measure.Measure,
	exps interfaces.Expansions,
	idx int,
	divisions uint8,
) model.Measure {
	xmlMeasure := model.Measure{
		XMLName: xml.Name{
			Local: "measure",
		},
		Number: idx + 1,
	}
	if idx == 0 {
		xmlMeasure.Attributes = model.NewAttributesWithKey(divisions)
	}
	if measure.Time != nil {
		xmlTime := model.NewTime(measure.Time)
		if xmlMeasure.Attributes != nil {
			xmlMeasure.Attributes.Time = xmlTime
		} else {
			xmlMeasure.Attributes = model.NewAttributesMinimal()
			xmlMeasure.Attributes.Time = xmlTime
		}
	}
	if measure.LeftBarline != nil {
		bar := barline.FromMusicModel(measure.LeftBarline, barline.Left)
		xmlMeasure.Barlines = append(xmlMeasure.Barlines, bar)
	}
	if measure.RightBarline != nil {
		bar := barline.FromMusicModel(measure.RightBarline, barline.Right)
		xmlMeasure.Barlines = append(xmlMeasure.Barlines, bar)
	}
	var measureNotes []model.Note
	noteCtx := &model.NoteContext{}

	// A tuplet is its own symbol bracketing a run of notes. Every note inside
	// the bracket carries a time-modification, and the first and last one also
	// carry the tuplet start/stop notation.
	var pendingTupletStart *mmtuplet.Tuplet

	for _, symbol := range measure.Symbols {
		if symbol.Tuplet != nil {
			switch symbol.Tuplet.BoundaryType {
			case boundary.Boundary_Start:
				noteCtx.CurrentTuplet = symbol.Tuplet
				pendingTupletStart = symbol.Tuplet
			case boundary.Boundary_End:
				if len(measureNotes) > 0 {
					last := &measureNotes[len(measureNotes)-1]
					last.SetTuplet(tuplet.FromMusicModel(symbol.Tuplet))
				}
				noteCtx.CurrentTuplet = nil
			}

			continue
		}

		if symbol.IsNote() {
			symbolNotes := model.NotesFromMusicModel(
				symbol,
				exps.Get(symbol.Note),
				noteCtx,
				divisions,
			)
			if pendingTupletStart != nil && len(symbolNotes) > 0 {
				// the melody note is the last one, graces come before it
				melody := &symbolNotes[len(symbolNotes)-1]
				melody.SetTuplet(tuplet.FromMusicModel(pendingTupletStart))
				pendingTupletStart = nil
			}
			measureNotes = append(measureNotes, symbolNotes...)
		}
		if symbol.Rest != nil {
			rest := model.RestFromMusicModel(symbol.Rest, divisions)
			measureNotes = append(measureNotes, rest)
		}
	}
	xmlMeasure.Notes = measureNotes
	return xmlMeasure
}
