package expander

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	emb "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/utils"
)

func doubleGrace(melody, grace pitch.Pitch) *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  melody,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type:  emb.Type_DoubleGrace,
				Pitch: grace,
			},
		},
	}
}

func Test_dblGraceExpander_ExpandSymbol(t *testing.T) {
	utils.SetupConsoleLogger()
	g := NewGomegaWithT(t)

	tests := []struct {
		name   string
		symbol *symbols.Symbol
		want   []pitch.Pitch
	}{
		{
			// the grace note, then the note below the melody note
			name:   "double grace on D",
			symbol: doubleGrace(pitch.Pitch_D, pitch.Pitch_HighG),
			want:   []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_C},
		},
		{
			name:   "double grace on E",
			symbol: doubleGrace(pitch.Pitch_E, pitch.Pitch_HighA),
			want:   []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_D},
		},
		{
			// there is no note below low G to play the second grace on
			name:   "cannot be played on low G",
			symbol: doubleGrace(pitch.Pitch_LowG, pitch.Pitch_HighG),
			want:   nil,
		},
		{
			// the grace note has to be above the melody note
			name:   "cannot be played below the melody note",
			symbol: doubleGrace(pitch.Pitch_HighA, pitch.Pitch_D),
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			got := NewDoubleGraceExpander().ExpandSymbol(tt.symbol, pitch.Pitch_NoPitch)

			g.Expect(fmt.Sprintf("%v", got)).To(Equal(fmt.Sprintf("%v", tt.want)))
		})
	}
}
