package musicxml

import (
	"os"

	"github.com/goccy/go-yaml"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/musicmodel"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/musicmodel/expander"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/utils"
)

var embExpander = expander.NewEmbellishmentExpander()

// exportToMusicXML regenerates a reference file. Call it from a spec, run the
// suite once, then comment the call out again — the written file becomes the
// expectation that spec compares against from then on.
//
//nolint:unused // kept as the way reference files are produced
func exportToMusicXML(score *model.Score, filePath string) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	Expect(err).ShouldNot(HaveOccurred())
	defer f.Close()

	err = WriteScore(score, f)
	Expect(err).ShouldNot(HaveOccurred())
}

func importFromMusicXML(filePath string) *model.Score {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0660)
	Expect(err).ShouldNot(HaveOccurred())
	defer f.Close()

	score, err := ReadScore(f)
	Expect(err).ShouldNot(HaveOccurred())

	return score
}

func importFromYaml(filePath string) musicmodel.MusicModel {
	muMo := make(musicmodel.MusicModel, 0)
	fileData, err := os.ReadFile(filePath)
	Expect(err).ShouldNot(HaveOccurred())
	err = yaml.Unmarshal(fileData, &muMo)
	Expect(err).ShouldNot(HaveOccurred())

	return muMo
}

var _ = Describe("measureAttributes", func() {
	quarters := &measure.TimeSignature{Beats: 4, BeatType: 4}

	It("gives the first bar the divisions and key", func() {
		attrs := measureAttributes(&measure.Measure{}, 0, 32)

		Expect(attrs).ToNot(BeNil())
		Expect(attrs.Divisions).To(Equal(uint8(32)))
	})

	It("adds the time signature to the first bar's attributes", func() {
		attrs := measureAttributes(&measure.Measure{Time: quarters}, 0, 32)

		Expect(attrs).ToNot(BeNil())
		Expect(attrs.Divisions).To(Equal(uint8(32)))
		Expect(attrs.Time).ToNot(BeNil())
	})

	It("gives a later bar attributes only to carry a time change", func() {
		attrs := measureAttributes(&measure.Measure{Time: quarters}, 3, 32)

		Expect(attrs).ToNot(BeNil())
		Expect(attrs.Time).ToNot(BeNil())
		// divisions and key belong to the first bar only
		Expect(attrs.Divisions).To(Equal(uint8(0)))
	})

	It("gives a later bar without a time change no attributes at all", func() {
		Expect(measureAttributes(&measure.Measure{}, 3, 32)).To(BeNil())
	})
})

var _ = Describe("ScoreFromMusicModelTune", func() {
	utils.SetupConsoleLogger()
	var err error
	var score *model.Score
	var readScore *model.Score

	Context("having a tune with four measures", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/four_measures.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/four_measures.musicxml")
			readScore = importFromMusicXML("./testfiles/four_measures.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with all melody notes", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/all_melody_notes.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/all_melody_notes.musicxml")
			readScore = importFromMusicXML("./testfiles/all_melody_notes.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with single grace notes", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/single_graces.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/single_graces.musicxml")
			readScore = importFromMusicXML("./testfiles/single_graces.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with doublings", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/doublings.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/doublings.musicxml")
			readScore = importFromMusicXML("./testfiles/doublings.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with strikes", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/strikes.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/strikes.musicxml")
			readScore = importFromMusicXML("./testfiles/strikes.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with grips", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/grips.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/grips.musicxml")
			readScore = importFromMusicXML("./testfiles/grips.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with taorluaths", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/taorluaths.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/taorluaths.musicxml")
			readScore = importFromMusicXML("./testfiles/taorluaths.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with bubblys", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/bubblys.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/bubblys.musicxml")
			readScore = importFromMusicXML("./testfiles/bubblys.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with birls", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/birls.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/birls.musicxml")
			readScore = importFromMusicXML("./testfiles/birls.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with throw on Ds", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/throwds.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/throwds.musicxml")
			readScore = importFromMusicXML("./testfiles/throwds.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with peles", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/peles.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/peles.musicxml")
			readScore = importFromMusicXML("./testfiles/peles.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with double strikes", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/double_strikes.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/double_strikes.musicxml")
			readScore = importFromMusicXML("./testfiles/double_strikes.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with triple strikes", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/triple_strikes.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/triple_strikes.musicxml")
			readScore = importFromMusicXML("./testfiles/triple_strikes.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with double grace", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/double_grace.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/double_grace.musicxml")
			readScore = importFromMusicXML("./testfiles/double_grace.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with repeats", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/tune_with_repeats.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/tune_with_repeats.musicxml")
			readScore = importFromMusicXML("./testfiles/tune_with_repeats.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with accidentals", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/accidentals.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/accidentals.musicxml")
			readScore = importFromMusicXML("./testfiles/accidentals.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with rests", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/rests.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/rests.musicxml")
			readScore = importFromMusicXML("./testfiles/rests.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with dots", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/dots.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/dots.musicxml")
			readScore = importFromMusicXML("./testfiles/dots.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with fermatas", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/fermatas.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/fermatas.musicxml")
			readScore = importFromMusicXML("./testfiles/fermatas.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with ties", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/ties.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/ties.musicxml")
			readScore = importFromMusicXML("./testfiles/ties.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a tune with a first and a second time", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/tune_with_normal_1_2_part.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/tune_with_normal_1_2_part.musicxml")
			readScore = importFromMusicXML("./testfiles/tune_with_normal_1_2_part.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a tune with a second time played in another part", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/tune_with_2_of_2_part.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/tune_with_2_of_2_part.musicxml")
			readScore = importFromMusicXML("./testfiles/tune_with_2_of_2_part.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})

	Context("having a file with irregular groups", func() {
		BeforeEach(func() {
			muMo := importFromYaml("./testfiles/irregular_groups.yaml")
			score, err = ScoreFromMusicModelTune(muMo[0], embExpander)
			//exportToMusicXML(score, "./testfiles/irregular_groups.musicxml")
			readScore = importFromMusicXML("./testfiles/irregular_groups.musicxml")
		})

		It("should succeed", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readScore).Should(BeComparableTo(score))
		})
	})
})
