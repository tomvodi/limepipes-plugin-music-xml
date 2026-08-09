package interfaces_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
)

func TestExpansions_Get(t *testing.T) {
	g := NewGomegaWithT(t)

	note := &symbols.Note{Pitch: pitch.Pitch_LowA}
	expanded := []pitch.Pitch{pitch.Pitch_HighG}

	exps := interfaces.Expansions{note: expanded}

	g.Expect(exps.Get(note)).To(Equal(expanded))

	// a note without an embellishment has no entry
	g.Expect(exps.Get(&symbols.Note{})).To(BeNil())

	// callers walk symbols that may hold no note at all, and a tune that was
	// never expanded has no map
	g.Expect(exps.Get(nil)).To(BeNil())
	g.Expect(interfaces.Expansions(nil).Get(note)).To(BeNil())
}
