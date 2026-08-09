package plugin_implementation

import (
	"bytes"
	"fmt"
	"os"

	plugininterfaces "github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"

	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/fileformat"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
	musicxml "github.com/tomvodi/limepipes-plugin-music-xml/internal"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/interfaces"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/model"
)

type plug struct {
	embExpander interfaces.EmbellishmentExpander
}

// ParseFromFile is not supported. This plugin only exports to MusicXML.
func (p *plug) ParseFromFile(string) ([]*messages.ParsedTune, error) {
	return nil, fmt.Errorf("parsing MusicXML files is not supported")
}

// Parse is not supported. This plugin only exports to MusicXML.
func (p *plug) Parse([]byte) ([]*messages.ParsedTune, error) {
	return nil, fmt.Errorf("parsing MusicXML files is not supported")
}

func (p *plug) Export(tunes []*tune.Tune) ([]byte, error) {
	score, err := p.scoreFromTunes(tunes)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := musicxml.WriteScore(score, buf); err != nil {
		return nil, fmt.Errorf("failed writing MusicXML score: %w", err)
	}

	return buf.Bytes(), nil
}

func (p *plug) ExportToFile(tunes []*tune.Tune, filePath string) error {
	score, err := p.scoreFromTunes(tunes)
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed creating file %s: %w", filePath, err)
	}
	defer f.Close()

	if err := musicxml.WriteScore(score, f); err != nil {
		return fmt.Errorf("failed writing MusicXML score to %s: %w", filePath, err)
	}

	return nil
}

// scoreFromTunes converts a single tune to a MusicXML score.
// MusicXML is a single tune format, so exporting more than one tune at a time
// is rejected.
func (p *plug) scoreFromTunes(tunes []*tune.Tune) (*model.Score, error) {
	if len(tunes) == 0 {
		return nil, fmt.Errorf("no tunes to export")
	}
	if len(tunes) > 1 {
		return nil, fmt.Errorf(
			"MusicXML is a single tune format, got %d tunes to export", len(tunes),
		)
	}

	t := tunes[0]
	exps := p.embExpander.ExpandTune(t)

	return musicxml.ScoreFromMusicModelTune(t, exps)
}

func (p *plug) PluginInfo() (*messages.PluginInfoResponse, error) {
	return &messages.PluginInfoResponse{
		Name:           "MusicXML Plugin",
		Description:    "Export tunes to the MusicXML format.",
		FileFormat:     fileformat.Format_MUSIC_XML,
		Type:           messages.PluginType_OUT,
		FileExtensions: []string{".xml", ".musicxml"},
	}, nil
}

func NewPluginImplementation(
	embExpander interfaces.EmbellishmentExpander,
) plugininterfaces.LimePipesPlugin {
	return &plug{
		embExpander: embExpander,
	}
}
