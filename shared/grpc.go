package shared

import (
	"context"

	"github.com/ymtdzzz/sashimi/gen/go/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type gRPCClient struct {
	client proto.JobClient
}

func (c *gRPCClient) SplitJob(ctx context.Context) ([]string, error) {
	res, err := c.client.SplitJob(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return res.GetCommands(), nil
}

type gRPCServer struct {
	Impl Job
	proto.UnimplementedJobServer
}

func (s *gRPCServer) SplitJob(ctx context.Context, req *emptypb.Empty) (*proto.SplitJobResponse, error) {
	v, err := s.Impl.SplitJob(ctx)
	return &proto.SplitJobResponse{Commands: v}, err
}
