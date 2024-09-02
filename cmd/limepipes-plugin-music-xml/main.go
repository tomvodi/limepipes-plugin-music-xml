package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/common"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/grpc_plugin"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/plugin_implementation"
	"google.golang.org/grpc"
)

// defaultGRPCServer returns a new gRPC server with the given options.
// Acts as a factory method for gRPC servers.
func defaultGRPCServer(opts []grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opts...)
}

func main() {
	impl := plugin_implementation.NewPluginImplementation()

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"xml": grpc_plugin.NewGrpcPlugin(
				impl,
			),
		},

		GRPCServer: defaultGRPCServer,
	})
}
