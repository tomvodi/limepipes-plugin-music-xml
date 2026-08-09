package expander

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	emb "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

// embKey identifies an embellishment for the expander lookup.
// The generated emb.Embellishment message can't be used as a map key itself
// because the embedded protobuf message state makes it non comparable.
type embKey struct {
	Type    emb.Type
	Pitch   pitch.Pitch
	Variant emb.Variant
	Weight  emb.Weight
}

func keyFromEmbellishment(e *emb.Embellishment) embKey {
	if e == nil {
		return embKey{}
	}

	return embKey{
		Type:    e.Type,
		Pitch:   e.Pitch,
		Variant: e.Variant,
		Weight:  e.Weight,
	}
}

type ExpandTable map[embKey]interfaces.SymbolExpander

//nolint:funlen // the table is a flat enumeration of all supported embellishments
func newSymbolExpanderTable() ExpandTable {
	singleGraceExp := NewSingleGraceExpander()
	dblExpander := NewDoublingsExpander()
	strikesExpander := NewStrikesExpander()
	gripsExpander := NewGripsExpander()
	taorExpander := NewTaorluathsExpander()
	birlsExpander := NewBirlsExpander()
	throwdExpander := NewThrowdsExpander()
	pelesExpander := NewPelesExpander()
	doubleStrikesExpander := NewDoubleStrikesExpander()
	tripleStrikesExpander := NewTripleStrikesExpander()
	dblGraceExpander := NewDoubleGraceExpander()

	return ExpandTable{
		{Type: emb.Type_SingleGrace, Pitch: pitch.Pitch_LowA}:  singleGraceExp,
		{Type: emb.Type_SingleGrace, Pitch: pitch.Pitch_B}:     singleGraceExp,
		{Type: emb.Type_SingleGrace, Pitch: pitch.Pitch_C}:     singleGraceExp,
		{Type: emb.Type_SingleGrace, Pitch: pitch.Pitch_D}:     singleGraceExp,
		{Type: emb.Type_SingleGrace, Pitch: pitch.Pitch_E}:     singleGraceExp,
		{Type: emb.Type_SingleGrace, Pitch: pitch.Pitch_F}:     singleGraceExp,
		{Type: emb.Type_SingleGrace, Pitch: pitch.Pitch_HighG}: singleGraceExp,
		{Type: emb.Type_SingleGrace, Pitch: pitch.Pitch_HighA}: singleGraceExp,

		{Type: emb.Type_Doubling}:                             dblExpander,
		{Type: emb.Type_Doubling, Variant: emb.Variant_Thumb}: dblExpander,
		{Type: emb.Type_Doubling, Variant: emb.Variant_Half}:  dblExpander,

		// A single strike carries no pitch: it always strikes the note below
		// the melody note.
		{Type: emb.Type_Strike}:                           strikesExpander,
		{Type: emb.Type_Strike, Pitch: pitch.Pitch_LowG}:  strikesExpander,
		{Type: emb.Type_Strike, Pitch: pitch.Pitch_LowA}:  strikesExpander,
		{Type: emb.Type_Strike, Pitch: pitch.Pitch_B}:     strikesExpander,
		{Type: emb.Type_Strike, Pitch: pitch.Pitch_C}:     strikesExpander,
		{Type: emb.Type_Strike, Pitch: pitch.Pitch_D}:     strikesExpander,
		{Type: emb.Type_Strike, Pitch: pitch.Pitch_E}:     strikesExpander,
		{Type: emb.Type_Strike, Pitch: pitch.Pitch_F}:     strikesExpander,
		{Type: emb.Type_Strike, Pitch: pitch.Pitch_HighG}: strikesExpander,
		{Type: emb.Type_Strike, Variant: emb.Variant_G}:   strikesExpander,
		{
			Type:    emb.Type_Strike,
			Variant: emb.Variant_G,
			Weight:  emb.Weight_Light,
		}: strikesExpander,
		{Type: emb.Type_Strike, Variant: emb.Variant_Thumb}: strikesExpander,
		{
			Type:    emb.Type_Strike,
			Variant: emb.Variant_Thumb,
			Weight:  emb.Weight_Light,
		}: strikesExpander,
		{Type: emb.Type_Strike, Variant: emb.Variant_Half}: strikesExpander,
		{
			Type:    emb.Type_Strike,
			Variant: emb.Variant_Half,
			Weight:  emb.Weight_Light,
		}: strikesExpander,

		{Type: emb.Type_Grip}:                         gripsExpander,
		{Type: emb.Type_Grip, Pitch: pitch.Pitch_B}:   gripsExpander,
		{Type: emb.Type_Grip, Variant: emb.Variant_G}: gripsExpander,
		{
			Type:    emb.Type_Grip,
			Variant: emb.Variant_G,
			Pitch:   pitch.Pitch_B,
		}: gripsExpander,
		{Type: emb.Type_Grip, Variant: emb.Variant_Thumb}: gripsExpander,
		{
			Type:    emb.Type_Grip,
			Variant: emb.Variant_Thumb,
			Pitch:   pitch.Pitch_B,
		}: gripsExpander,
		{Type: emb.Type_Grip, Variant: emb.Variant_Half}: gripsExpander,
		{
			Type:    emb.Type_Grip,
			Variant: emb.Variant_Half,
			Pitch:   pitch.Pitch_B,
		}: gripsExpander,

		{Type: emb.Type_Taorluath}:                       taorExpander,
		{Type: emb.Type_Taorluath, Pitch: pitch.Pitch_B}: taorExpander,

		{Type: emb.Type_Bubbly}: NewBubblysExpander(),

		{Type: emb.Type_Birl}:                             birlsExpander,
		{Type: emb.Type_Birl, Variant: emb.Variant_G}:     birlsExpander,
		{Type: emb.Type_Birl, Variant: emb.Variant_Thumb}: birlsExpander,
		{Type: emb.Type_ABirl}:                            birlsExpander,
		{Type: emb.Type_GraceBirl}:                        birlsExpander,

		// An unweighted throw on D is the heavy one.
		{Type: emb.Type_ThrowD}:                           throwdExpander,
		{Type: emb.Type_ThrowD, Weight: emb.Weight_Light}: throwdExpander,
		{Type: emb.Type_ThrowD, Weight: emb.Weight_Heavy}: throwdExpander,

		{Type: emb.Type_Pele}:                             pelesExpander,
		{Type: emb.Type_Pele, Weight: emb.Weight_Light}:   pelesExpander,
		{Type: emb.Type_Pele, Variant: emb.Variant_Thumb}: pelesExpander,
		{
			Type:    emb.Type_Pele,
			Variant: emb.Variant_Thumb,
			Weight:  emb.Weight_Light,
		}: pelesExpander,
		{Type: emb.Type_Pele, Variant: emb.Variant_Half}: pelesExpander,
		{
			Type:    emb.Type_Pele,
			Variant: emb.Variant_Half,
			Weight:  emb.Weight_Light,
		}: pelesExpander,

		{Type: emb.Type_DoubleStrike}:                           doubleStrikesExpander,
		{Type: emb.Type_DoubleStrike, Weight: emb.Weight_Light}: doubleStrikesExpander,
		{Type: emb.Type_DoubleStrike, Variant: emb.Variant_G}:   doubleStrikesExpander,
		{
			Type:    emb.Type_DoubleStrike,
			Variant: emb.Variant_G,
			Weight:  emb.Weight_Light,
		}: doubleStrikesExpander,
		{Type: emb.Type_DoubleStrike, Variant: emb.Variant_Thumb}: doubleStrikesExpander,
		{
			Type:    emb.Type_DoubleStrike,
			Variant: emb.Variant_Thumb,
			Weight:  emb.Weight_Light,
		}: doubleStrikesExpander,
		{Type: emb.Type_DoubleStrike, Variant: emb.Variant_Half}: doubleStrikesExpander,
		{
			Type:    emb.Type_DoubleStrike,
			Variant: emb.Variant_Half,
			Weight:  emb.Weight_Light,
		}: doubleStrikesExpander,

		{Type: emb.Type_TripleStrike}:                           tripleStrikesExpander,
		{Type: emb.Type_TripleStrike, Weight: emb.Weight_Light}: tripleStrikesExpander,
		{Type: emb.Type_TripleStrike, Variant: emb.Variant_G}:   tripleStrikesExpander,
		{
			Type:    emb.Type_TripleStrike,
			Variant: emb.Variant_G,
			Weight:  emb.Weight_Light,
		}: tripleStrikesExpander,
		{Type: emb.Type_TripleStrike, Variant: emb.Variant_Thumb}: tripleStrikesExpander,
		{
			Type:    emb.Type_TripleStrike,
			Variant: emb.Variant_Thumb,
			Weight:  emb.Weight_Light,
		}: tripleStrikesExpander,
		{Type: emb.Type_TripleStrike, Variant: emb.Variant_Half}: tripleStrikesExpander,
		{
			Type:    emb.Type_TripleStrike,
			Variant: emb.Variant_Half,
			Weight:  emb.Weight_Light,
		}: tripleStrikesExpander,

		{Type: emb.Type_DoubleGrace, Pitch: pitch.Pitch_D}:     dblGraceExpander,
		{Type: emb.Type_DoubleGrace, Pitch: pitch.Pitch_E}:     dblGraceExpander,
		{Type: emb.Type_DoubleGrace, Pitch: pitch.Pitch_F}:     dblGraceExpander,
		{Type: emb.Type_DoubleGrace, Pitch: pitch.Pitch_HighG}: dblGraceExpander,
		{Type: emb.Type_DoubleGrace, Pitch: pitch.Pitch_HighA}: dblGraceExpander,
	}
}
