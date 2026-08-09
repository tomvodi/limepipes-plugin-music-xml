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

func regBubbly() *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  pitch.Pitch_C,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type: emb.Type_Grip,
			},
		},
	}
}

func Test_bubblyExpander_regular_ExpandSymbol(t *testing.T) {
	utils.SetupConsoleLogger()
	g := NewGomegaWithT(t)
	type fields struct {
		symbol    *symbols.Symbol
		prevPitch pitch.Pitch
		want      []pitch.Pitch
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
	}{
		{
			name: "regular",
			prepare: func(f *fields) {
				f.symbol = regBubbly()
				f.want = []pitch.Pitch{pitch.Pitch_LowG, pitch.Pitch_D, pitch.Pitch_LowG, pitch.Pitch_C, pitch.Pitch_LowG}
			},
		},
		{
			name: "regular with previous low g => half bubbly",
			prepare: func(f *fields) {
				f.symbol = regBubbly()
				f.prevPitch = pitch.Pitch_LowG
				f.want = []pitch.Pitch{pitch.Pitch_D, pitch.Pitch_LowG, pitch.Pitch_C, pitch.Pitch_LowG}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewBubblysExpander()
			expanded := pack.ExpandSymbol(f.symbol, f.prevPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}
