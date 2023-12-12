package shared

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/ymtdzzz/sashimi/gen/go/proto"
	"google.golang.org/grpc"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SASHIMI_PLUGIN",
	MagicCookieValue: "sashimi",
}

var PluginMap = map[string]plugin.Plugin{
	"job": &JobGRPCPlugin{},
}

type Job interface {
	SplitJob(ctx context.Context) ([]string, error)
}

type JobGRPCPlugin struct {
	plugin.Plugin
	Impl Job
}

func (p *JobGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterJobServer(s, &gRPCServer{Impl: p.Impl})
	return nil
}

func (p *JobGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (any, error) {
	return &gRPCClient{client: proto.NewJobClient(c)}, nil
}
