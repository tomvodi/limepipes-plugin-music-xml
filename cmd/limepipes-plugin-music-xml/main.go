package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/common"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/fileformat"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/grpcplugin"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/musicmodel/expander"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/pluginimplementation"
	"google.golang.org/grpc"
)

// defaultGRPCServer returns a new gRPC server with the given options.
// Acts as a factory method for gRPC servers.
func defaultGRPCServer(opts []grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opts...)
}

func main() {
	impl := pluginimplementation.NewPluginImplementation(
		expander.NewEmbellishmentExpander(),
	)

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			fileformat.Format_MUSIC_XML.String(): grpcplugin.NewGrpcPlugin(
				impl,
			),
		},

		GRPCServer: defaultGRPCServer,
	})
}
