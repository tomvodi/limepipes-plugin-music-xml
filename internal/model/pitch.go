package model

import (
	"encoding/xml"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
)

type Pitch struct {
	XMLName xml.Name `xml:"pitch"`
	Step    string   `xml:"step"`
	Alter   int8     `xml:"alter,omitempty"` // -1 flat, 1 sharp
	Octave  uint8    `xml:"octave"`
}

func PitchFromMusicModel(p pitch.Pitch, acc accidental.Accidental) *Pitch {
	switch p {
	case pitch.Pitch_LowG:
		return createPitch("G", 4, 0, acc)
	case pitch.Pitch_LowA:
		return createPitch("A", 4, 0, acc)
	case pitch.Pitch_B:
		return createPitch("B", 4, 0, acc)
	case pitch.Pitch_C:
		return createPitch("C", 5, 1, acc)
	case pitch.Pitch_D:
		return createPitch("D", 5, 0, acc)
	case pitch.Pitch_E:
		return createPitch("E", 5, 0, acc)
	case pitch.Pitch_F:
		return createPitch("F", 5, 1, acc)
	case pitch.Pitch_HighG:
		return createPitch("G", 5, 0, acc)
	case pitch.Pitch_HighA:
		return createPitch("A", 5, 0, acc)
	}

	return &Pitch{}
}

func createPitch(
	step string,
	octave uint8,
	alter int8,
	acc accidental.Accidental,
) *Pitch {
	retPitch := &Pitch{
		XMLName: xml.Name{Local: "pitch"},
		Step:    step,
		Octave:  octave,
		Alter:   alter,
	}
	if acc == accidental.Accidental_Sharp {
		retPitch.Alter += 1
	}
	if acc == accidental.Accidental_Flat {
		retPitch.Alter -= 1
	}

	return retPitch
}
