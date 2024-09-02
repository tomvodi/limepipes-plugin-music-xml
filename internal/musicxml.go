package musicxml

import (
	"bytes"
	"encoding/xml"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model/barline"
	"io"
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

func ScoreFromMusicModelTune(tune *tune.Tune) (*model.Score, error) {
	var measures []model.Measure
	for i, measure := range tune.Measures {
		xmlMeasure := xmlMeasureFromMusicModelMeasure(measure, i, 32)
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

func xmlMeasureFromMusicModelMeasure(measure *measure.Measure, idx int, divisions uint8) model.Measure {
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
	for _, symbol := range measure.Symbols {
		if symbol.IsNote() {
			symbolNotes := model.NotesFromMusicModel(symbol.Note, noteCtx, divisions)
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
