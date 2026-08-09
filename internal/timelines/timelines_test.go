package timelines

import (
	"os"

	"github.com/goccy/go-yaml"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/helper"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/musicmodel"
	tl "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/timeline"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
)

// tuneFromYaml loads a single tune fixture. The fixtures are the music model as
// the bww plugin parses it, so they say exactly what the converter is handed.
func tuneFromYaml(path string) *tune.Tune {
	data, err := os.ReadFile(path)
	Expect(err).ShouldNot(HaveOccurred())

	muMo := make(musicmodel.MusicModel, 0)
	Expect(yaml.Unmarshal(data, &muMo)).To(Succeed())
	Expect(muMo).To(HaveLen(1))

	return muMo[0]
}

var _ = Describe("splitIntoParts", func() {
	When("a tune has two parts separated by heavy barlines", func() {
		It("returns one group of measures per part", func() {
			t := tuneFromYaml("./testfiles/tune_with_two_parts.yaml")

			parts := splitIntoParts(t.Measures)

			Expect(parts).To(HaveLen(2))
			Expect(parts[0]).To(HaveLen(3))
			Expect(parts[1]).To(HaveLen(3))
		})

		It("keeps every measure", func() {
			t := tuneFromYaml("./testfiles/tune_with_two_parts.yaml")

			parts := splitIntoParts(t.Measures)

			Expect(joinParts(parts)).To(HaveLen(len(t.Measures)))
		})
	})
})

var _ = Describe("splitMeasuresAtTimelines", func() {
	When("a bar holds a first and a second time", func() {
		var original, split []*measure.Measure

		BeforeEach(func() {
			t := tuneFromYaml("./testfiles/tune_with_normal_1_2_part.yaml")
			original = t.Measures
			split = splitMeasuresAtTimelines(t.Measures)
		})

		It("gives the notes before the brackets and each bracket a bar", func() {
			// | B ⌐1 D ¬ ⌐2 E ¬ |  ->  | B | ⌐1 D ¬ | ⌐2 E ¬ |
			// plus the empty bar the tune starts with
			Expect(split).To(HaveLen(4))
		})

		It("leaves each bracket spanning a whole bar", func() {
			for _, m := range split {
				if !Has(m) {
					continue
				}

				Expect(IsStart(m.Symbols[0])).To(BeTrue())
				Expect(IsEnd(m.Symbols[len(m.Symbols)-1])).To(BeTrue())
			}
		})

		It("keeps the outer barlines of the bar it cut up", func() {
			source := original[1]
			pieces := split[1:]

			Expect(pieces[0].LeftBarline).To(Equal(source.LeftBarline))
			Expect(pieces[len(pieces)-1].RightBarline).To(Equal(source.RightBarline))
		})
	})

	When("a bar holds no time line", func() {
		It("leaves it alone", func() {
			t := tuneFromYaml("./testfiles/tune_with_two_parts.yaml")

			Expect(splitMeasuresAtTimelines(t.Measures)).
				To(BeComparableTo(t.Measures, helper.MusicModelCompareOptions))
		})
	})
})

var _ = Describe("distributeTimelines", func() {
	// Each case is a pair of fixtures: the tune as written, and the same tune
	// with the second time already moved by hand. Normalizing the first has to
	// produce the second.
	DescribeTable("moves a second time to the part it is played in",
		func(written, expected string) {
			source := tuneFromYaml(written)
			want := tuneFromYaml(expected)

			got := joinParts(distributeTimelines(splitIntoParts(source.Measures)))

			Expect(got).To(BeComparableTo(want.Measures, helper.MusicModelCompareOptions))
		},
		Entry("when the bracket sits inside one bar",
			"./testfiles/tune_with_2_of_2_part.yaml",
			"./testfiles/tune_with_2_of_2_part_normalized.yaml"),
		Entry("when the bracket spans several bars",
			"./testfiles/tune_with_2_of_2_spanning_over_measures.yaml",
			"./testfiles/tune_with_2_of_2_spanning_over_measures_normalized.yaml"),
	)
})

// tuneFromYamlText builds a tune from an inline fixture, for the odd shapes
// that are easier to read here than as a file.
func tuneFromYamlText(text string) *tune.Tune {
	muMo := make(musicmodel.MusicModel, 0)
	Expect(yaml.Unmarshal([]byte(text), &muMo)).To(Succeed())
	Expect(muMo).To(HaveLen(1))

	return muMo[0]
}

// twoPartsWithBracketIn puts a single time line bracket of the given type into
// the given part of a two part tune. The other part is a bare note.
func twoPartsWithBracketIn(part int, timelineType string) string {
	bracket := `
  - right_barline: {type: Heavy}
    symbols:
    - timeline: {type: ` + timelineType + `, boundary_type: Start}
    - note: {pitch: B, length: Quarter}
    - timeline: {boundary_type: End}`

	plain := `
  - right_barline: {type: Heavy}
    symbols:
    - note: {pitch: LowA, length: Quarter}`

	if part == 1 {
		return "- title: t\n  measures:" + bracket + plain
	}

	return "- title: t\n  measures:" + plain + bracket
}

var _ = Describe("distributeTimelines, when the bracket cannot be moved", func() {
	It("leaves it alone if the tune has no such part", func() {
		// a "second of 4" in a tune that only has two parts
		text := twoPartsWithBracketIn(1, "SecondOf4")
		source := tuneFromYamlText(text)
		untouched := tuneFromYamlText(text)

		got := joinParts(distributeTimelines(splitIntoParts(source.Measures)))

		Expect(got).To(BeComparableTo(untouched.Measures, helper.MusicModelCompareOptions))
	})

	It("leaves it alone if the part it names has no first time", func() {
		text := twoPartsWithBracketIn(1, "SecondOf2")
		source := tuneFromYamlText(text)
		untouched := tuneFromYamlText(text)

		got := joinParts(distributeTimelines(splitIntoParts(source.Measures)))

		Expect(got).To(BeComparableTo(untouched.Measures, helper.MusicModelCompareOptions))
	})
})

var _ = Describe("distributeTimelines, when the bracket is already in place", func() {
	It("turns a second of 2 written in part 2 into an ordinary second time", func() {
		source := tuneFromYamlText(twoPartsWithBracketIn(2, "SecondOf2"))

		got := joinParts(distributeTimelines(splitIntoParts(source.Measures)))

		Expect(got).To(HaveLen(2))
		Expect(got[1].Symbols[0].Timeline.Type).To(Equal(tl.Type_Second))
	})
})

var _ = Describe("Normalize", func() {
	It("does not modify the tune it is given", func() {
		source := tuneFromYaml("./testfiles/tune_with_2_of_2_part.yaml")
		untouched := tuneFromYaml("./testfiles/tune_with_2_of_2_part.yaml")

		Normalize(source)

		Expect(source).To(BeComparableTo(untouched, helper.MusicModelCompareOptions))
	})

	It("leaves a tune without time lines as it is", func() {
		source := tuneFromYaml("./testfiles/tune_with_two_parts.yaml")

		Expect(Normalize(source)).
			To(BeComparableTo(source, helper.MusicModelCompareOptions))
	})
})
