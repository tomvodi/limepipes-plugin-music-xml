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

func singleStrike(p pitch.Pitch) *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  p + 1,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type:  emb.Type_Strike,
				Pitch: p,
			},
		},
	}
}

func halfStrike(p pitch.Pitch) *symbols.Symbol {
	return strikeVariant(p, emb.Variant_Half)
}

func thumbStrike(p pitch.Pitch) *symbols.Symbol {
	return strikeVariant(p, emb.Variant_Thumb)
}

func gStrike(p pitch.Pitch) *symbols.Symbol {
	return strikeVariant(p, emb.Variant_G)
}

func strikeVariant(p pitch.Pitch, variant emb.Variant) *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  p,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type:    emb.Type_Strike,
				Variant: variant,
			},
		},
	}
}

func Test_strikesExpander_single_ExpandSymbol(t *testing.T) {
	utils.SetupConsoleLogger()
	g := NewGomegaWithT(t)
	type fields struct {
		symbol *symbols.Symbol
		want   []pitch.Pitch
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
	}{
		{
			name: "low g",
			prepare: func(f *fields) {
				f.symbol = singleStrike(pitch.Pitch_LowG)
				f.want = []pitch.Pitch{pitch.Pitch_LowG}
			},
		},
		{
			name: "low a",
			prepare: func(f *fields) {
				f.symbol = singleStrike(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_LowA}
			},
		},
		{
			name: "b",
			prepare: func(f *fields) {
				f.symbol = singleStrike(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_B}
			},
		},
		{
			name: "c",
			prepare: func(f *fields) {
				f.symbol = singleStrike(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_C}
			},
		},
		{
			name: "d",
			prepare: func(f *fields) {
				f.symbol = singleStrike(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_D}
			},
		},
		{
			name: "e",
			prepare: func(f *fields) {
				f.symbol = singleStrike(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_E}
			},
		},
		{
			name: "f",
			prepare: func(f *fields) {
				f.symbol = singleStrike(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_F}
			},
		},
		{
			name: "hg",
			prepare: func(f *fields) {
				f.symbol = singleStrike(pitch.Pitch_HighG)
				f.want = []pitch.Pitch{pitch.Pitch_HighG}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewStrikesExpander()
			expanded := pack.ExpandSymbol(f.symbol, pitch.Pitch_NoPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}

func Test_strikesExpander_half_ExpandSymbol(t *testing.T) {
	utils.SetupConsoleLogger()
	g := NewGomegaWithT(t)
	type fields struct {
		symbol *symbols.Symbol
		want   []pitch.Pitch
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
	}{
		{
			name: "low a",
			prepare: func(f *fields) {
				f.symbol = halfStrike(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_LowA, pitch.Pitch_LowG}
			},
		},
		{
			name: "b",
			prepare: func(f *fields) {
				f.symbol = halfStrike(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_B, pitch.Pitch_LowG}
			},
		},
		{
			name: "c",
			prepare: func(f *fields) {
				f.symbol = halfStrike(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_C, pitch.Pitch_LowG}
			},
		},
		{
			name: "d",
			prepare: func(f *fields) {
				f.symbol = halfStrike(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_D, pitch.Pitch_LowG}
			},
		},
		{
			name: "d light",
			prepare: func(f *fields) {
				str := halfStrike(pitch.Pitch_D)
				str.Note.Embellishment.Weight = emb.Weight_Light
				f.symbol = str
				f.want = []pitch.Pitch{pitch.Pitch_D, pitch.Pitch_C}
			},
		},
		{
			name: "e",
			prepare: func(f *fields) {
				f.symbol = halfStrike(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_E, pitch.Pitch_LowA}
			},
		},
		{
			name: "f",
			prepare: func(f *fields) {
				f.symbol = halfStrike(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_F, pitch.Pitch_E}
			},
		},
		{
			name: "hg",
			prepare: func(f *fields) {
				f.symbol = halfStrike(pitch.Pitch_HighG)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_F}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewStrikesExpander()
			expanded := pack.ExpandSymbol(f.symbol, pitch.Pitch_NoPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}

func Test_strikesExpander_thumb_ExpandSymbol(t *testing.T) {
	utils.SetupConsoleLogger()
	g := NewGomegaWithT(t)
	type fields struct {
		symbol *symbols.Symbol
		want   []pitch.Pitch
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
	}{
		{
			name: "low a",
			prepare: func(f *fields) {
				f.symbol = thumbStrike(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_LowA, pitch.Pitch_LowG}
			},
		},
		{
			name: "b",
			prepare: func(f *fields) {
				f.symbol = thumbStrike(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_B, pitch.Pitch_LowG}
			},
		},
		{
			name: "c",
			prepare: func(f *fields) {
				f.symbol = thumbStrike(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_C, pitch.Pitch_LowG}
			},
		},
		{
			name: "d",
			prepare: func(f *fields) {
				f.symbol = thumbStrike(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_D, pitch.Pitch_LowG}
			},
		},
		{
			name: "d light",
			prepare: func(f *fields) {
				str := thumbStrike(pitch.Pitch_D)
				str.Note.Embellishment.Weight = emb.Weight_Light
				f.symbol = str
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_D, pitch.Pitch_C}
			},
		},
		{
			name: "e",
			prepare: func(f *fields) {
				f.symbol = thumbStrike(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_E, pitch.Pitch_LowA}
			},
		},
		{
			name: "f",
			prepare: func(f *fields) {
				f.symbol = thumbStrike(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_F, pitch.Pitch_E}
			},
		},
		{
			name: "hg",
			prepare: func(f *fields) {
				f.symbol = thumbStrike(pitch.Pitch_HighG)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_HighG, pitch.Pitch_F}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewStrikesExpander()
			expanded := pack.ExpandSymbol(f.symbol, pitch.Pitch_NoPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}

func Test_strikesExpander_g_ExpandSymbol(t *testing.T) {
	utils.SetupConsoleLogger()
	g := NewGomegaWithT(t)
	type fields struct {
		symbol *symbols.Symbol
		want   []pitch.Pitch
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
	}{
		{
			name: "low a",
			prepare: func(f *fields) {
				f.symbol = gStrike(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_LowA, pitch.Pitch_LowG}
			},
		},
		{
			name: "b",
			prepare: func(f *fields) {
				f.symbol = gStrike(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_B, pitch.Pitch_LowG}
			},
		},
		{
			name: "c",
			prepare: func(f *fields) {
				f.symbol = gStrike(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_C, pitch.Pitch_LowG}
			},
		},
		{
			name: "d",
			prepare: func(f *fields) {
				f.symbol = gStrike(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_D, pitch.Pitch_LowG}
			},
		},
		{
			name: "d light",
			prepare: func(f *fields) {
				str := gStrike(pitch.Pitch_D)
				str.Note.Embellishment.Weight = emb.Weight_Light
				f.symbol = str
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_D, pitch.Pitch_C}
			},
		},
		{
			name: "e",
			prepare: func(f *fields) {
				f.symbol = gStrike(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_E, pitch.Pitch_LowA}
			},
		},
		{
			name: "f",
			prepare: func(f *fields) {
				f.symbol = gStrike(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_F, pitch.Pitch_E}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewStrikesExpander()
			expanded := pack.ExpandSymbol(f.symbol, pitch.Pitch_NoPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}
