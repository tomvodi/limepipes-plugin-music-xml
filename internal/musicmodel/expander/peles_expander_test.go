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

func regPele(p pitch.Pitch) *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  p,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type: emb.Type_Pele,
			},
		},
	}
}

func halfPele(p pitch.Pitch) *symbols.Symbol {
	return peleVar(p, emb.Variant_Half)
}

func thumbPele(p pitch.Pitch) *symbols.Symbol {
	return peleVar(p, emb.Variant_Thumb)
}

func peleVar(p pitch.Pitch, variant emb.Variant) *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  p,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type:    emb.Type_Pele,
				Variant: variant,
			},
		},
	}
}

func makeLight(sym *symbols.Symbol) *symbols.Symbol {
	sym.Note.Embellishment.Weight = emb.Weight_Light
	return sym
}

func Test_peleExpander_regular_ExpandSymbol(t *testing.T) {
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
			name: "Low A",
			prepare: func(f *fields) {
				f.symbol = regPele(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_LowA, pitch.Pitch_E, pitch.Pitch_LowA, pitch.Pitch_LowG}
			},
		},
		{
			name: "B",
			prepare: func(f *fields) {
				f.symbol = regPele(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_B, pitch.Pitch_E, pitch.Pitch_B, pitch.Pitch_LowG}
			},
		},
		{
			name: "C",
			prepare: func(f *fields) {
				f.symbol = regPele(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_C, pitch.Pitch_E, pitch.Pitch_C, pitch.Pitch_LowG}
			},
		},
		{
			name: "D",
			prepare: func(f *fields) {
				f.symbol = regPele(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_D, pitch.Pitch_E, pitch.Pitch_D, pitch.Pitch_LowG}
			},
		},
		{
			name: "D light",
			prepare: func(f *fields) {
				f.symbol = makeLight(regPele(pitch.Pitch_D))
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_D, pitch.Pitch_E, pitch.Pitch_D, pitch.Pitch_C}
			},
		},
		{
			name: "E",
			prepare: func(f *fields) {
				f.symbol = regPele(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_E, pitch.Pitch_F, pitch.Pitch_E, pitch.Pitch_LowA}
			},
		},
		{
			name: "F",
			prepare: func(f *fields) {
				f.symbol = regPele(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_F, pitch.Pitch_HighG, pitch.Pitch_F, pitch.Pitch_E}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewPelesExpander()
			expanded := pack.ExpandSymbol(f.symbol, f.prevPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}

func Test_peleExpander_half_ExpandSymbol(t *testing.T) {
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
			name: "Low A",
			prepare: func(f *fields) {
				f.symbol = halfPele(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_LowA, pitch.Pitch_E, pitch.Pitch_LowA, pitch.Pitch_LowG}
			},
		},
		{
			name: "B",
			prepare: func(f *fields) {
				f.symbol = halfPele(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_B, pitch.Pitch_E, pitch.Pitch_B, pitch.Pitch_LowG}
			},
		},
		{
			name: "C",
			prepare: func(f *fields) {
				f.symbol = halfPele(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_C, pitch.Pitch_E, pitch.Pitch_C, pitch.Pitch_LowG}
			},
		},
		{
			name: "D",
			prepare: func(f *fields) {
				f.symbol = halfPele(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_D, pitch.Pitch_E, pitch.Pitch_D, pitch.Pitch_LowG}
			},
		},
		{
			name: "D light",
			prepare: func(f *fields) {
				f.symbol = makeLight(halfPele(pitch.Pitch_D))
				f.want = []pitch.Pitch{pitch.Pitch_D, pitch.Pitch_E, pitch.Pitch_D, pitch.Pitch_C}
			},
		},
		{
			name: "E",
			prepare: func(f *fields) {
				f.symbol = halfPele(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_E, pitch.Pitch_F, pitch.Pitch_E, pitch.Pitch_LowA}
			},
		},
		{
			name: "F",
			prepare: func(f *fields) {
				f.symbol = halfPele(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_F, pitch.Pitch_HighG, pitch.Pitch_F, pitch.Pitch_E}
			},
		},
		{
			name: "High G",
			prepare: func(f *fields) {
				f.symbol = halfPele(pitch.Pitch_HighG)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_HighA, pitch.Pitch_HighG, pitch.Pitch_F}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewPelesExpander()
			expanded := pack.ExpandSymbol(f.symbol, f.prevPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}

func Test_peleExpander_thumb_ExpandSymbol(t *testing.T) {
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
			name: "Low A",
			prepare: func(f *fields) {
				f.symbol = thumbPele(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_LowA, pitch.Pitch_E, pitch.Pitch_LowA, pitch.Pitch_LowG}
			},
		},
		{
			name: "B",
			prepare: func(f *fields) {
				f.symbol = thumbPele(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_B, pitch.Pitch_E, pitch.Pitch_B, pitch.Pitch_LowG}
			},
		},
		{
			name: "C",
			prepare: func(f *fields) {
				f.symbol = thumbPele(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_C, pitch.Pitch_E, pitch.Pitch_C, pitch.Pitch_LowG}
			},
		},
		{
			name: "D",
			prepare: func(f *fields) {
				f.symbol = thumbPele(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_D, pitch.Pitch_E, pitch.Pitch_D, pitch.Pitch_LowG}
			},
		},
		{
			name: "D light",
			prepare: func(f *fields) {
				f.symbol = makeLight(thumbPele(pitch.Pitch_D))
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_D, pitch.Pitch_E, pitch.Pitch_D, pitch.Pitch_C}
			},
		},
		{
			name: "E",
			prepare: func(f *fields) {
				f.symbol = thumbPele(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_E, pitch.Pitch_F, pitch.Pitch_E, pitch.Pitch_LowA}
			},
		},
		{
			name: "F",
			prepare: func(f *fields) {
				f.symbol = thumbPele(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_F, pitch.Pitch_HighG, pitch.Pitch_F, pitch.Pitch_E}
			},
		},
		{
			name: "High G",
			prepare: func(f *fields) {
				f.symbol = thumbPele(pitch.Pitch_HighG)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_HighG, pitch.Pitch_HighA, pitch.Pitch_HighG, pitch.Pitch_F}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewPelesExpander()
			expanded := pack.ExpandSymbol(f.symbol, f.prevPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}
