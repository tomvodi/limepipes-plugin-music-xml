package pluginimplementation_test

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/musicmodel"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/fileformat"
	plugininterfaces "github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/musicmodel/expander"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/pluginimplementation"
)

func newPlugin() plugininterfaces.LimePipesPlugin {
	return pluginimplementation.NewPluginImplementation(
		expander.NewEmbellishmentExpander(),
	)
}

func tuneFromYaml(path string) *tune.Tune {
	data, err := os.ReadFile(path)
	Expect(err).ShouldNot(HaveOccurred())

	muMo := make(musicmodel.MusicModel, 0)
	Expect(yaml.Unmarshal(data, &muMo)).To(Succeed())
	Expect(muMo).ToNot(BeEmpty())

	return muMo[0]
}

var _ = Describe("PluginInfo", func() {
	It("describes an export only MusicXML plugin", func() {
		info, err := newPlugin().PluginInfo()

		Expect(err).ShouldNot(HaveOccurred())
		Expect(info).To(Equal(&messages.PluginInfoResponse{
			Name:           "MusicXML Plugin",
			Description:    "Export tunes to the MusicXML format.",
			FileFormat:     fileformat.Format_MUSIC_XML,
			Type:           messages.PluginType_OUT,
			FileExtensions: []string{".xml", ".musicxml"},
		}))
	})
})

var _ = Describe("Parsing", func() {
	It("is not supported from data", func() {
		_, err := newPlugin().Parse([]byte("anything"))

		Expect(err).Should(HaveOccurred())
	})

	It("is not supported from a file", func() {
		_, err := newPlugin().ParseFromFile("tune.musicxml")

		Expect(err).Should(HaveOccurred())
	})
})

var _ = Describe("Export", func() {
	var tunes []*tune.Tune

	BeforeEach(func() {
		tunes = []*tune.Tune{tuneFromYaml("../testfiles/four_measures.yaml")}
	})

	It("returns a MusicXML document for a single tune", func() {
		data, err := newPlugin().Export(tunes)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(string(data)).To(ContainSubstring("score-partwise"))
	})

	It("fails when there is nothing to export", func() {
		_, err := newPlugin().Export(nil)

		Expect(err).Should(HaveOccurred())
	})

	It("fails on more than one tune, because MusicXML holds only one", func() {
		_, err := newPlugin().Export(append(tunes, tunes[0]))

		Expect(err).Should(HaveOccurred())
	})
})

var _ = Describe("ExportToFile", func() {
	var tunes []*tune.Tune

	BeforeEach(func() {
		tunes = []*tune.Tune{tuneFromYaml("../testfiles/four_measures.yaml")}
	})

	It("writes the document to the given path", func() {
		path := filepath.Join(GinkgoT().TempDir(), "tune.musicxml")

		Expect(newPlugin().ExportToFile(tunes, path)).To(Succeed())

		written, err := os.ReadFile(path)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(string(written)).To(ContainSubstring("score-partwise"))
	})

	It("fails when the file cannot be created", func() {
		path := filepath.Join(GinkgoT().TempDir(), "no-such-dir", "tune.musicxml")

		Expect(newPlugin().ExportToFile(tunes, path)).ToNot(Succeed())
	})

	It("fails on more than one tune", func() {
		path := filepath.Join(GinkgoT().TempDir(), "tune.musicxml")

		Expect(newPlugin().ExportToFile(append(tunes, tunes[0]), path)).ToNot(Succeed())
	})
})
