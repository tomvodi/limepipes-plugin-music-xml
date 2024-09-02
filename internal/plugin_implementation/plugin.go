package plugin_implementation

import (
	"fmt"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/file_type"
	plugininterfaces "github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
)

type plug struct {
}

func (p *plug) ImportLocalFile(string) (*messages.ImportFileResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *plug) Import([]byte) (*messages.ImportFileResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *plug) PluginInfo() (*messages.PluginInfoResponse, error) {
	return &messages.PluginInfoResponse{
		Name:           "MusicXML Plugin",
		Description:    "Import Bagpipe Music Writer and Bagpipe Player files.",
		FileType:       file_type.Type_MUSIC_XML,
		Type:           messages.PluginType_OUT,
		FileExtensions: []string{".xml", ".musicxml"},
	}, nil
}

func NewPluginImplementation() plugininterfaces.LimePipesPlugin {
	return &plug{}
}
