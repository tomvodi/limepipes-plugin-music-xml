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

func regDoubling(pitch pitch.Pitch) *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  pitch,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type: emb.Type_Doubling,
			},
		},
	}
}

func thumbDoubling(pitch pitch.Pitch) *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  pitch,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type:    emb.Type_Doubling,
				Variant: emb.Variant_Thumb,
			},
		},
	}
}

func halfDoubling(pitch pitch.Pitch) *symbols.Symbol {
	return &symbols.Symbol{
		Note: &symbols.Note{
			Pitch:  pitch,
			Length: length.Length_Quarter,
			Embellishment: &emb.Embellishment{
				Type:    emb.Type_Doubling,
				Variant: emb.Variant_Half,
			},
		},
	}
}

func Test_dblExpander_regular_ExpandSymbol(t *testing.T) {
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
				f.symbol = regDoubling(pitch.Pitch_LowG)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_LowG, pitch.Pitch_D}
			},
		},
		{
			name: "low a",
			prepare: func(f *fields) {
				f.symbol = regDoubling(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_LowA, pitch.Pitch_D}
			},
		},
		{
			name: "b",
			prepare: func(f *fields) {
				f.symbol = regDoubling(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_B, pitch.Pitch_D}
			},
		},
		{
			name: "c",
			prepare: func(f *fields) {
				f.symbol = regDoubling(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_C, pitch.Pitch_D}
			},
		},
		{
			name: "d",
			prepare: func(f *fields) {
				f.symbol = regDoubling(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_D, pitch.Pitch_E}
			},
		},
		{
			name: "e",
			prepare: func(f *fields) {
				f.symbol = regDoubling(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_E, pitch.Pitch_F}
			},
		},
		{
			name: "f",
			prepare: func(f *fields) {
				f.symbol = regDoubling(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_F, pitch.Pitch_HighG}
			},
		},
		{
			name: "hg",
			prepare: func(f *fields) {
				f.symbol = regDoubling(pitch.Pitch_HighG)
				f.want = []pitch.Pitch{pitch.Pitch_HighG, pitch.Pitch_F}
			},
		},
		{
			name: "ha",
			prepare: func(f *fields) {
				f.symbol = regDoubling(pitch.Pitch_HighA)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_HighG}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewDoublingsExpander()
			expanded := pack.ExpandSymbol(f.symbol, pitch.Pitch_NoPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}

func Test_dblUnapcker_thumb_ExpandSymbol(t *testing.T) {
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
				f.symbol = thumbDoubling(pitch.Pitch_LowG)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_LowG, pitch.Pitch_D}
			},
		},
		{
			name: "low a",
			prepare: func(f *fields) {
				f.symbol = thumbDoubling(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_LowA, pitch.Pitch_D}
			},
		},
		{
			name: "b",
			prepare: func(f *fields) {
				f.symbol = thumbDoubling(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_B, pitch.Pitch_D}
			},
		},
		{
			name: "c",
			prepare: func(f *fields) {
				f.symbol = thumbDoubling(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_C, pitch.Pitch_D}
			},
		},
		{
			name: "d",
			prepare: func(f *fields) {
				f.symbol = thumbDoubling(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_D, pitch.Pitch_E}
			},
		},
		{
			name: "e",
			prepare: func(f *fields) {
				f.symbol = thumbDoubling(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_E, pitch.Pitch_F}
			},
		},
		{
			name: "f",
			prepare: func(f *fields) {
				f.symbol = thumbDoubling(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_HighA, pitch.Pitch_F, pitch.Pitch_HighG}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewDoublingsExpander()
			expanded := pack.ExpandSymbol(f.symbol, pitch.Pitch_NoPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}

func Test_dblUnapcker_half_ExpandSymbol(t *testing.T) {
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
				f.symbol = halfDoubling(pitch.Pitch_LowG)
				f.want = []pitch.Pitch{pitch.Pitch_LowG, pitch.Pitch_D}
			},
		},
		{
			name: "low a",
			prepare: func(f *fields) {
				f.symbol = halfDoubling(pitch.Pitch_LowA)
				f.want = []pitch.Pitch{pitch.Pitch_LowA, pitch.Pitch_D}
			},
		},
		{
			name: "b",
			prepare: func(f *fields) {
				f.symbol = halfDoubling(pitch.Pitch_B)
				f.want = []pitch.Pitch{pitch.Pitch_B, pitch.Pitch_D}
			},
		},
		{
			name: "c",
			prepare: func(f *fields) {
				f.symbol = halfDoubling(pitch.Pitch_C)
				f.want = []pitch.Pitch{pitch.Pitch_C, pitch.Pitch_D}
			},
		},
		{
			name: "d",
			prepare: func(f *fields) {
				f.symbol = halfDoubling(pitch.Pitch_D)
				f.want = []pitch.Pitch{pitch.Pitch_D, pitch.Pitch_E}
			},
		},
		{
			name: "e",
			prepare: func(f *fields) {
				f.symbol = halfDoubling(pitch.Pitch_E)
				f.want = []pitch.Pitch{pitch.Pitch_E, pitch.Pitch_F}
			},
		},
		{
			name: "f",
			prepare: func(f *fields) {
				f.symbol = halfDoubling(pitch.Pitch_F)
				f.want = []pitch.Pitch{pitch.Pitch_F, pitch.Pitch_HighG}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			pack := NewDoublingsExpander()
			expanded := pack.ExpandSymbol(f.symbol, pitch.Pitch_NoPitch)
			want := fmt.Sprintf("%v", f.want)
			got := fmt.Sprintf("%v", expanded)
			g.Expect(got).To(Equal(want))
		})
	}
}
